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
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

func (r *fakeServiceCatalogRepository) List(
	_ context.Context,
	filter ports.ListServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	items := make([]domain.ServiceCatalogItem, 0, len(r.items))
	query := string(domain.NormalizeName(filter.Query))

	for _, item := range r.items {
		if !matchesStatus(item, filter.Status) {
			continue
		}

		if query != "" && !strings.Contains(string(item.NormalizedName()), query) {
			continue
		}

		items = append(items, item)
	}

	sortItemsByNormalizedName(items)

	return items, nil
}

func (r *fakeServiceCatalogRepository) Lookup(
	_ context.Context,
	filter ports.LookupServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	items := make([]domain.ServiceCatalogItem, 0, len(r.items))
	query := string(domain.NormalizeName(filter.Query))

	for _, item := range r.items {
		if filter.ActiveOnly && !item.IsActive() {
			continue
		}

		if query != "" && !strings.Contains(string(item.NormalizedName()), query) {
			continue
		}

		items = append(items, item)
	}

	sortItemsByNormalizedName(items)

	if filter.Limit > 0 && len(items) > filter.Limit {
		return items[:filter.Limit], nil
	}

	return items, nil
}
