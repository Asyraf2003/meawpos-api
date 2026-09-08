// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/core/audit"
	"pos-go/internal/core/money"
	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
	"pos-go/internal/modules/sales/ports"
)

type saleStoreFake struct {
	claim  ports.Claim
	sale   domain.Sale
	events []string
	fail   string
}

func (f *saleStoreFake) LoadSellableItemForSale(context.Context, string, string) (ports.SellableItem, error) {
	return ports.SellableItem{ID: "item", Name: "Kopi", UnitPrice: money.IDR(18_000)}, nil
}
func (f *saleStoreFake) ClaimIdempotency(context.Context, string, string, string, string, time.Time) (ports.Claim, error) {
	f.events = append(f.events, "claim")
	return f.claim, nil
}
func (f *saleStoreFake) BindIdempotency(context.Context, string, string, string, string, string) error {
	f.events = append(f.events, "bind")
	return nil
}
func (f *saleStoreFake) InsertSale(_ context.Context, s domain.Sale) error {
	f.sale = s
	f.events = append(f.events, "sale")
	return nil
}
func (f *saleStoreFake) InsertSaleLine(context.Context, domain.Line) error {
	f.events = append(f.events, "line")
	return nil
}
func (f *saleStoreFake) InsertCashPayment(context.Context, cash.Payment) error {
	f.events = append(f.events, "payment")
	return nil
}
func (f *saleStoreFake) GetSale(context.Context, string, string) (domain.Sale, error) {
	return f.sale, nil
}

type writerFake struct {
	store *saleStoreFake
	err   error
}

func (w writerFake) WriteAuditEvent(context.Context, audit.Event) error {
	w.store.events = append(w.store.events, "audit")
	return w.err
}

type txFake struct{ rolledBack bool }

func (t *txFake) RunInTx(ctx context.Context, fn func(context.Context) error) error {
	err := fn(ctx)
	if err != nil {
		t.rolledBack = true
	}
	return err
}

func TestPostCashSaleBuildsExactAtomicIntent(t *testing.T) {
	store, tx := &saleStoreFake{}, &txFake{}
	ids := 0
	uc := NewPostCashSale(store, store, store, store, writerFake{store: store}, tx, func() string { ids++; return "id" + string(rune('0'+ids)) }, func() time.Time { return time.Unix(1, 0) })
	got, err := uc.Execute(context.Background(), PostCashSaleCommand{RootID: "root", ActorAccountID: "actor", SessionID: "session", RequestID: "request", AuthorityUsed: "roles:cashier", IdempotencyKey: "key", Items: []SaleItemInput{{CatalogItemID: "item", Quantity: 2}}, PaymentType: "cash", TenderedRupiah: 50_000})
	if err != nil || got.Replayed || store.sale.Total.AmountRupiah() != 36_000 || store.sale.Cash.Change.AmountRupiah() != 14_000 {
		t.Fatalf("Execute()=%#v,%v sale=%#v", got, err, store.sale)
	}
	want := []string{"claim", "sale", "line", "payment", "audit", "bind"}
	if len(store.events) != len(want) {
		t.Fatalf("events=%v", store.events)
	}
	for i := range want {
		if store.events[i] != want[i] {
			t.Fatalf("events=%v", store.events)
		}
	}
}
