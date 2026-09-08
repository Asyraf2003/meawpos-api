// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/sales/domain"
	salesusecase "pos-go/internal/modules/sales/usecase"

	"github.com/google/uuid"
)

func TestCashSale_FullReversalPreservesTruthAndCannotDoubleRefund(t *testing.T) {
	f := newR4Fixture(t)
	price := int64(18_000)
	itemID := addR4CatalogItem(t, f, "Kopi", &price)
	posted, err := newR4PostSale(f, NewAuditWriter(f.pool)).Execute(context.Background(), r4SaleCommand(f, itemID, "sale-"+uuid.NewString()))
	if err != nil {
		t.Fatal(err)
	}
	store := NewSalesStore(f.pool)
	reverse := salesusecase.NewReverseSale(store, NewAuditWriter(f.pool), NewTransactor(f.pool), uuid.NewString, time.Now)
	cmd := salesusecase.ReverseSaleCommand{RootID: f.rootID, SaleID: posted.SaleID, ActorAccountID: f.accountID, SessionID: f.sessionID, RequestID: uuid.NewString(), AuthorityUsed: "primary_owner", Reason: "  wrong order  "}
	reversal, err := reverse.Execute(context.Background(), cmd)
	if err != nil || reversal.Reason != "wrong order" || reversal.Refund.Amount.AmountRupiah() != 36_000 {
		t.Fatalf("reversal=%#v err=%v", reversal, err)
	}
	if _, err := reverse.Execute(context.Background(), cmd); !errors.Is(err, domain.ErrSaleAlreadyReversed) {
		t.Fatalf("second reversal error=%v", err)
	}
	sale, err := store.GetSale(context.Background(), f.rootID, posted.SaleID)
	if err != nil || sale.Status != domain.StatusReversed || sale.Reversal == nil || sale.Cash.Applied.AmountRupiah() != 36_000 {
		t.Fatalf("sale=%#v err=%v", sale, err)
	}
	var sales, payments, reversals, refunds, audits int
	err = f.pool.QueryRow(context.Background(), `SELECT (SELECT count(*) FROM sales WHERE root_id=$1),(SELECT count(*) FROM cash_payments WHERE root_id=$1),(SELECT count(*) FROM sale_reversals WHERE root_id=$1),(SELECT count(*) FROM cash_refunds WHERE root_id=$1),(SELECT count(*) FROM audit_events WHERE root_id=$1)`, f.rootID).Scan(&sales, &payments, &reversals, &refunds, &audits)
	if err != nil || sales != 1 || payments != 1 || reversals != 1 || refunds != 1 || audits != 2 {
		t.Fatalf("counts=%d,%d,%d,%d,%d err=%v", sales, payments, reversals, refunds, audits, err)
	}
}
