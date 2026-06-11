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
	"net/http"

	"pos-go/internal/modules/system/ports"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	checker ports.HealthChecker
}

func NewHealthHandler(checker ports.HealthChecker) HealthHandler {
	return HealthHandler{checker: checker}
}

func (h HealthHandler) Register(group *echo.Group) {
	group.GET("/health", h.Get)
}

func (h HealthHandler) Get(c echo.Context) error {
	if err := h.checker.Check(c.Request().Context()); err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]any{
			"status":   "degraded",
			"database": "down",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":   "ok",
		"database": "up",
	})
}
