// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/core/audit"
	"pos-go/internal/modules/sales/usecase"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type r4Fixture struct {
	pool                                                      *pgxpool.Pool
	rootID, otherRootID, accountID, otherAccountID, sessionID string
}

func newR4Fixture(t *testing.T) r4Fixture {
	t.Helper()
	ctx := context.Background()
	pool := mustOpenIntegrationPool(t, ctx)
	f := r4Fixture{pool: pool, rootID: uuid.NewString(), otherRootID: uuid.NewString(), accountID: uuid.NewString(), otherAccountID: uuid.NewString(), sessionID: uuid.NewString()}
	_, err := pool.Exec(ctx, `INSERT INTO accounts (id,email) VALUES ($1,$2),($3,$4)`, f.accountID, f.accountID+"@test.invalid", f.otherAccountID, f.otherAccountID+"@test.invalid")
	if err != nil {
		pool.Close()
		t.Fatalf("insert accounts: %v", err)
	}
	_, err = pool.Exec(ctx, `INSERT INTO roots (id,name,primary_owner_account_id,created_at,updated_at) VALUES ($1,'Root A',$2,now(),now()),($3,'Root B',$4,now(),now())`, f.rootID, f.accountID, f.otherRootID, f.otherAccountID)
	if err != nil {
		pool.Close()
		t.Fatalf("insert roots: %v", err)
	}
	_, err = pool.Exec(ctx, `INSERT INTO root_memberships (id,root_id,account_id,created_at) VALUES ($1,$2,$3,now()),($4,$5,$6,now())`, uuid.NewString(), f.rootID, f.accountID, uuid.NewString(), f.otherRootID, f.otherAccountID)
	if err != nil {
		pool.Close()
		t.Fatalf("insert memberships: %v", err)
	}
	t.Cleanup(func() { cleanupR4Fixture(t, f); pool.Close() })
	return f
}

func cleanupR4Fixture(t *testing.T, f r4Fixture) {
	t.Helper()
	ctx := context.Background()
	for _, query := range []string{
		`DELETE FROM cash_refunds WHERE root_id IN ($1,$2)`, `DELETE FROM sale_reversals WHERE root_id IN ($1,$2)`, `DELETE FROM cash_payments WHERE root_id IN ($1,$2)`, `DELETE FROM sale_lines WHERE root_id IN ($1,$2)`, `DELETE FROM sales WHERE root_id IN ($1,$2)`, `DELETE FROM financial_idempotency WHERE root_id IN ($1,$2)`, `DELETE FROM audit_events WHERE root_id IN ($1,$2)`, `DELETE FROM catalog_item_prices WHERE root_id IN ($1,$2)`, `DELETE FROM catalog_items WHERE root_id IN ($1,$2)`, `DELETE FROM roots WHERE id IN ($1,$2)`, `DELETE FROM accounts WHERE id IN ($1,$2)`} {
		if _, err := f.pool.Exec(ctx, query, f.rootID, f.otherRootID); err != nil {
			t.Errorf("cleanup %q: %v", query, err)
		}
	}
}

func newR4PostSale(f r4Fixture, writer audit.Writer) *usecase.PostCashSale {
	store := NewSalesStore(f.pool)
	return usecase.NewPostCashSale(store, store, store, store, writer, NewTransactor(f.pool), uuid.NewString, time.Now)
}

func addR4CatalogItem(t *testing.T, f r4Fixture, name string, price *int64) string {
	t.Helper()
	ctx := context.Background()
	id := uuid.NewString()
	_, err := f.pool.Exec(ctx, `INSERT INTO catalog_items (id,root_id,name,created_at,updated_at) VALUES ($1,$2,$3,now(),now())`, id, f.rootID, name)
	if err != nil {
		t.Fatal(err)
	}
	if price != nil {
		if _, err := f.pool.Exec(ctx, `INSERT INTO catalog_item_prices (root_id,catalog_item_id,currency,amount_rupiah,created_at,updated_at) VALUES ($1,$2,'IDR',$3,now(),now())`, f.rootID, id, *price); err != nil {
			t.Fatal(err)
		}
	}
	return id
}

func r4SaleCommand(f r4Fixture, itemID, key string) usecase.PostCashSaleCommand {
	return usecase.PostCashSaleCommand{RootID: f.rootID, ActorAccountID: f.accountID, SessionID: f.sessionID, RequestID: uuid.NewString(), AuthorityUsed: "primary_owner", IdempotencyKey: key, Items: []usecase.SaleItemInput{{CatalogItemID: itemID, Quantity: 2}}, PaymentType: "cash", TenderedRupiah: 50_000}
}
