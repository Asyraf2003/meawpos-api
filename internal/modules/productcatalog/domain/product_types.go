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

type ProductStatus string

const (
	ProductStatusActive  ProductStatus = "active"
	ProductStatusDeleted ProductStatus = "deleted"
)

type ProductInput struct {
	ID                   string
	Code                 string
	Name                 string
	Brand                string
	Size                 *int
	SalePriceRupiah      int64
	ReorderPointQty      *int
	CriticalThresholdQty *int
}

type DeleteInput struct {
	DeletedAt        time.Time
	DeletedByActorID string
	Reason           string
}

type Product struct {
	id                   string
	code                 *string
	name                 string
	normalizedName       string
	brand                string
	normalizedBrand      string
	size                 *int
	salePriceRupiah      int64
	reorderPointQty      *int
	criticalThresholdQty *int
	deletedAt            *time.Time
	deletedByActorID     string
	deleteReason         string
}

func IntPtr(value int) *int {
	return &value
}
