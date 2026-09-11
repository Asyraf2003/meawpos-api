// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"
	"errors"
	"time"

	"pos-go/internal/core/money"
	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"

	"github.com/jackc/pgx/v5"
)

func (s *CatalogStore) GetItem(ctx context.Context, rootID, itemID string) (catalogports.ItemWithPrice, error) {
	var item catalogcore.Item
	var amount *int64
	var priceCreated, priceUpdated *time.Time
	err := s.pool.QueryRow(ctx, `SELECT i.id,i.root_id,i.name,i.created_at,i.updated_at,
		p.amount_rupiah,p.created_at,p.updated_at
		FROM catalog_items i LEFT JOIN catalog_item_prices p ON p.root_id=i.root_id AND p.catalog_item_id=i.id
		WHERE i.root_id=$1 AND i.id=$2`, rootID, itemID).Scan(&item.ID, &item.RootID,
		&item.Name, &item.CreatedAt, &item.UpdatedAt, &amount, &priceCreated, &priceUpdated)
	if errors.Is(err, pgx.ErrNoRows) {
		return catalogports.ItemWithPrice{}, catalogports.ErrItemNotFound
	}
	if err != nil {
		return catalogports.ItemWithPrice{}, err
	}
	result := catalogports.ItemWithPrice{Item: item}
	if amount != nil {
		result.Price = &pricing.Price{RootID: rootID, CatalogItemID: itemID,
			Amount: money.IDR(*amount), CreatedAt: priceCreated.UTC(), UpdatedAt: priceUpdated.UTC()}
	}
	return result, nil
}
