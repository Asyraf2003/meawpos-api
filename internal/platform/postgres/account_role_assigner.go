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

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRoleAssigner struct {
	pool *pgxpool.Pool
}

func NewAccountRoleAssigner(pool *pgxpool.Pool) *AccountRoleAssigner {
	return &AccountRoleAssigner{pool: pool}
}

func (a *AccountRoleAssigner) EnsureRole(ctx context.Context, accountID string, roleKey string) error {
	sql := `
		INSERT INTO account_roles (account_id, role_id)
		SELECT $1, r.id
		FROM roles r
		WHERE r.key = $2
		ON CONFLICT (account_id, role_id) DO NOTHING
	`

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, accountID, roleKey)
		return err
	}

	_, err := a.pool.Exec(ctx, sql, accountID, roleKey)
	return err
}

var _ ports.AccountRoleAssigner = (*AccountRoleAssigner)(nil)
