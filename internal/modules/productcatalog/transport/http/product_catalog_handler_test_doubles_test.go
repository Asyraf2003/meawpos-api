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

package http

import (
	"context"
	"time"

	productcatalogusecase "pos-go/internal/modules/productcatalog/usecase"
)

type fakeCreateProduct struct {
	got productcatalogusecase.CreateProductCommand
}

func (f *fakeCreateProduct) Execute(
	_ context.Context,
	cmd productcatalogusecase.CreateProductCommand,
) (productcatalogusecase.CreateProductResult, error) {
	f.got = cmd
	now := time.Date(2026, 6, 12, 10, 0, 0, 0, time.UTC)

	return productcatalogusecase.CreateProductResult{
		ID: "product-1", Code: stringPtr(cmd.Code), Name: cmd.Name, Brand: cmd.Brand,
		NormalizedName: "kampas rem", NormalizedBrand: "honda", Size: cmd.Size,
		SalePriceRupiah: cmd.SalePriceRupiah, ReorderPointQty: cmd.ReorderPointQty,
		CriticalThresholdQty: cmd.CriticalThresholdQty, Status: "active",
		CreatedAt: now, UpdatedAt: now,
	}, nil
}

type fakeListProducts struct {
	got productcatalogusecase.ListProductsQuery
}

func (f *fakeListProducts) Execute(
	_ context.Context,
	query productcatalogusecase.ListProductsQuery,
) (productcatalogusecase.ListProductsResult, error) {
	f.got = query

	return productcatalogusecase.ListProductsResult{
		Items: []productcatalogusecase.ListProductsItem{{
			ID: "product-1", Code: stringPtr("SKU-001"), Name: "Kampas Rem",
			Brand: "Honda", Size: intPtr(14), SalePriceRupiah: 40000, Status: "deleted",
		}},
	}, nil
}

type fakeGetProductDetail struct{ err error }

func (f *fakeGetProductDetail) Execute(
	_ context.Context,
	_ productcatalogusecase.GetProductDetailQuery,
) (productcatalogusecase.GetProductDetailResult, error) {
	return productcatalogusecase.GetProductDetailResult{}, f.err
}

func stringPtr(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func intPtr(value int) *int { return &value }
