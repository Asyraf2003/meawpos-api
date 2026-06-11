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

type ListProducts struct {
	reader ports.ProductReader
}

func NewListProducts(reader ports.ProductReader) *ListProducts {
	return &ListProducts{
		reader: reader,
	}
}

func (uc *ListProducts) Execute(
	ctx context.Context,
	query ListProductsQuery,
) (ListProductsResult, error) {
	items, err := uc.reader.List(ctx, ports.ProductListQuery{
		Search:  query.Search,
		Status:  query.Status,
		Page:    query.Page,
		PerPage: query.PerPage,
	})
	if err != nil {
		return ListProductsResult{}, err
	}

	result := ListProductsResult{
		Items: make([]ListProductsItem, 0, len(items)),
	}
	for _, item := range items {
		result.Items = append(result.Items, ListProductsItem{
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
