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

	"pos-go/internal/modules/productcatalog/ports"
)

func TestProductVersionRepository_AppendAndListByProductID(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	product := newProductCatalogTestProduct(t, "Kampas Rem Versioned")
	if err := repo.Create(txCtx, product); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	changedAt := time.Date(2026, 6, 12, 10, 0, 0, 0, time.UTC)
	record := ports.ProductVersionRecord{
		ProductID:        product.ID(),
		RevisionNo:       1,
		EventName:        "product.created",
		ChangedByActorID: "actor-1",
		ChangeReason:     "integration test",
		ChangedAt:        changedAt,
	}
	if err := repo.Append(txCtx, record); err != nil {
		t.Fatalf("Append() error = %v", err)
	}

	records, err := repo.ListByProductID(txCtx, product.ID())
	if err != nil {
		t.Fatalf("ListByProductID() error = %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("ListByProductID() len = %d, want 1", len(records))
	}
	assertProductVersionRecord(t, records[0], record)
}

func TestProductVersionRepository_ListByProductIDEmpty(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	records, err := repo.ListByProductID(txCtx, "missing-product")
	if err != nil {
		t.Fatalf("ListByProductID() error = %v", err)
	}
	if len(records) != 0 {
		t.Fatalf("ListByProductID() len = %d, want 0", len(records))
	}
}
