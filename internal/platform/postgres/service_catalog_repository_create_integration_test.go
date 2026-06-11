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

	"pos-go/internal/modules/servicecatalog/domain"
)

func TestServiceCatalogRepository_CreateAndFind(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenServiceCatalogRepoTx(t, ctx)
	repo := NewServiceCatalogRepository(pool)

	item := newServiceCatalogTestItem(t, "Express Wash")

	if err := repo.Create(txCtx, item); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	got, found, err := repo.FindByID(txCtx, item.ID())
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}
	if !found {
		t.Fatal("FindByID() found = false, want true")
	}
	if got.Name() != "Express Wash" {
		t.Fatalf("Name() = %q, want %q", got.Name(), "Express Wash")
	}
	if got.NormalizedName() != domain.NormalizedName("express wash") {
		t.Fatalf("NormalizedName() = %q", got.NormalizedName())
	}
}

func TestServiceCatalogRepository_FindMissing(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenServiceCatalogRepoTx(t, ctx)
	repo := NewServiceCatalogRepository(pool)

	_, found, err := repo.FindByID(txCtx, domain.ServiceCatalogItemID("missing-item"))
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}
	if found {
		t.Fatal("FindByID() found = true, want false")
	}
}
