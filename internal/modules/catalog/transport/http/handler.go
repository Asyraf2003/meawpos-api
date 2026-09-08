// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"context"
	stdhttp "net/http"

	catalogports "pos-go/internal/modules/catalog/ports"
	catalogusecase "pos-go/internal/modules/catalog/usecase"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type creator interface {
	Execute(context.Context, catalogusecase.CreateItemCommand) (catalogports.ItemWithPrice, error)
}
type getter interface {
	Execute(context.Context, string, string) (catalogports.ItemWithPrice, error)
}
type Handler struct {
	create        creator
	get           getter
	pricingActive bool
}

func NewHandler(create creator, get getter, pricingActive bool) *Handler {
	return &Handler{create: create, get: get, pricingActive: pricingActive}
}

func (h *Handler) RegisterCreate(group *echo.Group) { group.POST("/catalog/items", h.Create) }
func (h *Handler) RegisterRead(group *echo.Group) {
	group.GET("/catalog/items/:item_id", h.Show)
}

type createRequest struct {
	Name        string `json:"name"`
	PriceRupiah *int64 `json:"price_rupiah"`
}

func (h *Handler) Create(c echo.Context) error {
	var req createRequest
	if err := c.Bind(&req); err != nil {
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_request_body", "invalid request body")
	}
	if req.PriceRupiah != nil && !h.pricingActive {
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "catalog_pricing_inactive", "catalog pricing component is inactive")
	}
	result, err := h.create.Execute(c.Request().Context(), catalogusecase.CreateItemCommand{RootID: c.Param("root_id"), Name: req.Name, PriceRupiah: req.PriceRupiah})
	if err != nil {
		return mapError(err)
	}
	return c.JSON(stdhttp.StatusCreated, httpresponse.Success(present(result)))
}

func (h *Handler) Show(c echo.Context) error {
	if uuid.Validate(c.Param("item_id")) != nil {
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_catalog_item_id", "invalid catalog item id")
	}
	result, err := h.get.Execute(c.Request().Context(), c.Param("root_id"), c.Param("item_id"))
	if err != nil {
		return mapError(err)
	}
	return c.JSON(stdhttp.StatusOK, httpresponse.Success(present(result)))
}
