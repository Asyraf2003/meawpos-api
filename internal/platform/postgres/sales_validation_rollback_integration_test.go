// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"errors"
	"math"
	"testing"

	"pos-go/internal/core/money"
	"pos-go/internal/modules/payment/cash"

	"github.com/google/uuid"
)

func TestCashSale_DomainValidationLeavesNoDurableMutation(t *testing.T) {
	checks := []struct {
		name          string
		price, tender int64
		want          error
	}{
		{"insufficient tender", 18_000, 1, cash.ErrInsufficientTender},
		{"money overflow", math.MaxInt64, math.MaxInt64, money.ErrOverflow},
	}
	for _, check := range checks {
		t.Run(check.name, func(t *testing.T) {
			f := newR4Fixture(t)
			itemID := addR4CatalogItem(t, f, "Kopi", &check.price)
			cmd := r4SaleCommand(f, itemID, "validation-"+uuid.NewString())
			cmd.TenderedRupiah = check.tender
			if _, err := newR4PostSale(f, NewAuditWriter(f.pool)).Execute(context.Background(), cmd); !errors.Is(err, check.want) {
				t.Fatalf("error=%v want=%v", err, check.want)
			}
			var sales, idempotency, audits int
			err := f.pool.QueryRow(context.Background(), `SELECT (SELECT count(*) FROM sales WHERE root_id=$1),(SELECT count(*) FROM financial_idempotency WHERE root_id=$1),(SELECT count(*) FROM audit_events WHERE root_id=$1)`, f.rootID).Scan(&sales, &idempotency, &audits)
			if err != nil || sales != 0 || idempotency != 0 || audits != 0 {
				t.Fatalf("counts=%d,%d,%d err=%v", sales, idempotency, audits, err)
			}
		})
	}
}
