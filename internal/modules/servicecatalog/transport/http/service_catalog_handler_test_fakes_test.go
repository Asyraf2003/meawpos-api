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

package http

import (
	"context"

	uc "pos-go/internal/modules/servicecatalog/usecase"
)

type listFn func(context.Context, uc.ListServiceCatalogItemsCommand) ([]uc.ServiceCatalogItemResult, error)

func (f listFn) Execute(ctx context.Context, cmd uc.ListServiceCatalogItemsCommand) ([]uc.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type lookupFn func(context.Context, uc.LookupServiceCatalogItemsCommand) ([]uc.ServiceCatalogLookupResult, error)

func (f lookupFn) Execute(ctx context.Context, cmd uc.LookupServiceCatalogItemsCommand) ([]uc.ServiceCatalogLookupResult, error) {
	return f(ctx, cmd)
}

type showFn func(context.Context, uc.ShowServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error)

func (f showFn) Execute(ctx context.Context, cmd uc.ShowServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type createFn func(context.Context, uc.CreateServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error)

func (f createFn) Execute(ctx context.Context, cmd uc.CreateServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type updateFn func(context.Context, uc.UpdateServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error)

func (f updateFn) Execute(ctx context.Context, cmd uc.UpdateServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

func newTestServiceCatalogHandler() ServiceCatalogHandler {
	return NewServiceCatalogHandler(nil, nil, nil, nil, nil, nil, nil)
}
