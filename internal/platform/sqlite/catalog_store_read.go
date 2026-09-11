// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"pos-go/internal/core/money"
	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"
)

func (s *CatalogStore) GetItem(ctx context.Context, rootID, itemID string) (catalogports.ItemWithPrice, error) {
	rootID, err := canonicalUUID(rootID)
	if err != nil {
		return catalogports.ItemWithPrice{}, catalogports.ErrItemNotFound
	}
	itemID, err = canonicalUUID(itemID)
	if err != nil {
		return catalogports.ItemWithPrice{}, catalogports.ErrItemNotFound
	}
	var item catalogcore.Item
	var itemCreated, itemUpdated int64
	var amount, priceCreated, priceUpdated sql.NullInt64
	err = s.db.QueryRowContext(ctx, `SELECT i.id,i.root_id,i.name,i.created_at,i.updated_at,
		p.amount_rupiah,p.created_at,p.updated_at
		FROM catalog_items i LEFT JOIN catalog_item_prices p
		ON p.root_id=i.root_id AND p.catalog_item_id=i.id
		WHERE i.root_id=? AND i.id=?`, rootID, itemID).Scan(
		&item.ID, &item.RootID, &item.Name, &itemCreated, &itemUpdated,
		&amount, &priceCreated, &priceUpdated)
	if errors.Is(err, sql.ErrNoRows) {
		return catalogports.ItemWithPrice{}, catalogports.ErrItemNotFound
	}
	if err != nil {
		return catalogports.ItemWithPrice{}, err
	}
	item.CreatedAt = timeFromUnixMicro(itemCreated)
	item.UpdatedAt = timeFromUnixMicro(itemUpdated)
	result := catalogports.ItemWithPrice{Item: item}
	if amount.Valid {
		result.Price = &pricing.Price{RootID: rootID, CatalogItemID: itemID,
			Amount: money.IDR(amount.Int64), CreatedAt: timeFromUnixMicro(priceCreated.Int64),
			UpdatedAt: timeFromUnixMicro(priceUpdated.Int64)}
	}
	return result, nil
}
