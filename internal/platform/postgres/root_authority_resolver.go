// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"
	"errors"

	"pos-go/internal/core/authorization"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RootAuthorityResolver struct{ pool *pgxpool.Pool }

func NewRootAuthorityResolver(pool *pgxpool.Pool) *RootAuthorityResolver {
	return &RootAuthorityResolver{pool: pool}
}

func (r *RootAuthorityResolver) ResolveRootAuthority(ctx context.Context, rootID, accountID string) (authorization.Authority, error) {
	var owner bool
	err := r.pool.QueryRow(ctx, `SELECT primary_owner_account_id=$2 FROM roots WHERE id=$1`, rootID, accountID).Scan(&owner)
	if errors.Is(err, pgx.ErrNoRows) {
		return authorization.Authority{}, authorization.ErrRootAccessDenied
	}
	if err != nil {
		return authorization.Authority{}, err
	}
	roles, permissions, member, err := r.loadRootGrants(ctx, rootID, accountID)
	if err != nil {
		return authorization.Authority{}, err
	}
	if !owner && !member {
		return authorization.Authority{}, authorization.ErrRootAccessDenied
	}
	return authorization.Authority{RootID: rootID, AccountID: accountID, PrimaryOwner: owner, Roles: roles, Permissions: permissions}, nil
}

var _ authorization.Resolver = (*RootAuthorityResolver)(nil)
