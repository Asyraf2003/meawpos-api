// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func seedPostgresCatalogRoots(t *testing.T, pool *pgxpool.Pool, rootID, otherRootID, accountID, otherAccountID string) {
	t.Helper()
	ctx, now := context.Background(), time.Now().UTC()
	_, err := pool.Exec(ctx, `INSERT INTO accounts (id,email,created_at,updated_at)
		VALUES ($1,$2,$3,$3),($4,$5,$3,$3)`, accountID, accountID+"@test.invalid", now, otherAccountID, otherAccountID+"@test.invalid")
	if err != nil {
		t.Fatalf("seed PostgreSQL accounts: %v", err)
	}
	_, err = pool.Exec(ctx, `INSERT INTO roots (id,name,primary_owner_account_id,created_at,updated_at)
		VALUES ($1,'Root A',$2,$3,$3),($4,'Root B',$5,$3,$3)`, rootID, accountID, now, otherRootID, otherAccountID)
	if err != nil {
		t.Fatalf("seed PostgreSQL roots: %v", err)
	}
}

func cleanupPostgresCatalogRoots(t *testing.T, pool *pgxpool.Pool, rootID, otherRootID, accountID, otherAccountID string) {
	t.Helper()
	ctx := context.Background()
	for _, statement := range []string{
		`DELETE FROM catalog_item_prices WHERE root_id IN ($1,$2)`,
		`DELETE FROM catalog_items WHERE root_id IN ($1,$2)`,
		`DELETE FROM roots WHERE id IN ($1,$2)`,
	} {
		if _, err := pool.Exec(ctx, statement, rootID, otherRootID); err != nil {
			t.Errorf("cleanup PostgreSQL conformance fixture: %v", err)
		}
	}
	if _, err := pool.Exec(ctx, `DELETE FROM accounts WHERE id IN ($1,$2)`, accountID, otherAccountID); err != nil {
		t.Errorf("cleanup PostgreSQL conformance accounts: %v", err)
	}
}

func verifyPostgresCatalogSchema(ctx context.Context, pool *pgxpool.Pool) error {
	var foreignKeys int
	err := pool.QueryRow(ctx, `SELECT count(*) FROM information_schema.table_constraints
		WHERE constraint_type='FOREIGN KEY' AND table_name IN ('catalog_items','catalog_item_prices')`).Scan(&foreignKeys)
	if err != nil {
		return err
	}
	if foreignKeys < 2 {
		return fmt.Errorf("catalog foreign keys=%d, want at least 2", foreignKeys)
	}
	return nil
}
