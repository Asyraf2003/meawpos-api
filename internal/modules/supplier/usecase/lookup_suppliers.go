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

import (
	"context"

	"pos-go/internal/modules/supplier/ports"
)

type LookupSuppliers struct {
	repo ports.SupplierRepository
}

type LookupSuppliersCommand struct {
	Query      string
	Limit      int
	ActiveOnly *bool
}

func NewLookupSuppliers(repo ports.SupplierRepository) LookupSuppliers {
	return LookupSuppliers{repo: repo}
}

func (uc LookupSuppliers) Execute(ctx context.Context, cmd LookupSuppliersCommand) ([]SupplierLookupResult, error) {
	limit := cmd.Limit
	if limit == 0 {
		limit = defaultLookupLimit
	}
	if limit < 1 || limit > maxLookupLimit {
		return nil, ErrInvalidSupplierLookupLimit
	}

	activeOnly := true
	if cmd.ActiveOnly != nil {
		activeOnly = *cmd.ActiveOnly
	}

	suppliers, err := uc.repo.Lookup(ctx, ports.LookupSuppliersFilter{
		Query:      cmd.Query,
		Limit:      limit,
		ActiveOnly: activeOnly,
	})
	if err != nil {
		return nil, err
	}

	return mapSupplierLookupResults(suppliers), nil
}
