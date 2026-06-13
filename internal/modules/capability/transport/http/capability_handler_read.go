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
	"strings"

	capabilitypresenter "pos-go/internal/presentation/http/id/capability"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/labstack/echo/v4"
)

func (h *CapabilityHandler) List(c echo.Context) error {
	capabilities, err := h.listUsecase.Execute(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, httpresponse.Success(capabilitypresenter.FromDomainList(capabilities)))
}

func (h *CapabilityHandler) Show(c echo.Context) error {
	key := strings.TrimSpace(c.Param("key"))
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "capability key is required")
	}

	capability, err := h.showUsecase.Execute(c.Request().Context(), key)
	if err != nil {
		return capabilityHTTPError(err)
	}

	return c.JSON(http.StatusOK, httpresponse.Success(capabilitypresenter.FromDomain(capability)))
}
