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

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type RefreshTokenUsecase interface {
	Execute(ctx context.Context, in authusecase.RefreshTokenInput) (authusecase.RefreshTokenOutput, error)
}

type RefreshHandler struct {
	usecase RefreshTokenUsecase
}

func NewRefreshHandler(usecase RefreshTokenUsecase) *RefreshHandler {
	return &RefreshHandler{usecase: usecase}
}

func (h *RefreshHandler) Register(group *echo.Group) {
	group.POST("/refresh", h.Refresh)
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *RefreshHandler) Refresh(c echo.Context) error {
	var req refreshRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	out, err := h.usecase.Execute(c.Request().Context(), authusecase.RefreshTokenInput{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		if err == authusecase.ErrInvalidRefreshToken {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid refresh token")
		}
		return err
	}

	return c.JSON(http.StatusOK, out)
}
