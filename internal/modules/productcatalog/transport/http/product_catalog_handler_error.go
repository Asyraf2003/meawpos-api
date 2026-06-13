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
	stdhttp "net/http"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
	httpresponse "pos-go/internal/transport/http/response"
)

func mapProductCatalogError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, ports.ErrProductNotFound):
		return httpresponse.NewHTTPError(stdhttp.StatusNotFound, "product_not_found", "product not found")
	case errors.Is(err, ports.ErrDuplicateProductCode):
		return httpresponse.NewHTTPError(stdhttp.StatusConflict, "product_code_already_exists", "product code already exists")
	case errors.Is(err, ports.ErrDuplicateProductIdentity):
		return httpresponse.NewHTTPError(stdhttp.StatusConflict, "product_identity_already_exists", "product identity already exists")
	case errors.Is(err, domain.ErrProductIDRequired),
		errors.Is(err, domain.ErrProductNameRequired),
		errors.Is(err, domain.ErrProductBrandRequired),
		errors.Is(err, domain.ErrProductSalePriceMustBePositive),
		errors.Is(err, domain.ErrProductThresholdPairRequired),
		errors.Is(err, domain.ErrProductThresholdNegative),
		errors.Is(err, domain.ErrProductCriticalAboveReorder),
		errors.Is(err, domain.ErrProductDeleteTimeRequired),
		errors.Is(err, domain.ErrProductAlreadyDeleted),
		errors.Is(err, domain.ErrProductNotDeleted):
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "product_validation_failed", err.Error())
	default:
		return httpresponse.NewHTTPError(stdhttp.StatusInternalServerError, "product_catalog_request_failed", "product catalog request failed")
	}
}
