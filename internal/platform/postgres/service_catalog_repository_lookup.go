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
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

func (r *ServiceCatalogRepository) Lookup(
	ctx context.Context,
	filter ports.LookupServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	args := []any{}
	conditions := []string{}
	nextArg := 1

	if strings.TrimSpace(filter.Query) != "" {
		conditions = append(conditions, "normalized_name LIKE $"+itoa(nextArg))
		args = append(args, "%"+domain.NormalizeName(filter.Query)+"%")
		nextArg++
	}

	if filter.ActiveOnly {
		conditions = append(conditions, "is_active = true")
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 20
	}

	args = append(args, limit)

	sql := serviceCatalogItemSelectSQL()
	if len(conditions) > 0 {
		sql += "\n\t\tWHERE " + strings.Join(conditions, " AND ")
	}
	sql += "\n\t\tORDER BY normalized_name, id"
	sql += "\n\t\tLIMIT $" + itoa(nextArg)

	return r.findMany(ctx, sql, args...)
}
