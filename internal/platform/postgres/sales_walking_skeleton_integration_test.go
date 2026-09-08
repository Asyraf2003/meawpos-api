// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/sales/domain"

	"github.com/google/uuid"
)

func TestCashSale_AtomicCommitReplayConflictSnapshotAndReadback(t *testing.T) {
	f := newR4Fixture(t)
	price := int64(18_000)
	itemID := addR4CatalogItem(t, f, "Kopi Susu", &price)
	uc := newR4PostSale(f, NewAuditWriter(f.pool))
	cmd := r4SaleCommand(f, itemID, "key-"+uuid.NewString())
	first, err := uc.Execute(context.Background(), cmd)
	if err != nil || first.Replayed {
		t.Fatalf("first=%#v err=%v", first, err)
	}
	replay, err := uc.Execute(context.Background(), cmd)
	if err != nil || !replay.Replayed || replay.SaleID != first.SaleID {
		t.Fatalf("replay=%#v err=%v", replay, err)
	}
	changed := cmd
	changed.TenderedRupiah++
	if _, err := uc.Execute(context.Background(), changed); !errors.Is(err, domain.ErrIdempotencyConflict) {
		t.Fatalf("conflict error=%v", err)
	}
	_, err = f.pool.Exec(context.Background(), `UPDATE catalog_items SET name='Changed' WHERE root_id=$1 AND id=$2`, f.rootID, itemID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.pool.Exec(context.Background(), `UPDATE catalog_item_prices SET amount_rupiah=20000 WHERE root_id=$1 AND catalog_item_id=$2`, f.rootID, itemID)
	if err != nil {
		t.Fatal(err)
	}
	sale, err := NewSalesStore(f.pool).GetSale(context.Background(), f.rootID, first.SaleID)
	if err != nil || sale.Total.AmountRupiah() != 36_000 || sale.Cash.Change.AmountRupiah() != 14_000 || sale.Lines[0].ItemNameSnapshot != "Kopi Susu" || sale.Lines[0].UnitPrice.AmountRupiah() != 18_000 {
		t.Fatalf("sale=%#v err=%v", sale, err)
	}
	if _, err := NewSalesStore(f.pool).GetSale(context.Background(), f.otherRootID, first.SaleID); !errors.Is(err, domain.ErrSaleNotFound) {
		t.Fatalf("cross-root error=%v", err)
	}
	for table, want := range map[string]int{"sales": 1, "sale_lines": 1, "cash_payments": 1, "audit_events": 1, "financial_idempotency": 1} {
		var count int
		if err := f.pool.QueryRow(context.Background(), `SELECT count(*) FROM `+table+` WHERE root_id=$1`, f.rootID).Scan(&count); err != nil || count != want {
			t.Fatalf("%s count=%d err=%v", table, count, err)
		}
	}
}
