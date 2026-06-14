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
)

func TestSupplierRepository_SetActiveDeactivatesSupplier(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	supplier := newSupplierRepositoryTestSupplier(t, "Lancar Motor")
	if err := repo.Create(txCtx, supplier); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	got, found, err := repo.SetActive(txCtx, supplier.ID(), false)
	if err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}
	if !found || got.IsActive() {
		t.Fatalf("SetActive() = found %v active %v, want found true active false", found, got.IsActive())
	}
	if got.UpdatedAt().Equal(supplier.UpdatedAt()) {
		t.Fatal("UpdatedAt() did not change")
	}
	stored, found, err := repo.FindByID(txCtx, supplier.ID())
	if err != nil || !found || stored.IsActive() || !stored.UpdatedAt().Equal(got.UpdatedAt()) {
		t.Fatalf("stored supplier = found %v active %v err %v", found, stored.IsActive(), err)
	}
}

func TestSupplierRepository_SetActiveActivatesSupplier(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	supplier := newInactiveSupplierRepositoryTestSupplier(t, "Sumber Jaya")
	if err := repo.Create(txCtx, supplier); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	got, found, err := repo.SetActive(txCtx, supplier.ID(), true)
	if err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}
	if !found || !got.IsActive() {
		t.Fatalf("SetActive() = found %v active %v, want found true active true", found, got.IsActive())
	}
	if got.UpdatedAt().Equal(supplier.UpdatedAt()) {
		t.Fatal("UpdatedAt() did not change")
	}
	stored, found, err := repo.FindByID(txCtx, supplier.ID())
	if err != nil || !found || !stored.IsActive() || !stored.UpdatedAt().Equal(got.UpdatedAt()) {
		t.Fatalf("stored supplier = found %v active %v err %v", found, stored.IsActive(), err)
	}
}

func TestSupplierRepository_SetActiveMissing(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	_, found, err := repo.SetActive(txCtx, domain.SupplierID("missing-supplier"), false)
	if err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}
	if found {
		t.Fatal("SetActive() found = true, want false")
	}
}

func TestSupplierRepository_SetActiveRejectsDuplicateActivation(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	active := newSupplierRepositoryTestSupplier(t, "Prima Parts")
	inactive := newInactiveSupplierRepositoryTestSupplier(t, " prima   parts ")
	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create() active error = %v", err)
	}
	if err := repo.Create(txCtx, inactive); err != nil {
		t.Fatalf("Create() inactive error = %v", err)
	}
	_, _, err := repo.SetActive(txCtx, inactive.ID(), true)
	if !isSupplierActiveNameUniqueViolation(err) {
		t.Fatalf("SetActive() error = %v, want active name unique violation", err)
	}
}
