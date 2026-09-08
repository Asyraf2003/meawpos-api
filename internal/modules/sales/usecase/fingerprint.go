// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
)

func requestFingerprint(cmd PostCashSaleCommand) (string, error) {
	items := append([]SaleItemInput(nil), cmd.Items...)
	sort.Slice(items, func(i, j int) bool {
		if items[i].CatalogItemID == items[j].CatalogItemID {
			return items[i].Quantity < items[j].Quantity
		}
		return items[i].CatalogItemID < items[j].CatalogItemID
	})
	payload := struct {
		Items          []SaleItemInput `json:"items"`
		PaymentType    string          `json:"payment_type"`
		TenderedRupiah int64           `json:"tendered_rupiah"`
	}{items, cmd.PaymentType, cmd.TenderedRupiah}
	encoded, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(encoded)
	return hex.EncodeToString(sum[:]), nil
}
