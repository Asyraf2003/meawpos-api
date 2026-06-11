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

	"pos-go/internal/modules/servicecatalog/ports"
)

type LookupServiceCatalogItems struct {
	repo ports.ServiceCatalogRepository
}

type LookupServiceCatalogItemsCommand struct {
	Query           string
	Limit           int
	IncludeInactive bool
}

func NewLookupServiceCatalogItems(repo ports.ServiceCatalogRepository) LookupServiceCatalogItems {
	return LookupServiceCatalogItems{repo: repo}
}

func (uc LookupServiceCatalogItems) Execute(
	ctx context.Context,
	cmd LookupServiceCatalogItemsCommand,
) ([]ServiceCatalogLookupResult, error) {
	limit := cmd.Limit
	if limit == 0 {
		limit = defaultLookupLimit
	}

	if limit < 1 || limit > maxLookupLimit {
		return nil, ErrInvalidLookupLimit
	}

	items, err := uc.repo.Lookup(ctx, ports.LookupServiceCatalogItemsFilter{
		Query:      cmd.Query,
		Limit:      limit,
		ActiveOnly: !cmd.IncludeInactive,
	})
	if err != nil {
		return nil, err
	}

	return mapServiceCatalogLookupResults(items), nil
}
