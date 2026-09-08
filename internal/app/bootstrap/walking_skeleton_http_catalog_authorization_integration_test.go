// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"context"
	"net/http"
	"testing"

	"pos-go/internal/core/authorization"

	"github.com/google/uuid"
)

func TestWalkingSkeletonHTTP_CatalogRequiresOperationPermissions(t *testing.T) {
	app := newR4HTTPApp(t)
	owner := r4Login(t, app, "admin@example.com")
	member := r4Login(t, app, "kasir@example.com")
	if _, err := app.DB.Exec(t.Context(), `UPDATE api_capabilities SET enabled=false`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := app.DB.Exec(context.Background(), `UPDATE api_capabilities SET enabled=true`); err != nil {
			t.Errorf("restore legacy capabilities: %v", err)
		}
	})
	status, body := r4Request(t, app, http.MethodPost, "/api/roots", owner, "", map[string]any{"name": "Catalog authority"})
	if status != http.StatusCreated {
		t.Fatalf("root status=%d body=%v", status, body)
	}
	rootID := dataMap(t, body)["id"].(string)
	t.Cleanup(func() { cleanupR4HTTP(t, app, rootID) })
	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", owner, "", map[string]any{"name": "Owner item", "price_rupiah": 18000})
	if status != http.StatusCreated {
		t.Fatalf("owner catalog create status=%d body=%v", status, body)
	}
	itemID := dataMap(t, body)["id"].(string)

	addRootMembership(t, app, rootID, "kasir@example.com")
	assertR4Status(t, app, http.MethodGet, "/api/roots/"+rootID+"/catalog/items/"+itemID, member, "", nil, http.StatusForbidden)
	assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", member, "", map[string]any{"name": "Denied"}, http.StatusForbidden)

	grantRootRole(t, app, rootID, "kasir@example.com", "catalog-reader", authorization.PermissionCatalogItemRead)
	assertR4Status(t, app, http.MethodGet, "/api/roots/"+rootID+"/catalog/items/"+itemID, member, "", nil, http.StatusOK)
	assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", member, "", map[string]any{"name": "Still denied"}, http.StatusForbidden)

	grantRootRole(t, app, rootID, "kasir@example.com", "catalog-creator", authorization.PermissionCatalogItemCreate)
	assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", member, "", map[string]any{"name": "Created by grant"}, http.StatusCreated)
	saleBody := map[string]any{"items": []map[string]any{{"catalog_item_id": itemID, "quantity": 1}}, "payment": map[string]any{"type": "cash", "tendered_rupiah": 20000}}
	assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", member, "catalog-only-"+uuid.NewString(), saleBody, http.StatusForbidden)

	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", owner, "owner-"+uuid.NewString(), saleBody)
	if status != http.StatusCreated {
		t.Fatalf("owner sale status=%d body=%v", status, body)
	}
	saleID := dataMap(t, body)["id"].(string)
	assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales/"+saleID+"/reversals", member, "", map[string]any{"reason": "not allowed"}, http.StatusForbidden)
}

func assertR4Status(t *testing.T, app *App, method, path, token, key string, body any, want int) {
	t.Helper()
	status, response := r4Request(t, app, method, path, token, key, body)
	if status != want {
		t.Fatalf("%s %s status=%d want=%d body=%v", method, path, status, want, response)
	}
}
