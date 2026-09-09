// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"net/http"

	authdomain "pos-go/internal/modules/auth/domain"
	authusecase "pos-go/internal/modules/auth/usecase"
	httpmw "pos-go/internal/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

type BrowserLogoutHandler struct {
	usecase *authusecase.LogoutCurrentSession
	cookie  browserSessionCookie
}

func NewBrowserLogoutHandler(usecase *authusecase.LogoutCurrentSession, secureCookie bool) *BrowserLogoutHandler {
	return &BrowserLogoutHandler{
		usecase: usecase,
		cookie:  newBrowserSessionCookie(secureCookie),
	}
}

func (h *BrowserLogoutHandler) Register(group *echo.Group) {
	group.POST("/browser/logout", h.Logout)
}

func (h *BrowserLogoutHandler) Logout(c echo.Context) error {
	principal, ok := httpmw.PrincipalFromContext(c.Request().Context())
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "authentication required")
	}

	if err := h.usecase.Execute(c.Request().Context(), authdomain.Principal{
		AccountID:   principal.AccountID,
		SessionID:   principal.SessionID,
		Roles:       principal.Roles,
		Permissions: principal.Permissions,
		TrustLevel:  principal.TrustLevel,
	}); err != nil {
		return err
	}

	h.cookie.Clear(c)
	return c.NoContent(http.StatusNoContent)
}
