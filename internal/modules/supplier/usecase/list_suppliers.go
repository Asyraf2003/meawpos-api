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

type ListSuppliers struct {
	repo ports.SupplierRepository
}

type ListSuppliersCommand struct {
	Query   string
	Status  ports.ListStatusFilter
	Page    int
	PerPage int
}

func NewListSuppliers(repo ports.SupplierRepository) ListSuppliers {
	return ListSuppliers{repo: repo}
}

func (uc ListSuppliers) Execute(ctx context.Context, cmd ListSuppliersCommand) ([]SupplierResult, error) {
	page := cmd.Page
	if page == 0 {
		page = defaultListPage
	}
	if page < 0 {
		return nil, ErrInvalidSupplierListPage
	}

	perPage := cmd.PerPage
	if perPage == 0 {
		perPage = defaultListPerPage
	}
	if perPage < 1 || perPage > maxListPerPage {
		return nil, ErrInvalidSupplierListPageLimit
	}

	status := cmd.Status
	if status == "" {
		status = ports.ListStatusActive
	}

	suppliers, err := uc.repo.List(ctx, ports.ListSuppliersFilter{
		Query:   cmd.Query,
		Status:  status,
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return nil, err
	}

	return mapSupplierResults(suppliers), nil
}
