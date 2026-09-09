// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"net/http"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type BrowserRefreshHandler struct {
	usecase RefreshTokenUsecase
	cookie  browserSessionCookie
}

func NewBrowserRefreshHandler(usecase RefreshTokenUsecase, secureCookie bool) *BrowserRefreshHandler {
	return &BrowserRefreshHandler{
		usecase: usecase,
		cookie:  newBrowserSessionCookie(secureCookie),
	}
}

func (h *BrowserRefreshHandler) Register(group *echo.Group) {
	group.POST("/browser/refresh", h.Refresh)
}

func (h *BrowserRefreshHandler) Refresh(c echo.Context) error {
	refreshToken, err := h.cookie.Read(c)
	if err != nil {
		h.cookie.Clear(c)
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid refresh token")
	}

	out, err := h.usecase.Execute(c.Request().Context(), authusecase.RefreshTokenInput{
		RefreshToken: refreshToken,
	})
	if err != nil {
		if err == authusecase.ErrInvalidRefreshToken {
			h.cookie.Clear(c)
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid refresh token")
		}
		return err
	}

	h.cookie.Set(c, out.RefreshToken)
	return c.JSON(http.StatusOK, browserResponseFromRefresh(out))
}
