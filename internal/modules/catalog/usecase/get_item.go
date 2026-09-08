// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"

	catalogports "pos-go/internal/modules/catalog/ports"
)

type GetItem struct{ store catalogports.Store }

func NewGetItem(store catalogports.Store) *GetItem { return &GetItem{store: store} }
func (uc *GetItem) Execute(ctx context.Context, rootID, itemID string) (catalogports.ItemWithPrice, error) {
	return uc.store.GetItem(ctx, rootID, itemID)
}
