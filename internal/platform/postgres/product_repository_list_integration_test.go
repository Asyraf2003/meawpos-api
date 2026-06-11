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
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestProductRepository_ListActiveDefault(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	active := newProductCatalogTestProduct(t, "Aki Motor")
	deleted := newProductCatalogTestProduct(t, "Ban Dalam")

	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create() active error = %v", err)
	}
	if err := repo.Create(txCtx, deleted); err != nil {
		t.Fatalf("Create() deleted error = %v", err)
	}
	if err := deleted.SoftDelete(domain.DeleteInput{
		DeletedAt: time.Now().UTC(),
		Reason:    "integration test",
	}); err != nil {
		t.Fatalf("SoftDelete() error = %v", err)
	}
	if err := repo.Update(txCtx, deleted); err != nil {
		t.Fatalf("Update() deleted error = %v", err)
	}

	items, err := repo.List(txCtx, ports.ProductListQuery{
		Search:  "aki",
		PerPage: 10,
	})
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("len(items) = %d, want 1", len(items))
	}
	if items[0].ID != active.ID() {
		t.Fatalf("items[0].ID = %q, want %q", items[0].ID, active.ID())
	}
	if items[0].Status != "active" {
		t.Fatalf("items[0].Status = %q, want active", items[0].Status)
	}
}

func TestProductRepository_ListDeletedAndAll(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	active := newProductCatalogTestProduct(t, "Kabel Gas")
	deleted := newProductCatalogTestProduct(t, "Kabel Rem")

	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create() active error = %v", err)
	}
	if err := repo.Create(txCtx, deleted); err != nil {
		t.Fatalf("Create() deleted error = %v", err)
	}
	if err := deleted.SoftDelete(domain.DeleteInput{
		DeletedAt: time.Now().UTC(),
		Reason:    "integration test",
	}); err != nil {
		t.Fatalf("SoftDelete() error = %v", err)
	}
	if err := repo.Update(txCtx, deleted); err != nil {
		t.Fatalf("Update() deleted error = %v", err)
	}

	deletedItems, err := repo.List(txCtx, ports.ProductListQuery{
		Search:  "kabel",
		Status:  "deleted",
		PerPage: 10,
	})
	if err != nil {
		t.Fatalf("List() deleted error = %v", err)
	}
	if len(deletedItems) != 1 {
		t.Fatalf("len(deletedItems) = %d, want 1", len(deletedItems))
	}
	if deletedItems[0].ID != deleted.ID() {
		t.Fatalf("deletedItems[0].ID = %q, want %q", deletedItems[0].ID, deleted.ID())
	}
	if deletedItems[0].Status != "deleted" {
		t.Fatalf("deletedItems[0].Status = %q, want deleted", deletedItems[0].Status)
	}

	allItems, err := repo.List(txCtx, ports.ProductListQuery{
		Search:  "kabel",
		Status:  "all",
		PerPage: 10,
	})
	if err != nil {
		t.Fatalf("List() all error = %v", err)
	}
	if len(allItems) != 2 {
		t.Fatalf("len(allItems) = %d, want 2", len(allItems))
	}
}

func TestProductRepository_ListPagination(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	first := newProductCatalogTestProduct(t, "Busi Alpha")
	second := newProductCatalogTestProduct(t, "Busi Beta")

	if err := repo.Create(txCtx, first); err != nil {
		t.Fatalf("Create() first error = %v", err)
	}
	if err := repo.Create(txCtx, second); err != nil {
		t.Fatalf("Create() second error = %v", err)
	}

	items, err := repo.List(txCtx, ports.ProductListQuery{
		Search:  "busi",
		Page:    2,
		PerPage: 1,
	})
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("len(items) = %d, want 1", len(items))
	}
	if items[0].ID != second.ID() {
		t.Fatalf("items[0].ID = %q, want %q", items[0].ID, second.ID())
	}
}
