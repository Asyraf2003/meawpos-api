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

func TestUpdateProductReturnsDuplicateCheckerError(t *testing.T) {
	duplicateErr := errors.New("duplicate update failure")
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_002",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		&fakeUpdateProductDuplicateChecker{err: duplicateErr},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_002",
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})

	if !errors.Is(err, duplicateErr) {
		t.Fatalf("Execute() error = %v, want %v", err, duplicateErr)
	}
}
