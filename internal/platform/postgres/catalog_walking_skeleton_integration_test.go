// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"errors"
	"testing"
	"time"

	coretransaction "pos-go/internal/core/transaction"
	catalogports "pos-go/internal/modules/catalog/ports"
	catalogusecase "pos-go/internal/modules/catalog/usecase"
	salesports "pos-go/internal/modules/sales/ports"

	"github.com/google/uuid"
)

func TestCatalog_CoreWithoutPriceAndCrossRootIsolation(t *testing.T) {
	f := newR4Fixture(t)
	store := NewCatalogStore(f.pool)
	var _ coretransaction.Transactor = NewTransactor(f.pool)
	uc := catalogusecase.NewCreateItem(store, NewTransactor(f.pool), uuid.NewString, time.Now)
	item, err := uc.Execute(context.Background(), catalogusecase.CreateItemCommand{RootID: f.rootID, Name: "  Kopi  "})
	if err != nil || item.Item.Name != "Kopi" || item.Price != nil {
		t.Fatalf("item=%#v err=%v", item, err)
	}
	read, err := store.GetItem(context.Background(), f.rootID, item.Item.ID)
	if err != nil || read.Price != nil {
		t.Fatalf("read=%#v err=%v", read, err)
	}
	if _, err := store.GetItem(context.Background(), f.otherRootID, item.Item.ID); !errors.Is(err, catalogports.ErrItemNotFound) {
		t.Fatalf("cross-root error=%v", err)
	}
	_, err = NewSalesStore(f.pool).LoadSellableItemForSale(context.Background(), f.rootID, item.Item.ID)
	if !errors.Is(err, salesports.ErrCatalogItemNotSellable) {
		t.Fatalf("sellable error=%v", err)
	}
}

func TestCatalog_PriceCannotAttachAcrossRoots(t *testing.T) {
	f := newR4Fixture(t)
	amount := int64(18_000)
	itemID := addR4CatalogItem(t, f, "Kopi", &amount)
	_, err := f.pool.Exec(context.Background(), `INSERT INTO catalog_item_prices (root_id,catalog_item_id,currency,amount_rupiah,created_at,updated_at) VALUES ($1,$2,'IDR',100,now(),now())`, f.otherRootID, itemID)
	if err == nil {
		t.Fatal("expected cross-root price foreign key failure")
	}
}
