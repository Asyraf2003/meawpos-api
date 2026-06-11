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

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountIdentityRepository struct {
	pool *pgxpool.Pool
}

func NewAccountIdentityRepository(pool *pgxpool.Pool) *AccountIdentityRepository {
	return &AccountIdentityRepository{pool: pool}
}

func (r *AccountIdentityRepository) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.pool.QueryRow(ctx, sql, args...)
}

func (r *AccountIdentityRepository) exec(ctx context.Context, sql string, args ...any) error {
	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, args...)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}

var _ ports.AccountIdentityRepository = (*AccountIdentityRepository)(nil)
