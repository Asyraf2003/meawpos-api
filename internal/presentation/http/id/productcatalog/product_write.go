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

type ProductWriteResponse struct {
	ID                   string  `json:"id"`
	Code                 *string `json:"kode_barang"`
	Name                 string  `json:"nama_barang"`
	NormalizedName       string  `json:"nama_barang_normalized"`
	Brand                string  `json:"merek"`
	NormalizedBrand      string  `json:"merek_normalized"`
	Size                 *int    `json:"ukuran"`
	SalePriceRupiah      int64   `json:"harga_jual"`
	ReorderPointQty      *int    `json:"reorder_point_qty"`
	CriticalThresholdQty *int    `json:"critical_threshold_qty"`
	Status               string  `json:"status"`
	CreatedAt            string  `json:"created_at,omitempty"`
	UpdatedAt            string  `json:"updated_at,omitempty"`
	RevisionNo           int     `json:"revision_no,omitempty"`
}

func FromCreatedProduct(result productcatalogusecase.CreateProductResult) ProductWriteResponse {
	return ProductWriteResponse{
		ID:                   result.ID,
		Code:                 result.Code,
		Name:                 result.Name,
		NormalizedName:       result.NormalizedName,
		Brand:                result.Brand,
		NormalizedBrand:      result.NormalizedBrand,
		Size:                 result.Size,
		SalePriceRupiah:      result.SalePriceRupiah,
		ReorderPointQty:      result.ReorderPointQty,
		CriticalThresholdQty: result.CriticalThresholdQty,
		Status:               result.Status,
		CreatedAt:            formatRFC3339(result.CreatedAt),
		UpdatedAt:            formatRFC3339(result.UpdatedAt),
	}
}

func FromUpdatedProduct(result productcatalogusecase.UpdateProductResult) ProductWriteResponse {
	return ProductWriteResponse{
		ID:                   result.ID,
		Code:                 result.Code,
		Name:                 result.Name,
		NormalizedName:       result.NormalizedName,
		Brand:                result.Brand,
		NormalizedBrand:      result.NormalizedBrand,
		Size:                 result.Size,
		SalePriceRupiah:      result.SalePriceRupiah,
		ReorderPointQty:      result.ReorderPointQty,
		CriticalThresholdQty: result.CriticalThresholdQty,
		Status:               result.Status,
		UpdatedAt:            formatRFC3339(result.UpdatedAt),
		RevisionNo:           result.RevisionNo,
	}
}
