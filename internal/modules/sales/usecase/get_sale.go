// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"

	"pos-go/internal/modules/sales/domain"
	"pos-go/internal/modules/sales/ports"
)

type GetSale struct{ reader ports.SaleReader }

func NewGetSale(reader ports.SaleReader) *GetSale { return &GetSale{reader: reader} }
func (uc *GetSale) Execute(ctx context.Context, rootID, saleID string) (domain.Sale, error) {
	return uc.reader.GetSale(ctx, rootID, saleID)
}
