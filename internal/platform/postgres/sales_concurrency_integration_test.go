// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"sync"
	"testing"

	salesusecase "pos-go/internal/modules/sales/usecase"

	"github.com/google/uuid"
)

func TestCashSale_ConcurrentSameKeyCreatesOneSale(t *testing.T) {
	f := newR4Fixture(t)
	price := int64(18_000)
	itemID := addR4CatalogItem(t, f, "Kopi", &price)
	uc := newR4PostSale(f, NewAuditWriter(f.pool))
	cmd := r4SaleCommand(f, itemID, "concurrent-"+uuid.NewString())
	results := make([]salesusecase.PostCashSaleResult, 2)
	errs := make([]error, 2)
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-start
			results[index], errs[index] = uc.Execute(context.Background(), cmd)
		}(i)
	}
	close(start)
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatalf("concurrent error=%v", err)
		}
	}
	if results[0].SaleID != results[1].SaleID || results[0].Replayed == results[1].Replayed {
		t.Fatalf("results=%#v", results)
	}
	var sales, idempotency, audits int
	err := f.pool.QueryRow(context.Background(), `SELECT (SELECT count(*) FROM sales WHERE root_id=$1),(SELECT count(*) FROM financial_idempotency WHERE root_id=$1),(SELECT count(*) FROM audit_events WHERE root_id=$1 AND operation='sale.cash.post')`, f.rootID).Scan(&sales, &idempotency, &audits)
	if err != nil || sales != 1 || idempotency != 1 || audits != 1 {
		t.Fatalf("counts=%d,%d,%d err=%v", sales, idempotency, audits, err)
	}
}
