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
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestSoftDeleteProductRejectsAlreadyDeletedProduct(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	deletedAt := time.Date(2026, 6, 10, 9, 0, 0, 0, time.UTC)
	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: "actor-1",
		Reason:           "already deleted setup",
	}); err != nil {
		t.Fatalf("SoftDelete() setup error = %v", err)
	}

	usecase := NewSoftDeleteProduct(
		&softDeleteProductRepositoryDouble{found: product},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		func() time.Time { return time.Date(2026, 6, 10, 10, 0, 0, 0, time.UTC) },
	)

	_, err = usecase.Execute(context.Background(), SoftDeleteProductCommand{
		ID:      "product-1",
		ActorID: "actor-2",
		Reason:  "delete already deleted product",
	})

	if !errors.Is(err, domain.ErrProductAlreadyDeleted) {
		t.Fatalf("expected already deleted error, got %v", err)
	}
}
