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

	"pos-go/internal/modules/servicecatalog/domain"
)

type fakeServiceCatalogRepository struct {
	items map[domain.ServiceCatalogItemID]domain.ServiceCatalogItem
}

func newFakeServiceCatalogRepository() *fakeServiceCatalogRepository {
	return &fakeServiceCatalogRepository{
		items: make(map[domain.ServiceCatalogItemID]domain.ServiceCatalogItem),
	}
}

func (r *fakeServiceCatalogRepository) Create(
	_ context.Context,
	item domain.ServiceCatalogItem,
) error {
	r.items[item.ID()] = item
	return nil
}

func (r *fakeServiceCatalogRepository) Update(
	_ context.Context,
	item domain.ServiceCatalogItem,
) error {
	r.items[item.ID()] = item
	return nil
}

func (r *fakeServiceCatalogRepository) FindByID(
	_ context.Context,
	id domain.ServiceCatalogItemID,
) (domain.ServiceCatalogItem, bool, error) {
	item, found := r.items[id]
	return item, found, nil
}

func (r *fakeServiceCatalogRepository) FindByNormalizedName(
	_ context.Context,
	normalizedName domain.NormalizedName,
) (domain.ServiceCatalogItem, bool, error) {
	for _, item := range r.items {
		if item.NormalizedName() == normalizedName {
			return item, true, nil
		}
	}

	return domain.ServiceCatalogItem{}, false, nil
}
