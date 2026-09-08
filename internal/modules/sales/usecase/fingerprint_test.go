// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import "testing"

func TestRequestFingerprintTreatsLineOrderAsSameLogicalRequest(t *testing.T) {
	left, err := requestFingerprint(PostCashSaleCommand{Items: []SaleItemInput{{CatalogItemID: "b", Quantity: 2}, {CatalogItemID: "a", Quantity: 1}}, PaymentType: "cash", TenderedRupiah: 100})
	if err != nil {
		t.Fatal(err)
	}
	right, err := requestFingerprint(PostCashSaleCommand{Items: []SaleItemInput{{CatalogItemID: "a", Quantity: 1}, {CatalogItemID: "b", Quantity: 2}}, PaymentType: "cash", TenderedRupiah: 100})
	if err != nil || left != right {
		t.Fatalf("fingerprints %q and %q, err=%v", left, right, err)
	}
}
