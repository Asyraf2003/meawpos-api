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

	"pos-go/internal/modules/productcatalog/ports"
)

type LookupProducts struct {
	reader ports.ProductReader
}

func NewLookupProducts(reader ports.ProductReader) *LookupProducts {
	return &LookupProducts{
		reader: reader,
	}
}

func (uc *LookupProducts) Execute(
	ctx context.Context,
	query LookupProductsQuery,
) (LookupProductsResult, error) {
	items, err := uc.reader.Lookup(ctx, ports.ProductLookupQuery{
		Query:          query.Query,
		Limit:          query.Limit,
		IncludeDeleted: query.IncludeDeleted,
	})
	if err != nil {
		return LookupProductsResult{}, err
	}

	result := LookupProductsResult{
		Items: make([]LookupProductsItem, 0, len(items)),
	}
	for _, item := range items {
		result.Items = append(result.Items, LookupProductsItem{
			ID:              item.ID,
			Code:            item.Code,
			Name:            item.Name,
			Brand:           item.Brand,
			Size:            item.Size,
			SalePriceRupiah: item.SalePriceRupiah,
			Status:          item.Status,
		})
	}

	return result, nil
}
