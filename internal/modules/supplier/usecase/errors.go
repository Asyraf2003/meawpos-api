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
// along with gopos-api. If not, see https://www.gnu.org/licenses/.

package usecase

import "errors"

var (
	ErrSupplierNotFound             = errors.New("supplier not found")
	ErrDuplicateSupplierActiveName  = errors.New("supplier active normalized name already exists")
	ErrMissingSupplierIDGenerator   = errors.New("supplier id generator is required")
	ErrInvalidSupplierLookupLimit   = errors.New("supplier lookup limit must be between 1 and 50")
	ErrInvalidSupplierListPage      = errors.New("supplier list page must be greater than zero")
	ErrInvalidSupplierListPageLimit = errors.New("supplier list per_page must be between 1 and 50")
)
