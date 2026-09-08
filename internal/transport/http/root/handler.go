// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package root

import (
	"context"
	"net/http"

	rootdomain "pos-go/internal/core/root/domain"
	rootusecase "pos-go/internal/core/root/usecase"
	httpmw "pos-go/internal/transport/http/middleware"
	httprequest "pos-go/internal/transport/http/request"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/labstack/echo/v4"
)

type creator interface {
	Execute(context.Context, rootusecase.CreateRootCommand) (rootdomain.Root, error)
}
type lister interface {
	Execute(context.Context, string) ([]rootdomain.Root, error)
}

type Handler struct {
	create creator
	list   lister
}

func NewHandler(create creator, list lister) *Handler { return &Handler{create: create, list: list} }

func (h *Handler) Register(group *echo.Group) {
	group.POST("/roots", h.Create)
	group.GET("/roots", h.List)
}

type createRequest struct {
	Name string `json:"name"`
}
type rootResponse struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	PrimaryOwnerAccountID string `json:"primary_owner_account_id"`
	CreatedAt             string `json:"created_at"`
}

func (h *Handler) Create(c echo.Context) error {
	var req createRequest
	if err := httprequest.DecodeJSON(c, &req); err != nil {
		return err
	}
	principal, _ := httpmw.PrincipalFromContext(c.Request().Context())
	root, err := h.create.Execute(c.Request().Context(), rootusecase.CreateRootCommand{
		Name: req.Name, AccountID: principal.AccountID, SessionID: principal.SessionID,
		RequestID: httpmw.RequestIDFromContext(c),
	})
	if err != nil {
		return mapError(err)
	}
	return c.JSON(http.StatusCreated, httpresponse.Success(present(root)))
}
