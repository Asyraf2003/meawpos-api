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

func TestWalkingSkeletonHTTP_CashierCannotReverseUntilGrantedBothReversalPermissions(t *testing.T) {
	app, _, cashier, rootID, itemID := newR4FinancialAuthorityFixture(t, "Cashier authority")
	grantRootRole(t, app, rootID, "kasir@example.com", "cashier",
		authorization.PermissionSaleOrderCreate,
		authorization.PermissionSaleOrderRead,
		authorization.PermissionPaymentCreate,
	)
	saleID := postR4Sale(t, app, rootID, itemID, cashier)
	assertR4Status(t, app, http.MethodGet, "/api/roots/"+rootID+"/sales/"+saleID, cashier, "", nil, http.StatusOK)
	assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales/"+saleID+"/reversals", cashier, "", map[string]any{"reason": "cashier denied"}, http.StatusForbidden)

	grantRootRole(t, app, rootID, "kasir@example.com", "reversal",
		authorization.PermissionSaleOrderReverse,
		authorization.PermissionPaymentRefund,
	)
	assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales/"+saleID+"/reversals", cashier, "", map[string]any{"reason": "authorized reversal"}, http.StatusCreated)
}

func TestWalkingSkeletonHTTP_ReversalRequiresBothOperationPermissions(t *testing.T) {
	tests := []struct {
		name       string
		permission string
	}{
		{name: "missing refund", permission: authorization.PermissionSaleOrderReverse},
		{name: "missing reversal", permission: authorization.PermissionPaymentRefund},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app, owner, member, rootID, itemID := newR4FinancialAuthorityFixture(t, tc.name)
			saleID := postR4Sale(t, app, rootID, itemID, owner)
			grantRootRole(t, app, rootID, "kasir@example.com", "partial-reversal", tc.permission)
			assertR4Status(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales/"+saleID+"/reversals", member, "", map[string]any{"reason": tc.name}, http.StatusForbidden)
			var reversals, refunds int
			if err := app.DB.QueryRow(t.Context(), `SELECT count(*) FROM sale_reversals WHERE root_id=$1`, rootID).Scan(&reversals); err != nil {
				t.Fatal(err)
			}
			if err := app.DB.QueryRow(t.Context(), `SELECT count(*) FROM cash_refunds WHERE root_id=$1`, rootID).Scan(&refunds); err != nil {
				t.Fatal(err)
			}
			if reversals != 0 || refunds != 0 {
				t.Fatalf("denied reversal mutated truth: reversals=%d refunds=%d", reversals, refunds)
			}
		})
	}
}

func newR4FinancialAuthorityFixture(t *testing.T, name string) (*App, string, string, string, string) {
	t.Helper()
	app := newR4HTTPApp(t)
	owner := r4Login(t, app, "admin@example.com")
	member := r4Login(t, app, "kasir@example.com")
	status, body := r4Request(t, app, http.MethodPost, "/api/roots", owner, "", map[string]any{"name": name})
	if status != http.StatusCreated {
		t.Fatalf("root status=%d body=%v", status, body)
	}
	rootID := dataMap(t, body)["id"].(string)
	t.Cleanup(func() { cleanupR4HTTP(t, app, rootID) })
	status, body = r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", owner, "", map[string]any{"name": "Kopi", "price_rupiah": 18000})
	if status != http.StatusCreated {
		t.Fatalf("catalog status=%d body=%v", status, body)
	}
	return app, owner, member, rootID, dataMap(t, body)["id"].(string)
}

func postR4Sale(t *testing.T, app *App, rootID, itemID, token string) string {
	t.Helper()
	body := map[string]any{"items": []map[string]any{{"catalog_item_id": itemID, "quantity": 1}}, "payment": map[string]any{"type": "cash", "tendered_rupiah": 20000}}
	status, response := r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/sales", token, "financial-auth-"+uuid.NewString(), body)
	if status != http.StatusCreated {
		t.Fatalf("sale status=%d body=%v", status, response)
	}
	return dataMap(t, response)["id"].(string)
}
