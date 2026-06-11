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

package servicecatalog

import servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"

type ServiceCatalogLookupResponse struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	DefaultPriceRupiah int64  `json:"default_price_rupiah"`
}

func FromServiceCatalogLookup(
	result servicecatalogusecase.ServiceCatalogLookupResult,
) ServiceCatalogLookupResponse {
	return ServiceCatalogLookupResponse{
		ID:                 result.ID,
		Name:               result.Name,
		DefaultPriceRupiah: result.DefaultPriceRupiah,
	}
}

func FromServiceCatalogLookups(
	results []servicecatalogusecase.ServiceCatalogLookupResult,
) []ServiceCatalogLookupResponse {
	responses := make([]ServiceCatalogLookupResponse, 0, len(results))
	for _, result := range results {
		responses = append(responses, FromServiceCatalogLookup(result))
	}

	return responses
}
