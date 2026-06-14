// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

//go:build integration

package postgres

import (
	"context"
	"testing"

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"
)

func TestSupplierRepository_FindQueries(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	inactiveOnly := mustCreateSupplierQueryRow(t, txCtx, repo, "Cahaya Motor", false)
	inactive := mustCreateSupplierQueryRow(t, txCtx, repo, "Terang Abadi", false)
	active := mustCreateSupplierQueryRow(t, txCtx, repo, " terang  abadi ", true)
	_, found, err := repo.FindByID(txCtx, domain.SupplierID("missing-supplier"))
	if err != nil || found {
		t.Fatalf("FindByID() = found %v err %v, want found false nil", found, err)
	}
	got, found, err := repo.FindByNormalizedName(txCtx, active.NormalizedName())
	if err != nil || !found || got.ID() != active.ID() {
		t.Fatalf("FindByNormalizedName() = id %q found %v err %v", got.ID(), found, err)
	}
	got, found, err = repo.FindActiveByNormalizedName(txCtx, inactive.NormalizedName())
	if err != nil || !found || got.ID() != active.ID() {
		t.Fatalf("FindActiveByNormalizedName() = id %q found %v err %v", got.ID(), found, err)
	}
	_, found, err = repo.FindActiveByNormalizedName(txCtx, inactiveOnly.NormalizedName())
	if err != nil || found {
		t.Fatalf("FindActiveByNormalizedName() inactive = found %v err %v, want false nil", found, err)
	}
}

func TestSupplierRepository_ListAndLookup(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	for _, row := range []struct {
		name   string
		active bool
	}{{"Gamma Parts", true}, {"Beta Parts", false}, {"Omega Parts", true}, {"Delta Parts", false}, {"Alpha Parts", true}} {
		mustCreateSupplierQueryRow(t, txCtx, repo, row.name, row.active)
	}
	list := func(filter ports.ListSuppliersFilter, names ...string) {
		rows, err := repo.List(txCtx, filter)
		assertSupplierNames(t, rows, err, names...)
	}
	lookup := func(filter ports.LookupSuppliersFilter, names ...string) {
		rows, err := repo.Lookup(txCtx, filter)
		assertSupplierNames(t, rows, err, names...)
	}
	list(ports.ListSuppliersFilter{Status: ports.ListStatusActive, Page: 1, PerPage: 10}, "Alpha Parts", "Gamma Parts", "Omega Parts")
	list(ports.ListSuppliersFilter{Status: ports.ListStatusInactive, Page: 1, PerPage: 10}, "Beta Parts", "Delta Parts")
	list(ports.ListSuppliersFilter{Status: ports.ListStatusAll, Page: 1, PerPage: 10}, "Alpha Parts", "Beta Parts", "Delta Parts", "Gamma Parts", "Omega Parts")
	list(ports.ListSuppliersFilter{Status: ports.ListStatusAll, Page: 2, PerPage: 2}, "Delta Parts", "Gamma Parts")
	list(ports.ListSuppliersFilter{Query: "ta", Status: ports.ListStatusAll, Page: 1, PerPage: 10}, "Beta Parts", "Delta Parts")
	lookup(ports.LookupSuppliersFilter{Limit: 10, ActiveOnly: true}, "Alpha Parts", "Gamma Parts", "Omega Parts")
	lookup(ports.LookupSuppliersFilter{Limit: 10}, "Alpha Parts", "Beta Parts", "Delta Parts", "Gamma Parts", "Omega Parts")
	lookup(ports.LookupSuppliersFilter{Limit: 2, ActiveOnly: true}, "Alpha Parts", "Gamma Parts")
	lookup(ports.LookupSuppliersFilter{Query: "ta", Limit: 10}, "Beta Parts", "Delta Parts")
}

func mustCreateSupplierQueryRow(
	t *testing.T, ctx context.Context, repo *SupplierRepository, name string, active bool,
) domain.Supplier {
	t.Helper()
	supplier := newSupplierRepositoryTestSupplier(t, name)
	if !active {
		supplier.Deactivate(supplierRepositoryTestTime.Add(2))
	}
	if err := repo.Create(ctx, supplier); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	return supplier
}

func assertSupplierNames(t *testing.T, rows []domain.Supplier, err error, names ...string) {
	t.Helper()
	if err != nil {
		t.Fatalf("query error = %v", err)
	}
	if len(rows) != len(names) {
		t.Fatalf("len(rows) = %d, want %d", len(rows), len(names))
	}
	for i, name := range names {
		if rows[i].Name() != name {
			t.Fatalf("rows[%d].Name() = %q, want %q", i, rows[i].Name(), name)
		}
	}
}
