// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package catalogconformance

import (
	"context"
	"errors"
	"testing"
	"time"

	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"
	catalogusecase "pos-go/internal/modules/catalog/usecase"

	"github.com/google/uuid"
)

var proofTime = time.Date(2026, 9, 11, 2, 3, 4, 123456000, time.UTC)

func newCreate(f Fixture, id string) *catalogusecase.CreateItem {
	return catalogusecase.NewCreateItem(f.Store, f.Transactor, func() string { return id }, func() time.Time { return proofTime })
}

func mustCreate(t *testing.T, f Fixture, id, name string, amount *int64) catalogports.ItemWithPrice {
	t.Helper()
	got, err := newCreate(f, id).Execute(context.Background(), catalogusecase.CreateItemCommand{
		RootID: f.RootID, Name: name, PriceRupiah: amount,
	})
	if err != nil {
		t.Fatalf("create catalog item: %v", err)
	}
	return got
}

func mustCounts(t *testing.T, f Fixture, id string, wantItems, wantPrices int) {
	t.Helper()
	items, prices, err := f.Counts(context.Background(), id)
	if err != nil {
		t.Fatalf("count durable rows: %v", err)
	}
	if items != wantItems || prices != wantPrices {
		t.Fatalf("durable rows items=%d prices=%d, want %d/%d", items, prices, wantItems, wantPrices)
	}
}

func item(id, rootID, name string) catalogcore.Item {
	got, err := catalogcore.NewItem(id, rootID, name, proofTime)
	if err != nil {
		panic(err)
	}
	return got
}

func price(rootID, itemID string, amount int64) pricing.Price {
	got, err := pricing.NewPrice(rootID, itemID, amount, proofTime)
	if err != nil {
		panic(err)
	}
	return got
}

func assertNotFound(t *testing.T, f Fixture, rootID, itemID string) {
	t.Helper()
	_, err := f.Store.GetItem(context.Background(), rootID, itemID)
	if !errors.Is(err, catalogports.ErrItemNotFound) {
		t.Fatalf("GetItem error=%v, want catalog not found", err)
	}
}

func newID() string { return uuid.NewString() }

func sameItem(left, right catalogcore.Item) bool {
	return left.ID == right.ID && left.RootID == right.RootID && left.Name == right.Name &&
		left.CreatedAt.Equal(right.CreatedAt) && left.UpdatedAt.Equal(right.UpdatedAt)
}

type failingPriceStore struct{ catalogports.Store }

func (s failingPriceStore) CreatePrice(context.Context, pricing.Price) error {
	return errors.New("injected price write failure")
}
