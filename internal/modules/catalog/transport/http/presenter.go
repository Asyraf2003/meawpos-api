// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"errors"
	stdhttp "net/http"

	catalogcore "pos-go/internal/modules/catalog/core"
	catalogports "pos-go/internal/modules/catalog/ports"
	"pos-go/internal/modules/catalog/pricing"
	httpresponse "pos-go/internal/transport/http/response"
)

type itemResponse struct {
	ID     string         `json:"id"`
	RootID string         `json:"root_id"`
	Name   string         `json:"name"`
	Price  *priceResponse `json:"price,omitempty"`
}
type priceResponse struct {
	Currency     string `json:"currency"`
	AmountRupiah int64  `json:"amount_rupiah"`
}

func present(result catalogports.ItemWithPrice) itemResponse {
	response := itemResponse{ID: result.Item.ID, RootID: result.Item.RootID, Name: result.Item.Name}
	if result.Price != nil {
		response.Price = &priceResponse{Currency: result.Price.Amount.Currency(), AmountRupiah: result.Price.Amount.AmountRupiah()}
	}
	return response
}

func mapError(err error) error {
	switch {
	case errors.Is(err, catalogports.ErrItemNotFound):
		return httpresponse.NewHTTPError(stdhttp.StatusNotFound, "catalog_item_not_found", "catalog item not found")
	case errors.Is(err, catalogcore.ErrNameRequired), errors.Is(err, catalogcore.ErrNameTooLong), errors.Is(err, pricing.ErrPriceMustBePositive):
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "catalog_item_validation_failed", err.Error())
	default:
		return err
	}
}
