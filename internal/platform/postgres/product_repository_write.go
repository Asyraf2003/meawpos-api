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

package postgres

import (
	"context"

	"pos-go/internal/modules/productcatalog/domain"
)

func (r *ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	sql := `
		INSERT INTO products (
			id,
			kode_barang,
			nama_barang,
			nama_barang_normalized,
			merek,
			merek_normalized,
			ukuran,
			harga_jual,
			reorder_point_qty,
			critical_threshold_qty,
			deleted_at,
			deleted_by_actor_id,
			delete_reason
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7,
			$8, $9, $10, $11, $12, $13
		)
	`

	_, err := r.exec(ctx, sql, productArgs(product)...)
	return err
}

func (r *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	sql := `
		UPDATE products
		SET
			kode_barang = $2,
			nama_barang = $3,
			nama_barang_normalized = $4,
			merek = $5,
			merek_normalized = $6,
			ukuran = $7,
			harga_jual = $8,
			reorder_point_qty = $9,
			critical_threshold_qty = $10,
			deleted_at = $11,
			deleted_by_actor_id = $12,
			delete_reason = $13,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.exec(ctx, sql, productArgs(product)...)
	return err
}
