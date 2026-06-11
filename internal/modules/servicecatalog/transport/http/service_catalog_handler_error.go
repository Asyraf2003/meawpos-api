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
	"errors"

	"pos-go/internal/modules/servicecatalog/domain"
	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"

	"github.com/labstack/echo/v4"
)

func mapServiceCatalogError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, servicecatalogusecase.ErrServiceCatalogItemNotFound):
		return echo.NewHTTPError(404, "service catalog item not found")
	case errors.Is(err, servicecatalogusecase.ErrDuplicateServiceCatalogItemNormalizedName):
		return echo.NewHTTPError(409, "service catalog item name already exists")
	case errors.Is(err, servicecatalogusecase.ErrInvalidLookupLimit):
		return echo.NewHTTPError(400, "lookup limit must be between 1 and 50")
	case errors.Is(err, domain.ErrInvalidServiceCatalogItemID),
		errors.Is(err, domain.ErrBlankServiceCatalogItemName),
		errors.Is(err, domain.ErrInvalidServiceCatalogItemDefaultPrice):
		return echo.NewHTTPError(400, err.Error())
	default:
		return echo.NewHTTPError(500, "service catalog request failed")
	}
}
