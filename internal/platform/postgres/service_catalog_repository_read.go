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
	"context"
	"errors"

	"pos-go/internal/modules/servicecatalog/domain"

	"github.com/jackc/pgx/v5"
)

func (r *ServiceCatalogRepository) FindByID(
	ctx context.Context,
	id domain.ServiceCatalogItemID,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE id = $1
	`, string(id))

	return scanOptionalServiceCatalogItem(row)
}

func (r *ServiceCatalogRepository) FindByNormalizedName(
	ctx context.Context,
	normalizedName domain.NormalizedName,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE normalized_name = $1
	`, string(normalizedName))

	return scanOptionalServiceCatalogItem(row)
}

func scanOptionalServiceCatalogItem(
	row serviceCatalogItemScanner,
) (domain.ServiceCatalogItem, bool, error) {
	item, err := scanServiceCatalogItem(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ServiceCatalogItem{}, false, nil
	}
	if err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	return item, true, nil
}

func (r *ServiceCatalogRepository) findMany(
	ctx context.Context,
	sql string,
	args ...any,
) ([]domain.ServiceCatalogItem, error) {
	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []domain.ServiceCatalogItem{}
	for rows.Next() {
		item, err := scanServiceCatalogItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}
