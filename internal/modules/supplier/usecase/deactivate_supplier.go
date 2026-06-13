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

type DeactivateSupplier struct {
	repo ports.SupplierRepository
}

type DeactivateSupplierCommand struct {
	ID     string
	Reason string
}

func NewDeactivateSupplier(repo ports.SupplierRepository) DeactivateSupplier {
	return DeactivateSupplier{repo: repo}
}

func (uc DeactivateSupplier) Execute(ctx context.Context, cmd DeactivateSupplierCommand) (SupplierResult, error) {
	_ = strings.TrimSpace(cmd.Reason)

	id := domain.SupplierID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateSupplierID(id); err != nil {
		return SupplierResult{}, err
	}

	supplier, found, err := uc.repo.SetActive(ctx, id, false)
	if err != nil {
		return SupplierResult{}, err
	}
	if !found {
		return SupplierResult{}, ErrSupplierNotFound
	}

	return mapSupplierResult(supplier), nil
}
