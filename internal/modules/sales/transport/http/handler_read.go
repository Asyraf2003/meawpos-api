// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	stdhttp "net/http"

	salesusecase "pos-go/internal/modules/sales/usecase"
	httpmw "pos-go/internal/transport/http/middleware"
	httprequest "pos-go/internal/transport/http/request"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Show(c echo.Context) error {
	if uuid.Validate(c.Param("sale_id")) != nil {
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_sale_id", "invalid sale id")
	}
	sale, err := h.get.Execute(c.Request().Context(), c.Param("root_id"), c.Param("sale_id"))
	if err != nil {
		return mapError(err)
	}
	return c.JSON(stdhttp.StatusOK, httpresponse.Success(present(sale)))
}

func (h *Handler) Reverse(c echo.Context) error {
	if uuid.Validate(c.Param("sale_id")) != nil {
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_sale_id", "invalid sale id")
	}
	var req reversalRequest
	if err := httprequest.DecodeJSON(c, &req); err != nil {
		return err
	}
	actor, session, authority := requestAuthority(c)
	_, err := h.reverse.Execute(c.Request().Context(), salesusecase.ReverseSaleCommand{RootID: c.Param("root_id"), SaleID: c.Param("sale_id"), ActorAccountID: actor, SessionID: session, RequestID: httpmw.RequestIDFromContext(c), AuthorityUsed: authority, Reason: req.Reason})
	if err != nil {
		return mapError(err)
	}
	sale, err := h.get.Execute(c.Request().Context(), c.Param("root_id"), c.Param("sale_id"))
	if err != nil {
		return mapError(err)
	}
	return c.JSON(stdhttp.StatusCreated, httpresponse.Success(present(sale)))
}
