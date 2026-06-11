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

type SessionRevoker struct {
	pool *pgxpool.Pool
}

func NewSessionRevoker(pool *pgxpool.Pool) *SessionRevoker {
	return &SessionRevoker{pool: pool}
}

func (r *SessionRevoker) RevokeSession(ctx context.Context, sessionID string) error {
	sql := `
		UPDATE auth_sessions
		SET revoked_at = now()
		WHERE id = $1
		  AND revoked_at IS NULL
	`

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, sessionID)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, sessionID)
	return err
}

var _ ports.SessionRevoker = (*SessionRevoker)(nil)
