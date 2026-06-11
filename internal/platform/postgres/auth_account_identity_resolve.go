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
)

func (r *AccountIdentityRepository) ResolveOrCreateAccountByGoogle(ctx context.Context, claims ports.OIDCClaims) (string, error) {
	accountID, err := r.findAccountIDByProviderSubject(ctx, claims.Provider, claims.Subject)
	if err == nil {
		return accountID, nil
	}
	if err != pgx.ErrNoRows {
		return "", err
	}

	accountID, err = r.findAccountIDByEmail(ctx, claims.Email)
	if err == nil {
		return accountID, nil
	}
	if err != pgx.ErrNoRows {
		return "", err
	}

	return r.createAccount(ctx, claims.Email)
}

func (r *AccountIdentityRepository) findAccountIDByProviderSubject(ctx context.Context, provider string, subject string) (string, error) {
	var accountID string
	err := r.queryRow(
		ctx,
		`SELECT account_id FROM auth_identities WHERE provider = $1 AND subject = $2`,
		provider,
		subject,
	).Scan(&accountID)
	return accountID, err
}

func (r *AccountIdentityRepository) findAccountIDByEmail(ctx context.Context, email string) (string, error) {
	var accountID string
	err := r.queryRow(
		ctx,
		`SELECT id FROM accounts WHERE email = $1 LIMIT 1`,
		email,
	).Scan(&accountID)
	return accountID, err
}

func (r *AccountIdentityRepository) createAccount(ctx context.Context, email string) (string, error) {
	var accountID string
	err := r.queryRow(
		ctx,
		`INSERT INTO accounts (email) VALUES ($1) RETURNING id`,
		email,
	).Scan(&accountID)
	return accountID, err
}
