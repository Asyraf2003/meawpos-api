// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"errors"
	stdhttp "net/http"

	"pos-go/internal/core/money"
	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
	"pos-go/internal/modules/sales/ports"
	httpresponse "pos-go/internal/transport/http/response"
)

func mapError(err error) error {
	switch {
	case errors.Is(err, ports.ErrCatalogItemNotFound):
		return httpresponse.NewHTTPError(stdhttp.StatusNotFound, "catalog_item_not_found", "catalog item not found")
	case errors.Is(err, ports.ErrCatalogItemNotSellable):
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "catalog_item_not_sellable", "catalog item is not sellable")
	case errors.Is(err, domain.ErrSaleNotFound):
		return httpresponse.NewHTTPError(stdhttp.StatusNotFound, "sale_not_found", "sale not found")
	case errors.Is(err, domain.ErrSaleAlreadyReversed):
		return httpresponse.NewHTTPError(stdhttp.StatusConflict, "sale_already_reversed", "sale already reversed")
	case errors.Is(err, domain.ErrIdempotencyConflict):
		return httpresponse.NewHTTPError(stdhttp.StatusConflict, "idempotency_conflict", "idempotency key conflicts with the original request")
	case errors.Is(err, money.ErrOverflow):
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "money_overflow", "money calculation overflow")
	case errors.Is(err, cash.ErrInsufficientTender):
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "insufficient_cash_tender", err.Error())
	case errors.Is(err, domain.ErrInvalidQuantity):
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_quantity", err.Error())
	case errors.Is(err, domain.ErrItemsRequired), errors.Is(err, domain.ErrUnsupportedPayment), errors.Is(err, domain.ErrIdempotencyKeyRequired), errors.Is(err, domain.ErrReversalReasonRequired):
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "sale_validation_failed", err.Error())
	default:
		return err
	}
}
