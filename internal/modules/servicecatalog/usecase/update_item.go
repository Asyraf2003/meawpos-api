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

package usecase

import (
	"context"
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

type UpdateServiceCatalogItem struct {
	repo ports.ServiceCatalogRepository
	now  Clock
}

type UpdateServiceCatalogItemCommand struct {
	ID                 string
	Name               string
	DefaultPriceRupiah int64
}

func NewUpdateServiceCatalogItem(
	repo ports.ServiceCatalogRepository,
	now Clock,
) UpdateServiceCatalogItem {
	return UpdateServiceCatalogItem{
		repo: repo,
		now:  ensureClock(now),
	}
}

func (uc UpdateServiceCatalogItem) Execute(
	ctx context.Context,
	cmd UpdateServiceCatalogItemCommand,
) (ServiceCatalogItemResult, error) {
	id := domain.ServiceCatalogItemID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if err := domain.ValidateServiceCatalogItemName(cmd.Name); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	item, found, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if !found {
		return ServiceCatalogItemResult{}, ErrServiceCatalogItemNotFound
	}

	normalizedName := domain.NormalizeName(cmd.Name)
	existing, duplicateFound, err := uc.repo.FindByNormalizedName(ctx, normalizedName)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if duplicateFound && existing.ID() != item.ID() {
		return ServiceCatalogItemResult{}, ErrDuplicateServiceCatalogItemNormalizedName
	}

	if err := item.Update(cmd.Name, domain.MoneyRupiah(cmd.DefaultPriceRupiah), uc.now()); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if err := uc.repo.Update(ctx, item); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	return mapServiceCatalogItemResult(item), nil
}
