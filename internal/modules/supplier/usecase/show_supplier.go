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
	"strings"

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"
)

type ShowSupplier struct {
	repo ports.SupplierRepository
}

type ShowSupplierCommand struct {
	ID string
}

func NewShowSupplier(repo ports.SupplierRepository) ShowSupplier {
	return ShowSupplier{repo: repo}
}

func (uc ShowSupplier) Execute(ctx context.Context, cmd ShowSupplierCommand) (SupplierResult, error) {
	id := domain.SupplierID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateSupplierID(id); err != nil {
		return SupplierResult{}, err
	}

	supplier, found, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return SupplierResult{}, err
	}
	if !found {
		return SupplierResult{}, ErrSupplierNotFound
	}

	return mapSupplierResult(supplier), nil
}
