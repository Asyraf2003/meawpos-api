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
	"sort"
	"time"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

func (r *fakeServiceCatalogRepository) SetActive(
	_ context.Context,
	id domain.ServiceCatalogItemID,
	active bool,
) (domain.ServiceCatalogItem, bool, error) {
	item, found := r.items[id]
	if !found {
		return domain.ServiceCatalogItem{}, false, nil
	}

	if active {
		item.Activate(time.Now())
	} else {
		item.Deactivate(time.Now())
	}

	r.items[id] = item

	return item, true, nil
}

func matchesStatus(
	item domain.ServiceCatalogItem,
	status ports.ListStatusFilter,
) bool {
	switch status {
	case ports.ListStatusInactive:
		return !item.IsActive()
	case ports.ListStatusAll:
		return true
	case ports.ListStatusActive, "":
		return item.IsActive()
	default:
		return item.IsActive()
	}
}

func sortItemsByNormalizedName(items []domain.ServiceCatalogItem) {
	sort.Slice(items, func(left, right int) bool {
		return items[left].NormalizedName() < items[right].NormalizedName()
	})
}
