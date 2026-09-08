// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package ports

import (
	"context"
	"errors"
	"time"

	"pos-go/internal/core/money"
	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
)

var (
	ErrCatalogItemNotFound    = errors.New("catalog item not found")
	ErrCatalogItemNotSellable = errors.New("catalog item not sellable")
)

type SellableItem struct {
	ID, Name  string
	UnitPrice money.Money
}
type Claim struct {
	Existing                bool
	Fingerprint, ResourceID string
}

type PricingReader interface {
	LoadSellableItemForSale(context.Context, string, string) (SellableItem, error)
}
type PostingStore interface {
	InsertSale(context.Context, domain.Sale) error
	InsertSaleLine(context.Context, domain.Line) error
	InsertCashPayment(context.Context, cash.Payment) error
}
type IdempotencyStore interface {
	ClaimIdempotency(context.Context, string, string, string, string, time.Time) (Claim, error)
	BindIdempotency(context.Context, string, string, string, string, string) error
}
type SaleReader interface {
	GetSale(context.Context, string, string) (domain.Sale, error)
}
type ReversalStore interface {
	LockSaleForReversal(context.Context, string, string) (domain.Sale, error)
	MarkSaleReversed(context.Context, string, string, time.Time) error
	InsertReversal(context.Context, domain.Reversal) error
	InsertCashRefund(context.Context, cash.Refund) error
}
