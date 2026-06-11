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

func TestServiceCatalogRepository_Update(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenServiceCatalogRepoTx(t, ctx)
	repo := NewServiceCatalogRepository(pool)

	item := newServiceCatalogTestItem(t, "Basic Wash")
	if err := repo.Create(txCtx, item); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	if err := item.Update("Premium Wash", domain.MoneyRupiah(25000), item.UpdatedAt().Add(1)); err != nil {
		t.Fatalf("Update domain error = %v", err)
	}

	if err := repo.Update(txCtx, item); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, found, err := repo.FindByNormalizedName(txCtx, domain.NormalizedName("premium wash"))
	if err != nil {
		t.Fatalf("FindByNormalizedName() error = %v", err)
	}
	if !found {
		t.Fatal("FindByNormalizedName() found = false, want true")
	}
	if got.DefaultPriceRupiah() != domain.MoneyRupiah(25000) {
		t.Fatalf("DefaultPriceRupiah() = %d", got.DefaultPriceRupiah())
	}
}

func TestServiceCatalogRepository_SetActive(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenServiceCatalogRepoTx(t, ctx)
	repo := NewServiceCatalogRepository(pool)

	item := newServiceCatalogTestItem(t, "Interior Detail")
	if err := repo.Create(txCtx, item); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	got, found, err := repo.SetActive(txCtx, item.ID(), false)
	if err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}
	if !found {
		t.Fatal("SetActive() found = false, want true")
	}
	if got.IsActive() {
		t.Fatal("IsActive() = true, want false")
	}
}
