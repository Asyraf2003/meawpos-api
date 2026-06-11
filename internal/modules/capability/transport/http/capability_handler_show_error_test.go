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
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCapabilityHandler_ShowRejectsMissingKey(t *testing.T) {
	e := echo.New()
	c, _ := newCapabilityContext(e, http.MethodGet, "/api/admin/capabilities/", "")
	c.SetParamNames("key")
	c.SetParamValues(" ")
	handler, fake := newCapabilityHandlerForTest(t)

	err := handler.Show(c)
	if err == nil {
		t.Fatal("Show() error = nil, want error")
	}
	assertHTTPError(t, err, http.StatusBadRequest)
	assertEqual(t, fake.showCalls, 0, "show calls")
}

func TestCapabilityHandler_ShowMapsNotFound(t *testing.T) {
	e := echo.New()
	c, _ := newCapabilityContext(e, http.MethodGet, "/api/admin/capabilities/missing", "")
	c.SetParamNames("key")
	c.SetParamValues("missing")
	handler, _ := newCapabilityHandlerForTest(t)

	err := handler.Show(c)
	if err == nil {
		t.Fatal("Show() error = nil, want error")
	}
	assertHTTPError(t, err, http.StatusNotFound)
}
