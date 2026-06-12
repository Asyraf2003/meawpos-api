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

package productcatalog

import productcatalogusecase "pos-go/internal/modules/productcatalog/usecase"

type ProductLifecycleResponse struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	DeletedAt  string `json:"deleted_at,omitempty"`
	RestoredAt string `json:"restored_at,omitempty"`
	RevisionNo int    `json:"revision_no"`
}

func FromDeletedProduct(result productcatalogusecase.SoftDeleteProductResult) ProductLifecycleResponse {
	return ProductLifecycleResponse{
		ID:         result.ID,
		Status:     result.Status,
		DeletedAt:  formatRFC3339(result.DeletedAt),
		RevisionNo: result.RevisionNo,
	}
}

func FromRestoredProduct(result productcatalogusecase.RestoreProductResult) ProductLifecycleResponse {
	return ProductLifecycleResponse{
		ID:         result.ID,
		Status:     result.Status,
		RestoredAt: formatRFC3339(result.RestoredAt),
		RevisionNo: result.RevisionNo,
	}
}
