// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"

	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CatalogStore struct{ pool *pgxpool.Pool }

func NewCatalogStore(pool *pgxpool.Pool) *CatalogStore { return &CatalogStore{pool: pool} }

func (s *CatalogStore) CreateItem(ctx context.Context, item catalogcore.Item) error {
	exec := s.pool.Exec
	if tx, ok := TxFromContext(ctx); ok {
		exec = tx.Exec
	}
	_, err := exec(ctx, `INSERT INTO catalog_items (id,root_id,name,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)`, item.ID, item.RootID, item.Name, item.CreatedAt, item.UpdatedAt)
	return err
}

func (s *CatalogStore) CreatePrice(ctx context.Context, price pricing.Price) error {
	exec := s.pool.Exec
	if tx, ok := TxFromContext(ctx); ok {
		exec = tx.Exec
	}
	_, err := exec(ctx, `INSERT INTO catalog_item_prices (root_id,catalog_item_id,currency,amount_rupiah,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6)`, price.RootID, price.CatalogItemID, price.Amount.Currency(), price.Amount.AmountRupiah(), price.CreatedAt, price.UpdatedAt)
	return err
}

var _ catalogports.Store = (*CatalogStore)(nil)
