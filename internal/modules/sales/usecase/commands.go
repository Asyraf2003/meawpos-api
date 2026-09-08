// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

type SaleItemInput struct {
	CatalogItemID string `json:"catalog_item_id"`
	Quantity      int64  `json:"quantity"`
}

type PostCashSaleCommand struct {
	RootID, ActorAccountID, SessionID, RequestID, AuthorityUsed, IdempotencyKey string
	Items                                                                       []SaleItemInput
	PaymentType                                                                 string
	TenderedRupiah                                                              int64
}

type PostCashSaleResult struct {
	SaleID   string
	Replayed bool
}

type ReverseSaleCommand struct {
	RootID, SaleID, ActorAccountID, SessionID, RequestID, AuthorityUsed, Reason string
}
