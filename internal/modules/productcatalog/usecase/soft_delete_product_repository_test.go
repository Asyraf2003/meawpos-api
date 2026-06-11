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

package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestSoftDeleteProductUpdatesRepository(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	repository := &softDeleteProductRepositoryDouble{found: product}
	deletedAt := time.Date(2026, 6, 10, 10, 0, 0, 0, time.UTC)
	usecase := NewSoftDeleteProduct(
		repository,
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		func() time.Time { return deletedAt },
	)

	_, err = usecase.Execute(context.Background(), SoftDeleteProductCommand{
		ID:      "product-1",
		ActorID: "actor-1",
		Reason:  "obsolete product",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.updated == nil {
		t.Fatal("expected repository update, got nil")
	}
	if repository.updated.Status() != domain.ProductStatusDeleted {
		t.Fatalf("Status() = %v, want %v", repository.updated.Status(), domain.ProductStatusDeleted)
	}
	if repository.updated.DeletedAt() == nil || !repository.updated.DeletedAt().Equal(deletedAt) {
		t.Fatalf("DeletedAt() = %v, want %v", repository.updated.DeletedAt(), deletedAt)
	}
}
