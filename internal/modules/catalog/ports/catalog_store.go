// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package ports

import (
	"context"
	"errors"

	catalogcore "pos-go/internal/modules/catalog/core"
	"pos-go/internal/modules/catalog/pricing"
)

var ErrItemNotFound = errors.New("catalog item not found")

type ItemWithPrice struct {
	Item  catalogcore.Item
	Price *pricing.Price
}

type Store interface {
	CreateItem(ctx context.Context, item catalogcore.Item) error
	CreatePrice(ctx context.Context, price pricing.Price) error
	GetItem(ctx context.Context, rootID, itemID string) (ItemWithPrice, error)
}
