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

type AccountRoleRemover struct {
	pool *pgxpool.Pool
}

func NewAccountRoleRemover(pool *pgxpool.Pool) *AccountRoleRemover {
	return &AccountRoleRemover{pool: pool}
}

func (r *AccountRoleRemover) RemoveRole(ctx context.Context, accountID string, roleKey string) error {
	sql := `
		DELETE FROM account_roles ar
		USING roles r
		WHERE ar.role_id = r.id
		  AND ar.account_id = $1
		  AND r.key = $2
	`

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, accountID, roleKey)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, accountID, roleKey)
	return err
}

var _ ports.AccountRoleRemover = (*AccountRoleRemover)(nil)
