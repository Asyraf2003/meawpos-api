// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"time"

	"pos-go/internal/modules/sales/domain"
)

type saleResponse struct {
	ID             string              `json:"id"`
	RootID         string              `json:"root_id"`
	Status         string              `json:"status"`
	TotalRupiah    int64               `json:"total_rupiah"`
	ActorAccountID string              `json:"actor_account_id"`
	SessionID      string              `json:"session_id"`
	PostedAt       time.Time           `json:"posted_at"`
	Lines          []lineResponse      `json:"lines"`
	Payment        cashPaymentResponse `json:"payment"`
	Reversal       *reversalResponse   `json:"reversal,omitempty"`
}
type lineResponse struct {
	CatalogItemID   string `json:"catalog_item_id"`
	ItemName        string `json:"item_name_snapshot"`
	UnitPriceRupiah int64  `json:"unit_price_rupiah"`
	Quantity        int64  `json:"quantity"`
	LineTotalRupiah int64  `json:"line_total_rupiah"`
}
type cashPaymentResponse struct {
	Type           string    `json:"type"`
	ID             string    `json:"id"`
	TenderedRupiah int64     `json:"tendered_rupiah"`
	AppliedRupiah  int64     `json:"applied_rupiah"`
	ChangeRupiah   int64     `json:"change_rupiah"`
	PaidAt         time.Time `json:"paid_at"`
}
type reversalResponse struct {
	ID         string             `json:"id"`
	Reason     string             `json:"reason"`
	ReversedAt time.Time          `json:"reversed_at"`
	Refund     cashRefundResponse `json:"refund"`
}
type cashRefundResponse struct {
	Type         string    `json:"type"`
	ID           string    `json:"id"`
	AmountRupiah int64     `json:"amount_rupiah"`
	RefundedAt   time.Time `json:"refunded_at"`
}

func present(sale domain.Sale) saleResponse {
	result := saleResponse{ID: sale.ID, RootID: sale.RootID, Status: sale.Status, TotalRupiah: sale.Total.AmountRupiah(), ActorAccountID: sale.ActorAccountID, SessionID: sale.SessionID, PostedAt: sale.PostedAt, Payment: cashPaymentResponse{Type: "cash", ID: sale.Cash.ID, TenderedRupiah: sale.Cash.Tendered.AmountRupiah(), AppliedRupiah: sale.Cash.Applied.AmountRupiah(), ChangeRupiah: sale.Cash.Change.AmountRupiah(), PaidAt: sale.Cash.PaidAt}, Lines: make([]lineResponse, len(sale.Lines))}
	for i, line := range sale.Lines {
		result.Lines[i] = lineResponse{CatalogItemID: line.CatalogItemID, ItemName: line.ItemNameSnapshot, UnitPriceRupiah: line.UnitPrice.AmountRupiah(), Quantity: line.Quantity, LineTotalRupiah: line.LineTotal.AmountRupiah()}
	}
	if sale.Reversal != nil {
		result.Reversal = &reversalResponse{ID: sale.Reversal.ID, Reason: sale.Reversal.Reason, ReversedAt: sale.Reversal.ReversedAt, Refund: cashRefundResponse{Type: "cash", ID: sale.Reversal.Refund.ID, AmountRupiah: sale.Reversal.Refund.Amount.AmountRupiah(), RefundedAt: sale.Reversal.Refund.RefundedAt}}
	}
	return result
}
