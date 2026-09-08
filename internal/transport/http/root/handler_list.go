// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package root

import (
	"net/http"

	httpmw "pos-go/internal/transport/http/middleware"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/labstack/echo/v4"
)

func (h *Handler) List(c echo.Context) error {
	principal, _ := httpmw.PrincipalFromContext(c.Request().Context())
	roots, err := h.list.Execute(c.Request().Context(), principal.AccountID)
	if err != nil {
		return err
	}
	result := make([]rootResponse, len(roots))
	for i := range roots {
		result[i] = present(roots[i])
	}
	return c.JSON(http.StatusOK, httpresponse.Success(result))
}
