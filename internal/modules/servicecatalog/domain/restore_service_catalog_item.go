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

func RestoreServiceCatalogItem(
	id ServiceCatalogItemID,
	name string,
	normalizedName NormalizedName,
	defaultPriceRupiah MoneyRupiah,
	isActive bool,
	createdAt time.Time,
	updatedAt time.Time,
) (ServiceCatalogItem, error) {
	id = ServiceCatalogItemID(strings.TrimSpace(string(id)))
	normalizedName = NormalizedName(strings.TrimSpace(string(normalizedName)))

	if err := ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItem{}, err
	}

	if err := ValidateServiceCatalogItemName(name); err != nil {
		return ServiceCatalogItem{}, err
	}

	if normalizedName == "" {
		return ServiceCatalogItem{}, ErrInvalidServiceCatalogItemNormalizedName
	}

	if err := ValidateServiceCatalogItemDefaultPrice(defaultPriceRupiah); err != nil {
		return ServiceCatalogItem{}, err
	}

	return ServiceCatalogItem{
		id:                 id,
		name:               normalizeDisplayName(name),
		normalizedName:     normalizedName,
		defaultPriceRupiah: defaultPriceRupiah,
		isActive:           isActive,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
	}, nil
}
