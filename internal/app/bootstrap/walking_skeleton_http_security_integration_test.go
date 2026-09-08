// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestWalkingSkeletonHTTP_CrossRootAccessIsRedacted(t *testing.T) {
	app := newR4HTTPApp(t)
	admin := r4Login(t, app, "admin@example.com")
	cashier := r4Login(t, app, "kasir@example.com")
	status, body := r4Request(t, app, http.MethodPost, "/api/roots", admin, "", map[string]any{"name": "Private"})
	if status != http.StatusCreated {
		t.Fatalf("root status=%d body=%v", status, body)
	}
	rootID := dataMap(t, body)["id"].(string)
	t.Cleanup(func() { cleanupR4HTTP(t, app, rootID) })
	status, body = r4Request(t, app, http.MethodGet, "/api/roots/"+rootID+"/catalog/items/00000000-0000-0000-0000-000000000000", cashier, "", nil)
	if status != http.StatusForbidden {
		t.Fatalf("status=%d body=%v", status, body)
	}
	errorBody := body["error"].(map[string]any)
	if errorBody["code"] != "root_access_denied" {
		t.Fatalf("body=%v", body)
	}
	status, _ = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", cashier, "", map[string]any{"name": "Denied", "price_rupiah": 100})
	if status != http.StatusForbidden {
		t.Fatalf("cross-root catalog write status=%d", status)
	}
	saleBody := map[string]any{"items": []map[string]any{{"catalog_item_id": "00000000-0000-0000-0000-000000000000", "quantity": 1}}, "payment": map[string]any{"type": "cash", "tendered_rupiah": 100}}
	status, _ = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", cashier, "denied-"+uuid.NewString(), saleBody)
	if status != http.StatusForbidden {
		t.Fatalf("cross-root sale status=%d", status)
	}
	var count int
	if err := app.DB.QueryRow(context.Background(), `SELECT count(*) FROM sales WHERE root_id=$1`, rootID).Scan(&count); err != nil || count != 0 {
		t.Fatalf("sales after denial=%d err=%v", count, err)
	}
}

func cleanupR4HTTP(t *testing.T, app *App, rootID string) {
	t.Helper()
	ctx := context.Background()
	queries := []string{`DELETE FROM cash_refunds WHERE root_id=$1`, `DELETE FROM sale_reversals WHERE root_id=$1`, `DELETE FROM cash_payments WHERE root_id=$1`, `DELETE FROM sale_lines WHERE root_id=$1`, `DELETE FROM sales WHERE root_id=$1`, `DELETE FROM financial_idempotency WHERE root_id=$1`, `DELETE FROM audit_events WHERE root_id=$1`, `DELETE FROM catalog_item_prices WHERE root_id=$1`, `DELETE FROM catalog_items WHERE root_id=$1`, `DELETE FROM roots WHERE id=$1`}
	for _, query := range queries {
		if _, err := app.DB.Exec(ctx, query, rootID); err != nil {
			t.Errorf("cleanup: %v", err)
		}
	}
}
