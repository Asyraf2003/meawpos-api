// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"

	"pos-go/internal/modules/sales/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SalesStore struct{ pool *pgxpool.Pool }

func NewSalesStore(pool *pgxpool.Pool) *SalesStore { return &SalesStore{pool: pool} }

func (s *SalesStore) exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Exec(ctx, sql, args...)
	}
	return s.pool.Exec(ctx, sql, args...)
}
func (s *SalesStore) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return s.pool.QueryRow(ctx, sql, args...)
}
func (s *SalesStore) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Query(ctx, sql, args...)
	}
	return s.pool.Query(ctx, sql, args...)
}

var _ ports.PricingReader = (*SalesStore)(nil)
var _ ports.PostingStore = (*SalesStore)(nil)
var _ ports.IdempotencyStore = (*SalesStore)(nil)
var _ ports.SaleReader = (*SalesStore)(nil)
var _ ports.ReversalStore = (*SalesStore)(nil)
