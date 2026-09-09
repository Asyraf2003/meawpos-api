// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"net/http"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type BrowserManualLoginHandler struct {
	usecase ManualLoginUsecase
	cookie  browserSessionCookie
}

func NewBrowserManualLoginHandler(usecase ManualLoginUsecase, secureCookie bool) *BrowserManualLoginHandler {
	return &BrowserManualLoginHandler{
		usecase: usecase,
		cookie:  newBrowserSessionCookie(secureCookie),
	}
}

func (h *BrowserManualLoginHandler) Register(group *echo.Group) {
	group.POST("/browser/manual/login", h.Login)
}

func (h *BrowserManualLoginHandler) Login(c echo.Context) error {
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

	h.cookie.Set(c, out.RefreshToken)
	return c.JSON(http.StatusOK, browserResponseFromManual(out))
}
