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

type UpdateSupplier struct {
	repo ports.SupplierRepository
	now  Clock
}

type UpdateSupplierCommand struct {
	ID      string
	Name    string
	Contact SupplierContactInput
}

func NewUpdateSupplier(repo ports.SupplierRepository, now Clock) UpdateSupplier {
	return UpdateSupplier{
		repo: repo,
		now:  ensureClock(now),
	}
}

func (uc UpdateSupplier) Execute(ctx context.Context, cmd UpdateSupplierCommand) (SupplierResult, error) {
	id := domain.SupplierID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateSupplierID(id); err != nil {
		return SupplierResult{}, err
	}
	if err := domain.ValidateSupplierName(cmd.Name); err != nil {
		return SupplierResult{}, err
	}

	supplier, found, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return SupplierResult{}, err
	}
	if !found {
		return SupplierResult{}, ErrSupplierNotFound
	}

	normalizedName := domain.NormalizeName(cmd.Name)
	existing, duplicateFound, err := uc.repo.FindByNormalizedName(ctx, normalizedName)
	if err != nil {
		return SupplierResult{}, err
	}
	if supplier.IsActive() && duplicateFound && existing.ID() != supplier.ID() && existing.IsActive() {
		return SupplierResult{}, ErrDuplicateSupplierActiveName
	}

	if err := supplier.Update(cmd.Name, mapContactInput(cmd.Contact), uc.now()); err != nil {
		return SupplierResult{}, err
	}

	if err := uc.repo.Update(ctx, supplier); err != nil {
		return SupplierResult{}, err
	}

	return mapSupplierResult(supplier), nil
}
