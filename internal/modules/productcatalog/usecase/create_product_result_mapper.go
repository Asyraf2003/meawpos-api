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
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func createProductResultFromDomain(product *domain.Product, now time.Time) CreateProductResult {
	return CreateProductResult{
		ID:                   product.ID(),
		Code:                 product.Code(),
		Name:                 product.Name(),
		NormalizedName:       product.NormalizedName(),
		Brand:                product.Brand(),
		NormalizedBrand:      product.NormalizedBrand(),
		Size:                 product.Size(),
		SalePriceRupiah:      product.SalePriceRupiah(),
		ReorderPointQty:      product.ReorderPointQty(),
		CriticalThresholdQty: product.CriticalThresholdQty(),
		Status:               string(product.Status()),
		CreatedAt:            now,
		UpdatedAt:            now,
	}
}
