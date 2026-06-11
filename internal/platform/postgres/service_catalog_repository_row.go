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

package postgres

import (
	"time"

	"pos-go/internal/modules/servicecatalog/domain"
)

type serviceCatalogItemScanner interface {
	Scan(dest ...any) error
}

func serviceCatalogItemSelectSQL() string {
	return `
		SELECT
			id,
			name,
			normalized_name,
			default_price_rupiah,
			is_active,
			created_at,
			updated_at
		FROM service_catalog_items
	`
}

func scanServiceCatalogItem(row serviceCatalogItemScanner) (domain.ServiceCatalogItem, error) {
	var id string
	var name string
	var normalizedName string
	var defaultPriceRupiah int64
	var isActive bool
	var createdAt time.Time
	var updatedAt time.Time

	err := row.Scan(
		&id,
		&name,
		&normalizedName,
		&defaultPriceRupiah,
		&isActive,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return domain.ServiceCatalogItem{}, err
	}

	return domain.RestoreServiceCatalogItem(
		domain.ServiceCatalogItemID(id),
		name,
		domain.NormalizedName(normalizedName),
		domain.MoneyRupiah(defaultPriceRupiah),
		isActive,
		createdAt,
		updatedAt,
	)
}

func serviceCatalogItemArgs(item domain.ServiceCatalogItem) []any {
	return []any{
		string(item.ID()),
		item.Name(),
		string(item.NormalizedName()),
		int64(item.DefaultPriceRupiah()),
		item.IsActive(),
		item.CreatedAt(),
		item.UpdatedAt(),
	}
}
