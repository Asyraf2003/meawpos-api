// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"testing"
	"time"

	coretransaction "pos-go/internal/core/transaction"
	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"
)

type storeFake struct {
	item  catalogcore.Item
	price *pricing.Price
}

func (f *storeFake) CreateItem(_ context.Context, item catalogcore.Item) error {
	f.item = item
	return nil
}
func (f *storeFake) CreatePrice(_ context.Context, price pricing.Price) error {
	f.price = &price
	return nil
}
func (f *storeFake) GetItem(context.Context, string, string) (catalogports.ItemWithPrice, error) {
	return catalogports.ItemWithPrice{}, nil
}

type txFake struct{}

func (txFake) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

var _ coretransaction.Transactor = txFake{}

func TestCreateItemAllowsCoreWithoutPrice(t *testing.T) {
	store := &storeFake{}
	uc := NewCreateItem(store, txFake{}, func() string { return "item" }, func() time.Time { return time.Unix(1, 0) })
	got, err := uc.Execute(context.Background(), CreateItemCommand{RootID: "root", Name: "  Kopi  "})
	if err != nil || got.Item.Name != "Kopi" || got.Price != nil || store.price != nil {
		t.Fatalf("Execute() = %#v, %v", got, err)
	}
}

func TestCreateItemAddsPositiveIDRPrice(t *testing.T) {
	amount := int64(18_000)
	store := &storeFake{}
	uc := NewCreateItem(store, txFake{}, func() string { return "item" }, time.Now)
	got, err := uc.Execute(context.Background(), CreateItemCommand{RootID: "root", Name: "Kopi", PriceRupiah: &amount})
	if err != nil || got.Price.Amount.AmountRupiah() != amount {
		t.Fatalf("Execute() = %#v, %v", got, err)
	}
}
