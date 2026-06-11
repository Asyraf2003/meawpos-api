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

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Code() *string {
	if p.code == nil {
		return nil
	}
	code := *p.code
	return &code
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) NormalizedName() string {
	return p.normalizedName
}

func (p *Product) Brand() string {
	return p.brand
}

func (p *Product) NormalizedBrand() string {
	return p.normalizedBrand
}

func (p *Product) Size() *int {
	return copyIntPtr(p.size)
}

func (p *Product) SalePriceRupiah() int64 {
	return p.salePriceRupiah
}

func (p *Product) ReorderPointQty() *int {
	return copyIntPtr(p.reorderPointQty)
}

func (p *Product) CriticalThresholdQty() *int {
	return copyIntPtr(p.criticalThresholdQty)
}

func (p *Product) Status() ProductStatus {
	if p.deletedAt != nil {
		return ProductStatusDeleted
	}
	return ProductStatusActive
}

func (p *Product) DeletedAt() *time.Time {
	if p.deletedAt == nil {
		return nil
	}
	deletedAt := *p.deletedAt
	return &deletedAt
}

func (p *Product) DeletedByActorID() string {
	return p.deletedByActorID
}

func (p *Product) DeleteReason() string {
	return p.deleteReason
}
