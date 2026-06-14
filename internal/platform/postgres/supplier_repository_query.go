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

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *SupplierRepository) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Query(ctx, sql, args...)
	}

	return r.pool.Query(ctx, sql, args...)
}

func (r *SupplierRepository) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}

	return r.pool.QueryRow(ctx, sql, args...)
}

func (r *SupplierRepository) exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Exec(ctx, sql, args...)
	}

	return r.pool.Exec(ctx, sql, args...)
}

func (r *SupplierRepository) FindByID(
	ctx context.Context,
	id domain.SupplierID,
) (domain.Supplier, bool, error) {
	return domain.Supplier{}, false, errSupplierRepositoryNotImplemented
}

func (r *SupplierRepository) FindByNormalizedName(
	ctx context.Context,
	normalizedName domain.NormalizedName,
) (domain.Supplier, bool, error) {
	return domain.Supplier{}, false, errSupplierRepositoryNotImplemented
}

func (r *SupplierRepository) FindActiveByNormalizedName(
	ctx context.Context,
	normalizedName domain.NormalizedName,
) (domain.Supplier, bool, error) {
	return domain.Supplier{}, false, errSupplierRepositoryNotImplemented
}

func (r *SupplierRepository) List(
	ctx context.Context,
	filter ports.ListSuppliersFilter,
) ([]domain.Supplier, error) {
	return nil, errSupplierRepositoryNotImplemented
}

func (r *SupplierRepository) Lookup(
	ctx context.Context,
	filter ports.LookupSuppliersFilter,
) ([]domain.Supplier, error) {
	return nil, errSupplierRepositoryNotImplemented
}
