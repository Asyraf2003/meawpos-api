// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/core/audit"

	"github.com/google/uuid"
)

type failingAuditWriter struct{}

func (failingAuditWriter) WriteAuditEvent(context.Context, audit.Event) error {
	return errors.New("injected audit failure")
}

func TestCashSale_AuditFailureRollsBackAllAuthoritativeWrites(t *testing.T) {
	f := newR4Fixture(t)
	price := int64(18_000)
	itemID := addR4CatalogItem(t, f, "Kopi", &price)
	_, err := newR4PostSale(f, failingAuditWriter{}).Execute(context.Background(), r4SaleCommand(f, itemID, "rollback-"+uuid.NewString()))
	if err == nil {
		t.Fatal("expected injected failure")
	}
	for _, table := range []string{"sales", "sale_lines", "cash_payments", "audit_events", "financial_idempotency"} {
		var count int
		if err := f.pool.QueryRow(context.Background(), `SELECT count(*) FROM `+table+` WHERE root_id=$1`, f.rootID).Scan(&count); err != nil || count != 0 {
			t.Fatalf("%s count=%d err=%v", table, count, err)
		}
	}
}
