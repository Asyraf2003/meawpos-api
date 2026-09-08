// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCashSale_PriceReadHoldsCoherentRowLock(t *testing.T) {
	f := newR4Fixture(t)
	price := int64(18_000)
	itemID := addR4CatalogItem(t, f, "Kopi", &price)
	ctx := context.Background()
	tx, err := f.pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	txCtx := contextWithTx(ctx, tx)
	item, err := NewSalesStore(f.pool).LoadSellableItemForSale(txCtx, f.rootID, itemID)
	if err != nil || item.UnitPrice.AmountRupiah() != 18_000 {
		t.Fatalf("item=%#v err=%v", item, err)
	}
	updateCtx, cancel := context.WithTimeout(ctx, 150*time.Millisecond)
	defer cancel()
	_, err = f.pool.Exec(updateCtx, `UPDATE catalog_item_prices SET amount_rupiah=20000 WHERE root_id=$1 AND catalog_item_id=$2`, f.rootID, itemID)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("blocked update error=%v", err)
	}
	if err := tx.Rollback(ctx); err != nil {
		t.Fatal(err)
	}
	if _, err := f.pool.Exec(ctx, `UPDATE catalog_item_prices SET amount_rupiah=20000 WHERE root_id=$1 AND catalog_item_id=$2`, f.rootID, itemID); err != nil {
		t.Fatalf("update after rollback: %v", err)
	}
}
