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

	"pos-go/internal/modules/servicecatalog/ports"
)

func TestServiceCatalogRepository_ListAndLookup(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenServiceCatalogRepoTx(t, ctx)
	repo := NewServiceCatalogRepository(pool)

	active := newServiceCatalogTestItem(t, "Alpha Service")
	inactive := newServiceCatalogTestItem(t, "Beta Service")

	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create active error = %v", err)
	}
	if err := repo.Create(txCtx, inactive); err != nil {
		t.Fatalf("Create inactive error = %v", err)
	}
	if _, _, err := repo.SetActive(txCtx, inactive.ID(), false); err != nil {
		t.Fatalf("SetActive inactive error = %v", err)
	}

	activeRows, err := repo.List(txCtx, ports.ListServiceCatalogItemsFilter{
		Status:  ports.ListStatusActive,
		Page:    1,
		PerPage: 10,
	})
	if err != nil {
		t.Fatalf("List active error = %v", err)
	}
	if len(activeRows) != 1 {
		t.Fatalf("List active len = %d, want 1", len(activeRows))
	}

	allRows, err := repo.List(txCtx, ports.ListServiceCatalogItemsFilter{
		Status:  ports.ListStatusAll,
		Page:    1,
		PerPage: 10,
	})
	if err != nil {
		t.Fatalf("List all error = %v", err)
	}
	if len(allRows) != 2 {
		t.Fatalf("List all len = %d, want 2", len(allRows))
	}

	lookupRows, err := repo.Lookup(txCtx, ports.LookupServiceCatalogItemsFilter{
		Query:      "service",
		Limit:      10,
		ActiveOnly: true,
	})
	if err != nil {
		t.Fatalf("Lookup error = %v", err)
	}
	if len(lookupRows) != 1 {
		t.Fatalf("Lookup len = %d, want 1", len(lookupRows))
	}
}
