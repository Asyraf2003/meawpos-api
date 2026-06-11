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

	"pos-go/internal/modules/productcatalog/ports"
)

func TestUpdateProductReturnsNotFoundWhenProductDoesNotExist(t *testing.T) {
	uc := NewUpdateProduct(
		&fakeProductRepository{},
		&fakeProductDuplicateChecker{},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err := uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "missing_product",
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})

	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("Execute() error = %v, want %v", err, ports.ErrProductNotFound)
	}
}
