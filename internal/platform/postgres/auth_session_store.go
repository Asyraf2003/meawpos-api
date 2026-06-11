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
	"encoding/json"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionStore struct {
	pool *pgxpool.Pool
}

func NewSessionStore(pool *pgxpool.Pool) *SessionStore {
	return &SessionStore{pool: pool}
}

func (s *SessionStore) Create(ctx context.Context, session domain.Session) error {
	metaJSON, err := json.Marshal(session.Meta)
	if err != nil {
		return err
	}

	sql := `
		INSERT INTO auth_sessions (
			id, account_id, refresh_token_hash, expires_at, revoked_at, meta_json, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	args := []any{
		session.ID,
		session.AccountID,
		session.RefreshTokenHash,
		session.ExpiresAt,
		session.RevokedAt,
		metaJSON,
		session.CreatedAt,
	}

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, args...)
		return err
	}

	_, err = s.pool.Exec(ctx, sql, args...)
	return err
}

var _ ports.SessionStore = (*SessionStore)(nil)
