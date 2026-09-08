// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"
	"errors"

	"pos-go/internal/core/money"
	"pos-go/internal/modules/sales/ports"

	"github.com/jackc/pgx/v5"
)

func (s *SalesStore) LoadSellableItemForSale(ctx context.Context, rootID, itemID string) (ports.SellableItem, error) {
	var name string
	err := s.queryRow(ctx, `SELECT name FROM catalog_items WHERE root_id=$1 AND id=$2 FOR SHARE`, rootID, itemID).Scan(&name)
	if errors.Is(err, pgx.ErrNoRows) {
		return ports.SellableItem{}, ports.ErrCatalogItemNotFound
	}
	if err != nil {
		return ports.SellableItem{}, err
	}
	var amount int64
	err = s.queryRow(ctx, `SELECT amount_rupiah FROM catalog_item_prices WHERE root_id=$1 AND catalog_item_id=$2 FOR SHARE`, rootID, itemID).Scan(&amount)
	if errors.Is(err, pgx.ErrNoRows) {
		return ports.SellableItem{}, ports.ErrCatalogItemNotSellable
	}
	if err != nil {
		return ports.SellableItem{}, err
	}
	return ports.SellableItem{ID: itemID, Name: name, UnitPrice: money.IDR(amount)}, nil
}
