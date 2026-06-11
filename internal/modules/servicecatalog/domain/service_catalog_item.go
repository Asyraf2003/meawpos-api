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

import "time"

type ServiceCatalogItemID string

type ServiceCatalogItemStatus string

const (
	ServiceCatalogItemStatusActive   ServiceCatalogItemStatus = "active"
	ServiceCatalogItemStatusInactive ServiceCatalogItemStatus = "inactive"
)

type NormalizedName string

type MoneyRupiah int64

type ServiceCatalogItem struct {
	id                 ServiceCatalogItemID
	name               string
	normalizedName     NormalizedName
	defaultPriceRupiah MoneyRupiah
	isActive           bool
	createdAt          time.Time
	updatedAt          time.Time
}

func (i ServiceCatalogItem) ID() ServiceCatalogItemID {
	return i.id
}

func (i ServiceCatalogItem) Name() string {
	return i.name
}

func (i ServiceCatalogItem) NormalizedName() NormalizedName {
	return i.normalizedName
}

func (i ServiceCatalogItem) DefaultPriceRupiah() MoneyRupiah {
	return i.defaultPriceRupiah
}

func (i ServiceCatalogItem) IsActive() bool {
	return i.isActive
}

func (i ServiceCatalogItem) Status() ServiceCatalogItemStatus {
	if i.isActive {
		return ServiceCatalogItemStatusActive
	}

	return ServiceCatalogItemStatusInactive
}

func (i ServiceCatalogItem) CreatedAt() time.Time {
	return i.createdAt
}

func (i ServiceCatalogItem) UpdatedAt() time.Time {
	return i.updatedAt
}
