// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

package http

import (
	"context"
	"net/http"
	"strings"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type AssignAccountRoleUsecase interface {
	Execute(ctx context.Context, accountID string, roleKey string) error
}

type RemoveAccountRoleUsecase interface {
	Execute(ctx context.Context, accountID string, roleKey string) error
}

type AccountRoleHandler struct {
	assignUsecase AssignAccountRoleUsecase
	removeUsecase RemoveAccountRoleUsecase
}

func NewAccountRoleHandler(
	assignUsecase AssignAccountRoleUsecase,
	removeUsecase RemoveAccountRoleUsecase,
) *AccountRoleHandler {
	return &AccountRoleHandler{
		assignUsecase: assignUsecase,
		removeUsecase: removeUsecase,
	}
}

func (h *AccountRoleHandler) Register(group *echo.Group) {
	h.RegisterAssign(group)
	h.RegisterRemove(group)
}

func (h *AccountRoleHandler) RegisterAssign(group *echo.Group) {
	group.POST("/accounts/:account_id/roles", h.Assign)
}

func (h *AccountRoleHandler) RegisterRemove(group *echo.Group) {
	group.DELETE("/accounts/:account_id/roles/:role_key", h.Remove)
}

type assignAccountRoleRequest struct {
	RoleKey string `json:"role_key"`
}

func (h *AccountRoleHandler) Assign(c echo.Context) error {
	var req assignAccountRoleRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	accountID := strings.TrimSpace(c.Param("account_id"))
	roleKey := strings.TrimSpace(req.RoleKey)

	if accountID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "account id is required")
	}
	if roleKey == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "role key is required")
	}

	if err := h.assignUsecase.Execute(c.Request().Context(), accountID, roleKey); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *AccountRoleHandler) Remove(c echo.Context) error {
	accountID := strings.TrimSpace(c.Param("account_id"))
	roleKey := strings.TrimSpace(c.Param("role_key"))

	if accountID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "account id is required")
	}
	if roleKey == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "role key is required")
	}

	if err := h.removeUsecase.Execute(c.Request().Context(), accountID, roleKey); err != nil {
		if err == authusecase.ErrBaseRoleRemovalNotAllowed {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
