// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"time"

	coretransaction "pos-go/internal/core/transaction"
	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"
)

type CreateItemCommand struct {
	RootID, Name string
	PriceRupiah  *int64
}

type CreateItem struct {
	store catalogports.Store
	tx    coretransaction.Transactor
	newID func() string
	now   func() time.Time
}

func NewCreateItem(store catalogports.Store, tx coretransaction.Transactor, newID func() string, now func() time.Time) *CreateItem {
	return &CreateItem{store: store, tx: tx, newID: newID, now: now}
}

func (uc *CreateItem) Execute(ctx context.Context, cmd CreateItemCommand) (catalogports.ItemWithPrice, error) {
	now := uc.now().UTC()
	item, err := catalogcore.NewItem(uc.newID(), cmd.RootID, cmd.Name, now)
	if err != nil {
		return catalogports.ItemWithPrice{}, err
	}
	result := catalogports.ItemWithPrice{Item: item}
	if cmd.PriceRupiah != nil {
		price, err := pricing.NewPrice(cmd.RootID, item.ID, *cmd.PriceRupiah, now)
		if err != nil {
			return catalogports.ItemWithPrice{}, err
		}
		result.Price = &price
	}
	err = uc.tx.RunInTx(ctx, func(txCtx context.Context) error {
		if err := uc.store.CreateItem(txCtx, item); err != nil {
			return err
		}
		if result.Price != nil {
			return uc.store.CreatePrice(txCtx, *result.Price)
		}
		return nil
	})
	return result, err
}
