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

func TestProductVersionRepository_ListByProductIDOrdered(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	product := newProductCatalogTestProduct(t, "Busi Versioned")
	if err := repo.Create(txCtx, product); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	baseTime := time.Date(2026, 6, 12, 10, 0, 0, 0, time.UTC)
	records := []ports.ProductVersionRecord{
		newProductVersionRecord(product.ID(), 2, "product.updated", baseTime.Add(time.Minute)),
		newProductVersionRecord(product.ID(), 1, "product.created", baseTime),
	}
	for _, record := range records {
		if err := repo.Append(txCtx, record); err != nil {
			t.Fatalf("Append() error = %v", err)
		}
	}

	got, err := repo.ListByProductID(txCtx, product.ID())
	if err != nil {
		t.Fatalf("ListByProductID() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("ListByProductID() len = %d, want 2", len(got))
	}
	if got[0].RevisionNo != 1 || got[1].RevisionNo != 2 {
		t.Fatalf("Revision order = [%d %d], want [1 2]",
			got[0].RevisionNo,
			got[1].RevisionNo,
		)
	}
}
