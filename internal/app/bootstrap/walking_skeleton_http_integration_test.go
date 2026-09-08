// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestWalkingSkeletonHTTP_AuthenticatedRootCatalogSaleReplayAndReversal(t *testing.T) {
	app := newR4HTTPApp(t)
	admin := r4Login(t, app, "admin@example.com")
	status, _ := r4Request(t, app, http.MethodPost, "/api/roots", "", "", map[string]any{"name": "Denied"})
	if status != http.StatusUnauthorized {
		t.Fatalf("unauth root status=%d", status)
	}
	status, body := r4Request(t, app, http.MethodPost, "/api/roots", admin, "", map[string]any{"name": "Toko R4"})
	if status != http.StatusCreated {
		t.Fatalf("root status=%d body=%v", status, body)
	}
	rootID := dataMap(t, body)["id"].(string)
	t.Cleanup(func() { cleanupR4HTTP(t, app, rootID) })
	status, body = r4Request(t, app, http.MethodGet, "/api/roots", admin, "", nil)
	if status != http.StatusOK || body["success"] != true {
		t.Fatalf("list status=%d body=%v", status, body)
	}
	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", admin, "", map[string]any{"name": "Kopi Susu", "price_rupiah": 18000})
	if status != http.StatusCreated {
		t.Fatalf("catalog status=%d body=%v", status, body)
	}
	itemID := dataMap(t, body)["id"].(string)
	status, body = r4Request(t, app, http.MethodGet, "/api/roots/"+rootID+"/catalog/items/"+itemID, admin, "", nil)
	if status != http.StatusOK || dataMap(t, body)["name"] != "Kopi Susu" {
		t.Fatalf("catalog read status=%d body=%v", status, body)
	}
	saleBody := map[string]any{"items": []map[string]any{{"catalog_item_id": itemID, "quantity": 2}}, "payment": map[string]any{"type": "cash", "tendered_rupiah": 50000}}
	status, _ = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", admin, "", saleBody)
	if status != http.StatusBadRequest {
		t.Fatalf("missing idempotency status=%d", status)
	}
	key := "http-" + uuid.NewString()
	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", admin, key, saleBody)
	if status != http.StatusCreated {
		t.Fatalf("sale status=%d body=%v", status, body)
	}
	sale := dataMap(t, body)
	saleID := sale["id"].(string)
	if sale["total_rupiah"] != float64(36000) {
		t.Fatalf("sale=%v", sale)
	}
	status, replay := r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", admin, key, saleBody)
	if status != http.StatusOK || dataMap(t, replay)["id"] != saleID {
		t.Fatalf("replay status=%d body=%v", status, replay)
	}
	conflictBody := map[string]any{"items": []map[string]any{{"catalog_item_id": itemID, "quantity": 2}}, "payment": map[string]any{"type": "cash", "tendered_rupiah": 50001}}
	status, _ = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", admin, key, conflictBody)
	if status != http.StatusConflict {
		t.Fatalf("idempotency conflict status=%d", status)
	}
	status, body = r4Request(t, app, http.MethodGet, "/api/roots/"+rootID+"/sales/"+saleID, admin, "", nil)
	if status != http.StatusOK {
		t.Fatalf("read status=%d body=%v", status, body)
	}
	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales/"+saleID+"/reversals", admin, "", map[string]any{"reason": "wrong order"})
	if status != http.StatusCreated || dataMap(t, body)["status"] != "REVERSED" {
		t.Fatalf("reverse status=%d body=%v", status, body)
	}
	status, _ = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales/"+saleID+"/reversals", admin, "", map[string]any{"reason": "again"})
	if status != http.StatusConflict {
		t.Fatalf("second reverse status=%d", status)
	}
}
