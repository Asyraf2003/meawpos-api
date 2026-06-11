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

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestGetProductDetailReturnsNotFound(t *testing.T) {
	usecase := NewGetProductDetail(&getProductDetailReaderDouble{
		err: ports.ErrProductNotFound,
	})

	_, err := usecase.Execute(context.Background(), GetProductDetailQuery{
		ID: "product-404",
	})

	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("expected product not found, got %v", err)
	}
}

type getProductDetailReaderDouble struct {
	found *domain.Product
	err   error
}

func (d *getProductDetailReaderDouble) GetByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	if d.err != nil {
		return nil, d.err
	}

	return d.found, nil
}

func (d *getProductDetailReaderDouble) List(
	_ context.Context,
	_ ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	return nil, nil
}

func (d *getProductDetailReaderDouble) Lookup(
	_ context.Context,
	_ ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	return nil, nil
}
