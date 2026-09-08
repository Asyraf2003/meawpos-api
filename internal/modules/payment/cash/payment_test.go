// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package cash

import (
	"errors"
	"testing"
	"time"

	"pos-go/internal/core/money"
)

func TestSettleCalculatesExactChange(t *testing.T) {
	payment, err := Settle("payment", "root", "sale", money.IDR(73_000), 100_000, time.Time{})
	if err != nil || payment.Applied.AmountRupiah() != 73_000 || payment.Change.AmountRupiah() != 27_000 {
		t.Fatalf("Settle() = %#v, %v", payment, err)
	}
}

func TestSettleRejectsInsufficientTender(t *testing.T) {
	_, err := Settle("payment", "root", "sale", money.IDR(73_000), 72_999, time.Time{})
	if !errors.Is(err, ErrInsufficientTender) {
		t.Fatalf("error = %v", err)
	}
}
