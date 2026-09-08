// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"net/http"
	"testing"

	"pos-go/internal/core/authorization"

	"github.com/google/uuid"
)

func TestWalkingSkeletonHTTP_DelegatedMultiRolePermissionsCombine(t *testing.T) {
	app := newR4HTTPApp(t)
	admin := r4Login(t, app, "admin@example.com")
	cashier := r4Login(t, app, "kasir@example.com")
	status, body := r4Request(t, app, http.MethodPost, "/api/roots", admin, "", map[string]any{"name": "Delegated"})
	if status != http.StatusCreated {
		t.Fatalf("root status=%d body=%v", status, body)
	}
	rootID := dataMap(t, body)["id"].(string)
	t.Cleanup(func() { cleanupR4HTTP(t, app, rootID) })
	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", admin, "", map[string]any{"name": "Kopi", "price_rupiah": 18000})
	if status != http.StatusCreated {
		t.Fatalf("catalog status=%d body=%v", status, body)
	}
	itemID := dataMap(t, body)["id"].(string)
	grantSplitSaleRoles(t, app, rootID)
	saleBody := map[string]any{"items": []map[string]any{{"catalog_item_id": itemID, "quantity": 1}}, "payment": map[string]any{"type": "cash", "tendered_rupiah": 20000}}
	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", cashier, "roles-"+uuid.NewString(), saleBody)
	if status != http.StatusCreated {
		t.Fatalf("delegated sale status=%d body=%v", status, body)
	}
	saleID := dataMap(t, body)["id"].(string)
	status, _ = r4Request(t, app, http.MethodGet, "/api/roots/"+rootID+"/sales/"+saleID, cashier, "", nil)
	if status != http.StatusForbidden {
		t.Fatalf("read without grant status=%d", status)
	}
}

func grantSplitSaleRoles(t *testing.T, app *App, rootID string) {
	t.Helper()
	grantRootRole(t, app, rootID, "kasir@example.com", "sales", authorization.PermissionSaleOrderCreate)
	grantRootRole(t, app, rootID, "kasir@example.com", "payments", authorization.PermissionPaymentCreate)
}
