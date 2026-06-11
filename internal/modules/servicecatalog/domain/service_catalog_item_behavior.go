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

package domain

import (
	"strings"
	"time"
)

func NewServiceCatalogItem(
	id ServiceCatalogItemID,
	name string,
	defaultPriceRupiah MoneyRupiah,
	now time.Time,
) (ServiceCatalogItem, error) {
	id = ServiceCatalogItemID(strings.TrimSpace(string(id)))

	if err := ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItem{}, err
	}

	if err := ValidateServiceCatalogItemName(name); err != nil {
		return ServiceCatalogItem{}, err
	}

	if err := ValidateServiceCatalogItemDefaultPrice(defaultPriceRupiah); err != nil {
		return ServiceCatalogItem{}, err
	}

	return ServiceCatalogItem{
		id:                 id,
		name:               normalizeDisplayName(name),
		normalizedName:     NormalizeName(name),
		defaultPriceRupiah: defaultPriceRupiah,
		isActive:           true,
		createdAt:          now,
		updatedAt:          now,
	}, nil
}

func (i *ServiceCatalogItem) Update(
	name string,
	defaultPriceRupiah MoneyRupiah,
	now time.Time,
) error {
	if err := ValidateServiceCatalogItemName(name); err != nil {
		return err
	}

	if err := ValidateServiceCatalogItemDefaultPrice(defaultPriceRupiah); err != nil {
		return err
	}

	i.name = normalizeDisplayName(name)
	i.normalizedName = NormalizeName(name)
	i.defaultPriceRupiah = defaultPriceRupiah
	i.updatedAt = now

	return nil
}

func (i *ServiceCatalogItem) Activate(now time.Time) {
	i.isActive = true
	i.updatedAt = now
}

func (i *ServiceCatalogItem) Deactivate(now time.Time) {
	i.isActive = false
	i.updatedAt = now
}
