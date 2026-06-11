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

type ManualLoginUsecase interface {
	Execute(ctx context.Context, in authusecase.ManualLoginInput) (authusecase.ManualLoginOutput, error)
}

type ManualLoginHandler struct {
	usecase ManualLoginUsecase
}

func NewManualLoginHandler(usecase ManualLoginUsecase) *ManualLoginHandler {
	return &ManualLoginHandler{usecase: usecase}
}

func (h *ManualLoginHandler) Register(group *echo.Group) {
	group.POST("/manual/login", h.Login)
}

type manualLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *ManualLoginHandler) Login(c echo.Context) error {
	var req manualLoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	out, err := h.usecase.Execute(c.Request().Context(), authusecase.ManualLoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if err == authusecase.ErrManualLoginInvalidCredentials {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid manual login credentials")
		}
		return err
	}

	return c.JSON(http.StatusOK, out)
}
