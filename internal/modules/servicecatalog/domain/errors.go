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
	ErrInvalidServiceCatalogItemID             = errors.New("service catalog item id is required")
	ErrBlankServiceCatalogItemName             = errors.New("service catalog item name is required")
	ErrInvalidServiceCatalogItemNormalizedName = errors.New("service catalog item normalized name is required")
	ErrInvalidServiceCatalogItemDefaultPrice   = errors.New("service catalog item default price must be greater than zero")
)
