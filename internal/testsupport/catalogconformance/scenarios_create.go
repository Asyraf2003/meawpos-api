// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package catalogconformance

import (
	"context"
	"testing"
	"time"

	catalogusecase "pos-go/internal/modules/catalog/usecase"
)

func scenarioSchema(t *testing.T, f Fixture) {
	if err := f.Verify(context.Background()); err != nil {
		t.Fatalf("schema/connection policy: %v", err)
	}
}

func scenarioWithoutPrice(t *testing.T, f Fixture) {
	id := newID()
	created := mustCreate(t, f, id, "  Kopi  ", nil)
	if created.Item.Name != "Kopi" || created.Price != nil {
		t.Fatalf("created=%#v", created)
	}
	read, err := f.Store.GetItem(context.Background(), f.RootID, id)
	if err != nil || !sameItem(read.Item, created.Item) || read.Price != nil {
		t.Fatalf("read=%#v err=%v", read, err)
	}
}

func scenarioExactPrice(t *testing.T, f Fixture) {
	id, amount := newID(), int64(18_500)
	created := mustCreate(t, f, id, "Kopi Susu", &amount)
	read, err := f.Store.GetItem(context.Background(), f.RootID, id)
	if err != nil {
		t.Fatal(err)
	}
	if !sameItem(read.Item, created.Item) || read.Price == nil || read.Price.Amount.AmountRupiah() != amount {
		t.Fatalf("read=%#v created=%#v", read, created)
	}
	if !read.Price.CreatedAt.Equal(proofTime) || !read.Price.UpdatedAt.Equal(proofTime) {
		t.Fatalf("price timestamps=%v/%v, want %v", read.Price.CreatedAt, read.Price.UpdatedAt, proofTime)
	}
}

func scenarioInvalidInput(t *testing.T, f Fixture) {
	blankID := newID()
	_, err := newCreate(f, blankID).Execute(context.Background(), catalogusecase.CreateItemCommand{RootID: f.RootID, Name: "  "})
	if err == nil {
		t.Fatal("blank name unexpectedly accepted")
	}
	mustCounts(t, f, blankID, 0, 0)

	priceID, invalid := newID(), int64(0)
	_, err = newCreate(f, priceID).Execute(context.Background(), catalogusecase.CreateItemCommand{RootID: f.RootID, Name: "Tea", PriceRupiah: &invalid})
	if err == nil {
		t.Fatal("zero price unexpectedly accepted")
	}
	mustCounts(t, f, priceID, 0, 0)
}

func scenarioRollback(t *testing.T, f Fixture) {
	id, amount := newID(), int64(1)
	store := failingPriceStore{Store: f.Store}
	uc := catalogusecase.NewCreateItem(store, f.Transactor, func() string { return id }, func() time.Time { return proofTime })
	_, err := uc.Execute(context.Background(), catalogusecase.CreateItemCommand{RootID: f.RootID, Name: "Rollback", PriceRupiah: &amount})
	if err == nil {
		t.Fatal("injected price failure unexpectedly succeeded")
	}
	mustCounts(t, f, id, 0, 0)
}
