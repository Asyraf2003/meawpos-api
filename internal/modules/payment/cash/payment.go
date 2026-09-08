// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package cash

import (
	"errors"
	"time"

	"pos-go/internal/core/money"
)

var ErrInsufficientTender = errors.New("insufficient cash tender")

type Payment struct {
	ID, RootID, SaleID        string
	Tendered, Applied, Change money.Money
	PaidAt                    time.Time
}

type Refund struct {
	ID, RootID, ReversalID, CashPaymentID string
	Amount                                money.Money
	RefundedAt                            time.Time
}

func Settle(id, rootID, saleID string, total money.Money, tenderedRupiah int64, paidAt time.Time) (Payment, error) {
	if tenderedRupiah < total.AmountRupiah() {
		return Payment{}, ErrInsufficientTender
	}
	change, err := money.Subtract(money.IDR(tenderedRupiah), total)
	if err != nil {
		return Payment{}, err
	}
	return Payment{ID: id, RootID: rootID, SaleID: saleID, Tendered: money.IDR(tenderedRupiah), Applied: total, Change: change, PaidAt: paidAt}, nil
}
