// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"database/sql"

	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"
)

type CatalogStore struct{ db *sql.DB }

func NewCatalogStore(db *sql.DB) *CatalogStore { return &CatalogStore{db: db} }

type sqlExecutor interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}

func (s *CatalogStore) executor(ctx context.Context) sqlExecutor {
	if tx, ok := txFromContext(ctx); ok {
		return tx
	}
	return s.db
}

func (s *CatalogStore) CreateItem(ctx context.Context, item catalogcore.Item) error {
	id, err := canonicalUUID(item.ID)
	if err != nil {
		return err
	}
	rootID, err := canonicalUUID(item.RootID)
	if err != nil {
		return err
	}
	_, err = s.executor(ctx).ExecContext(ctx, `INSERT INTO catalog_items
		(id,root_id,name,created_at,updated_at) VALUES (?,?,?,?,?)`,
		id, rootID, item.Name, unixMicro(item.CreatedAt), unixMicro(item.UpdatedAt))
	return err
}

func (s *CatalogStore) CreatePrice(ctx context.Context, price pricing.Price) error {
	rootID, err := canonicalUUID(price.RootID)
	if err != nil {
		return err
	}
	itemID, err := canonicalUUID(price.CatalogItemID)
	if err != nil {
		return err
	}
	_, err = s.executor(ctx).ExecContext(ctx, `INSERT INTO catalog_item_prices
		(root_id,catalog_item_id,currency,amount_rupiah,created_at,updated_at)
		VALUES (?,?,?,?,?,?)`, rootID, itemID, price.Amount.Currency(),
		price.Amount.AmountRupiah(), unixMicro(price.CreatedAt), unixMicro(price.UpdatedAt))
	return err
}

var _ catalogports.Store = (*CatalogStore)(nil)
