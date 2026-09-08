// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"context"
	stdhttp "net/http"

	"pos-go/internal/modules/sales/domain"
	salesusecase "pos-go/internal/modules/sales/usecase"
	httpmw "pos-go/internal/transport/http/middleware"
	httprequest "pos-go/internal/transport/http/request"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type poster interface {
	Execute(context.Context, salesusecase.PostCashSaleCommand) (salesusecase.PostCashSaleResult, error)
}
type getter interface {
	Execute(context.Context, string, string) (domain.Sale, error)
}
type reverser interface {
	Execute(context.Context, salesusecase.ReverseSaleCommand) (domain.Reversal, error)
}
type Handler struct {
	post    poster
	get     getter
	reverse reverser
}

func NewHandler(post poster, get getter, reverse reverser) *Handler {
	return &Handler{post: post, get: get, reverse: reverse}
}

func (h *Handler) RegisterCreate(group *echo.Group) { group.POST("/sales", h.Create) }
func (h *Handler) RegisterRead(group *echo.Group)   { group.GET("/sales/:sale_id", h.Show) }
func (h *Handler) RegisterReverse(group *echo.Group) {
	group.POST("/sales/:sale_id/reversals", h.Reverse)
}

type createRequest struct {
	Items   []salesusecase.SaleItemInput `json:"items"`
	Payment paymentRequest               `json:"payment"`
}
type paymentRequest struct {
	Type           string `json:"type"`
	TenderedRupiah int64  `json:"tendered_rupiah"`
}
type reversalRequest struct {
	Reason string `json:"reason"`
}

func requestAuthority(c echo.Context) (string, string, string) {
	principal, _ := httpmw.PrincipalFromContext(c.Request().Context())
	authority, _ := httpmw.RootAuthorityFromContext(c.Request().Context())
	return principal.AccountID, principal.SessionID, authority.Snapshot()
}

func (h *Handler) Create(c echo.Context) error {
	var req createRequest
	if err := httprequest.DecodeJSON(c, &req); err != nil {
		return err
	}
	for _, item := range req.Items {
		if uuid.Validate(item.CatalogItemID) != nil {
			return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_catalog_item_id", "invalid catalog item id")
		}
	}
	actor, session, authority := requestAuthority(c)
	result, err := h.post.Execute(c.Request().Context(), salesusecase.PostCashSaleCommand{RootID: c.Param("root_id"), ActorAccountID: actor, SessionID: session, RequestID: httpmw.RequestIDFromContext(c), AuthorityUsed: authority, IdempotencyKey: c.Request().Header.Get("Idempotency-Key"), Items: req.Items, PaymentType: req.Payment.Type, TenderedRupiah: req.Payment.TenderedRupiah})
	if err != nil {
		return mapError(err)
	}
	sale, err := h.get.Execute(c.Request().Context(), c.Param("root_id"), result.SaleID)
	if err != nil {
		return mapError(err)
	}
	status := stdhttp.StatusCreated
	if result.Replayed {
		status = stdhttp.StatusOK
	}
	return c.JSON(status, httpresponse.Success(present(sale)))
}
