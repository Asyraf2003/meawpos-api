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

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

type CreateServiceCatalogItem struct {
	repo  ports.ServiceCatalogRepository
	newID ServiceCatalogItemIDGenerator
	now   Clock
}

type CreateServiceCatalogItemCommand struct {
	Name               string
	DefaultPriceRupiah int64
}

func NewCreateServiceCatalogItem(
	repo ports.ServiceCatalogRepository,
	newID ServiceCatalogItemIDGenerator,
	now Clock,
) CreateServiceCatalogItem {
	return CreateServiceCatalogItem{
		repo:  repo,
		newID: newID,
		now:   ensureClock(now),
	}
}

func (uc CreateServiceCatalogItem) Execute(
	ctx context.Context,
	cmd CreateServiceCatalogItemCommand,
) (ServiceCatalogItemResult, error) {
	if uc.newID == nil {
		return ServiceCatalogItemResult{}, ErrMissingServiceCatalogItemIDGenerator
	}

	normalizedName := domain.NormalizeName(cmd.Name)
	existing, found, err := uc.repo.FindByNormalizedName(ctx, normalizedName)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if found && existing.ID() != "" {
		return ServiceCatalogItemResult{}, ErrDuplicateServiceCatalogItemNormalizedName
	}

	id, err := uc.newID()
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	item, err := domain.NewServiceCatalogItem(
		id,
		cmd.Name,
		domain.MoneyRupiah(cmd.DefaultPriceRupiah),
		uc.now(),
	)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if err := uc.repo.Create(ctx, item); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	return mapServiceCatalogItemResult(item), nil
}
