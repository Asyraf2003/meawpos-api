// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package domain

import (
	"errors"
	"strings"
	"time"

	"pos-go/internal/core/money"
	"pos-go/internal/modules/payment/cash"
)

const (
	StatusPosted            = "POSTED"
	StatusReversed          = "REVERSED"
	MaxReversalReasonLength = 500
	MaxSaleItems            = 100
	MaxIdempotencyKeyLength = 200
)

var (
	ErrItemsRequired          = errors.New("sale items are required")
	ErrTooManyItems           = errors.New("too many sale items")
	ErrInvalidQuantity        = errors.New("invalid quantity")
	ErrSaleNotFound           = errors.New("sale not found")
	ErrSaleAlreadyReversed    = errors.New("sale already reversed")
	ErrReversalReasonRequired = errors.New("reversal reason is required")
	ErrReversalReasonTooLong  = errors.New("reversal reason is too long")
	ErrUnsupportedPayment     = errors.New("unsupported payment type")
	ErrIdempotencyKeyRequired = errors.New("idempotency key is required")
	ErrIdempotencyKeyTooLong  = errors.New("idempotency key is too long")
	ErrIdempotencyConflict    = errors.New("idempotency conflict")
)

type Sale struct {
	ID, RootID, Status, ActorAccountID, SessionID string
	Total                                         money.Money
	Lines                                         []Line
	Cash                                          cash.Payment
	PostedAt                                      time.Time
	Reversal                                      *Reversal
}

type Line struct {
	ID, RootID, SaleID, CatalogItemID, ItemNameSnapshot string
	UnitPrice                                           money.Money
	Quantity                                            int64
	LineTotal                                           money.Money
	CreatedAt                                           time.Time
}

type Reversal struct {
	ID, RootID, SaleID, ActorAccountID, SessionID, Reason string
	ReversedAt                                            time.Time
	Refund                                                cash.Refund
}

func NormalizeReason(reason string) (string, error) {
	reason = strings.TrimSpace(reason)
	if reason == "" {
		return "", ErrReversalReasonRequired
	}
	if len([]rune(reason)) > MaxReversalReasonLength {
		return "", ErrReversalReasonTooLong
	}
	return reason, nil
}
