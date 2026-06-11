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

type SessionStatusChecker struct {
	pool *pgxpool.Pool
}

func NewSessionStatusChecker(pool *pgxpool.Pool) *SessionStatusChecker {
	return &SessionStatusChecker{pool: pool}
}

func (s *SessionStatusChecker) IsSessionActive(ctx context.Context, sessionID string) (bool, error) {
	sql := `
		SELECT EXISTS (
			SELECT 1
			FROM auth_sessions
			WHERE id = $1
			  AND revoked_at IS NULL
		)
	`

	var active bool

	if tx, ok := TxFromContext(ctx); ok {
		err := tx.QueryRow(ctx, sql, sessionID).Scan(&active)
		return active, err
	}

	err := s.pool.QueryRow(ctx, sql, sessionID).Scan(&active)
	return active, err
}

var _ ports.SessionStatusChecker = (*SessionStatusChecker)(nil)
