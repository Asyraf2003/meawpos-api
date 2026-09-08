// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package pricing

import (
	"errors"
	"time"

	"pos-go/internal/core/money"
)

var ErrPriceMustBePositive = errors.New("catalog price must be positive")

type Price struct {
	RootID, CatalogItemID string
	Amount                money.Money
	CreatedAt, UpdatedAt  time.Time
}

func NewPrice(rootID, itemID string, amountRupiah int64, now time.Time) (Price, error) {
	if amountRupiah <= 0 {
		return Price{}, ErrPriceMustBePositive
	}
	return Price{RootID: rootID, CatalogItemID: itemID, Amount: money.IDR(amountRupiah), CreatedAt: now, UpdatedAt: now}, nil
}
