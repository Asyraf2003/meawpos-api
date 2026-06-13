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
	"errors"
	"net/http"

	"pos-go/internal/modules/capability/ports"

	"github.com/labstack/echo/v4"
)

type disableCapabilityRequest struct {
	Reason string `json:"reason"`
}

func capabilityHTTPError(err error) error {
	if errors.Is(err, ports.ErrCapabilityNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "capability not found")
	}

	return err
}
