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

import "errors"

var (
	ErrProductIDRequired              = errors.New("product id is required")
	ErrProductNameRequired            = errors.New("product name is required")
	ErrProductBrandRequired           = errors.New("product brand is required")
	ErrProductSalePriceMustBePositive = errors.New("product sale price must be greater than zero")
	ErrProductThresholdPairRequired   = errors.New("product reorder point and critical threshold must be both null or both filled")
	ErrProductThresholdNegative       = errors.New("product threshold must be non-negative")
	ErrProductCriticalAboveReorder    = errors.New("product critical threshold must not exceed reorder point")
	ErrProductDeleteTimeRequired      = errors.New("product delete time is required")
	ErrProductAlreadyDeleted          = errors.New("product is already deleted")
	ErrProductNotDeleted              = errors.New("product is not deleted")
)
