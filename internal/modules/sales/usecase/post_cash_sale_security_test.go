// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"pos-go/internal/modules/sales/domain"
)

func TestPostCashSaleRejectsAbusiveCardinalityBeforeTransaction(t *testing.T) {
	store, tx := &saleStoreFake{}, &txFake{}
	uc := NewPostCashSale(store, store, store, store, writerFake{store: store}, tx, func() string { return "id" }, time.Now)
	items := make([]SaleItemInput, domain.MaxSaleItems+1)
	for i := range items {
		items[i] = SaleItemInput{CatalogItemID: "item", Quantity: 1}
	}
	_, err := uc.Execute(context.Background(), PostCashSaleCommand{IdempotencyKey: "key", Items: items, PaymentType: "cash"})
	if !errors.Is(err, domain.ErrTooManyItems) || tx.rolledBack {
		t.Fatalf("Execute() error = %v, rollback = %v", err, tx.rolledBack)
	}
}

func TestPostCashSaleRejectsOversizedIdempotencyKeyBeforeTransaction(t *testing.T) {
	store, tx := &saleStoreFake{}, &txFake{}
	uc := NewPostCashSale(store, store, store, store, writerFake{store: store}, tx, func() string { return "id" }, time.Now)
	cmd := PostCashSaleCommand{IdempotencyKey: strings.Repeat("k", domain.MaxIdempotencyKeyLength+1), Items: []SaleItemInput{{CatalogItemID: "item", Quantity: 1}}, PaymentType: "cash"}
	_, err := uc.Execute(context.Background(), cmd)
	if !errors.Is(err, domain.ErrIdempotencyKeyTooLong) || tx.rolledBack {
		t.Fatalf("Execute() error = %v, rollback = %v", err, tx.rolledBack)
	}
}
