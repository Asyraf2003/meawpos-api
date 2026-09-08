// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"

	"pos-go/internal/core/money"
	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
)

func (uc *PostCashSale) buildSale(ctx context.Context, cmd PostCashSaleCommand) (domain.Sale, error) {
	now := uc.now().UTC()
	sale := domain.Sale{ID: uc.newID(), RootID: cmd.RootID, Status: domain.StatusPosted, ActorAccountID: cmd.ActorAccountID, SessionID: cmd.SessionID, PostedAt: now, Total: money.IDR(0)}
	for _, requested := range cmd.Items {
		item, err := uc.pricing.LoadSellableItemForSale(ctx, cmd.RootID, requested.CatalogItemID)
		if err != nil {
			return domain.Sale{}, err
		}
		lineTotal, err := money.Multiply(item.UnitPrice, requested.Quantity)
		if err != nil {
			return domain.Sale{}, err
		}
		sale.Total, err = money.Add(sale.Total, lineTotal)
		if err != nil {
			return domain.Sale{}, err
		}
		sale.Lines = append(sale.Lines, domain.Line{ID: uc.newID(), RootID: cmd.RootID, SaleID: sale.ID, CatalogItemID: item.ID, ItemNameSnapshot: item.Name, UnitPrice: item.UnitPrice, Quantity: requested.Quantity, LineTotal: lineTotal, CreatedAt: now})
	}
	payment, err := cash.Settle(uc.newID(), cmd.RootID, sale.ID, sale.Total, cmd.TenderedRupiah, now)
	if err != nil {
		return domain.Sale{}, err
	}
	sale.Cash = payment
	return sale, nil
}
