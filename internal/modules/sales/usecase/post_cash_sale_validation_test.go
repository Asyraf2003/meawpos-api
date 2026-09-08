// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
	"pos-go/internal/modules/sales/ports"
)

func TestPostCashSaleRejectsConflictAndInsufficientTender(t *testing.T) {
	store, tx := &saleStoreFake{claim: ports.Claim{Existing: true, Fingerprint: "different", ResourceID: "sale"}}, &txFake{}
	uc := NewPostCashSale(store, store, store, store, writerFake{store: store}, tx, func() string { return "id" }, time.Now)
	cmd := PostCashSaleCommand{RootID: "root", IdempotencyKey: "key", Items: []SaleItemInput{{CatalogItemID: "item", Quantity: 1}}, PaymentType: "cash", TenderedRupiah: 1}
	if _, err := uc.Execute(context.Background(), cmd); !errors.Is(err, domain.ErrIdempotencyConflict) {
		t.Fatalf("conflict error=%v", err)
	}
	store.claim = ports.Claim{}
	if _, err := uc.Execute(context.Background(), cmd); !errors.Is(err, cash.ErrInsufficientTender) || !tx.rolledBack {
		t.Fatalf("tender error=%v rollback=%v", err, tx.rolledBack)
	}
}
