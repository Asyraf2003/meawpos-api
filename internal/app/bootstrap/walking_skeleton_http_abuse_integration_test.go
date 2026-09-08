// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"pos-go/internal/modules/sales/domain"
	httprequest "pos-go/internal/transport/http/request"
)

func TestWalkingSkeletonHTTP_RejectsUnknownAndOversizedJSONWithoutMutation(t *testing.T) {
	app, owner, _, rootID, _ := newR4FinancialAuthorityFixture(t, "Abuse boundary")
	path := "/api/roots/" + rootID + "/catalog/items"
	status, _ := r4RawRequest(t, app, http.MethodPost, path, owner, "", "application/json", `{"name":"Unknown","is_admin":true}`)
	if status != http.StatusBadRequest {
		t.Fatalf("unknown field status = %d", status)
	}
	body := `{"name":"` + strings.Repeat("x", int(httprequest.MaxJSONBodyBytes)) + `"}`
	status, _ = r4RawRequest(t, app, http.MethodPost, path, owner, "", "application/json", body)
	if status != http.StatusRequestEntityTooLarge {
		t.Fatalf("oversized body status = %d", status)
	}
	var count int
	if err := app.DB.QueryRow(context.Background(), `SELECT count(*) FROM catalog_items WHERE root_id=$1`, rootID).Scan(&count); err != nil || count != 1 {
		t.Fatalf("catalog count after rejected requests = %d, err = %v", count, err)
	}
}

func TestWalkingSkeletonHTTP_RejectsAbusiveSaleBoundsWithoutMutation(t *testing.T) {
	app, owner, _, rootID, itemID := newR4FinancialAuthorityFixture(t, "Sale bounds")
	item := fmt.Sprintf(`{"catalog_item_id":%q,"quantity":1}`, itemID)
	body := `{"items":[` + strings.TrimSuffix(strings.Repeat(item+",", domain.MaxSaleItems+1), ",") + `],"payment":{"type":"cash","tendered_rupiah":9999999}}`
	path := "/api/roots/" + rootID + "/sales"
	status, _ := r4RawRequest(t, app, http.MethodPost, path, owner, "bounded", "application/json", body)
	if status != http.StatusBadRequest {
		t.Fatalf("too many items status = %d", status)
	}
	status, _ = r4RawRequest(t, app, http.MethodPost, path, owner, strings.Repeat("k", domain.MaxIdempotencyKeyLength+1), "application/json", strings.Replace(body, strings.TrimSuffix(strings.Repeat(item+",", domain.MaxSaleItems+1), ","), item, 1))
	if status != http.StatusBadRequest {
		t.Fatalf("oversized idempotency key status = %d", status)
	}
	var count int
	if err := app.DB.QueryRow(context.Background(), `SELECT count(*) FROM sales WHERE root_id=$1`, rootID).Scan(&count); err != nil || count != 0 {
		t.Fatalf("sales after rejected requests = %d, err = %v", count, err)
	}
}

func TestWalkingSkeletonHTTP_TreatsInjectionLikeCatalogNameAsData(t *testing.T) {
	app, owner, _, rootID, _ := newR4FinancialAuthorityFixture(t, "Injection data")
	name := `<script>alert(1)</script>'); DELETE FROM roots; --`
	status, body := r4Request(t, app, http.MethodPost, "/api/roots/"+rootID+"/catalog/items", owner, "", map[string]any{"name": name})
	if status != http.StatusCreated || dataMap(t, body)["name"] != name {
		t.Fatalf("status = %d, body = %v", status, body)
	}
	var count int
	if err := app.DB.QueryRow(context.Background(), `SELECT count(*) FROM roots WHERE id=$1`, rootID).Scan(&count); err != nil || count != 1 {
		t.Fatalf("root count = %d, err = %v", count, err)
	}
}
