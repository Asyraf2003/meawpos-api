// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"

	rootdomain "pos-go/internal/core/root/domain"
	rootports "pos-go/internal/core/root/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RootStore struct{ pool *pgxpool.Pool }

func NewRootStore(pool *pgxpool.Pool) *RootStore { return &RootStore{pool: pool} }

func (s *RootStore) CreateRoot(ctx context.Context, root rootdomain.Root, membershipID string) error {
	exec := s.pool.Exec
	if tx, ok := TxFromContext(ctx); ok {
		exec = tx.Exec
	}
	if _, err := exec(ctx, `INSERT INTO roots (id,name,primary_owner_account_id,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)`, root.ID, root.Name, root.PrimaryOwnerAccountID, root.CreatedAt, root.UpdatedAt); err != nil {
		return err
	}
	_, err := exec(ctx, `INSERT INTO root_memberships (id,root_id,account_id,created_at) VALUES ($1,$2,$3,$4)`, membershipID, root.ID, root.PrimaryOwnerAccountID, root.CreatedAt)
	return err
}

var _ rootports.Store = (*RootStore)(nil)
