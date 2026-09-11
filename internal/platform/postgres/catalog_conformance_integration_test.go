// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"

	"pos-go/internal/testsupport/catalogconformance"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func TestCatalogConformance(t *testing.T) {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Fatal("DATABASE_URL is required for PostgreSQL catalog conformance")
	}
	pool, err := NewPool(context.Background(), dsn)
	if err != nil {
		t.Fatalf("open PostgreSQL conformance pool: %v", err)
	}
	t.Cleanup(pool.Close)
	catalogconformance.Run(t, postgresCatalogFactory(pool))
}

func postgresCatalogFactory(pool *pgxpool.Pool) catalogconformance.Factory {
	return func(t *testing.T) catalogconformance.Fixture {
		t.Helper()
		rootID, otherRootID := uuid.NewString(), uuid.NewString()
		accountID, otherAccountID := uuid.NewString(), uuid.NewString()
		seedPostgresCatalogRoots(t, pool, rootID, otherRootID, accountID, otherAccountID)
		t.Cleanup(func() { cleanupPostgresCatalogRoots(t, pool, rootID, otherRootID, accountID, otherAccountID) })
		store := NewCatalogStore(pool)
		return catalogconformance.Fixture{
			Store: store, Transactor: NewTransactor(pool), RootID: rootID, OtherRootID: otherRootID,
			Verify: func(ctx context.Context) error { return verifyPostgresCatalogSchema(ctx, pool) },
			Counts: func(ctx context.Context, id string) (int, int, error) { return postgresCatalogCounts(ctx, pool, id) },
		}
	}
}

func postgresCatalogCounts(ctx context.Context, pool *pgxpool.Pool, id string) (int, int, error) {
	var items, prices int
	if err := pool.QueryRow(ctx, `SELECT count(*) FROM catalog_items WHERE id=$1`, id).Scan(&items); err != nil {
		return 0, 0, err
	}
	if err := pool.QueryRow(ctx, `SELECT count(*) FROM catalog_item_prices WHERE catalog_item_id=$1`, id).Scan(&prices); err != nil {
		return 0, 0, err
	}
	return items, prices, nil
}
