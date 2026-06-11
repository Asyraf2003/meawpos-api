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

	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"
)

type listServiceCatalogItems interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.ListServiceCatalogItemsCommand,
	) ([]servicecatalogusecase.ServiceCatalogItemResult, error)
}

type lookupServiceCatalogItems interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.LookupServiceCatalogItemsCommand,
	) ([]servicecatalogusecase.ServiceCatalogLookupResult, error)
}

type showServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.ShowServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type createServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.CreateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type updateServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.UpdateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type activateServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.ActivateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type deactivateServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.DeactivateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}
