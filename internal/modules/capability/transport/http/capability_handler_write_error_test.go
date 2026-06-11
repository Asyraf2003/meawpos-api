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

func TestCapabilityHandler_DisableRejectsInvalidBody(t *testing.T) {
	e := echo.New()
	c, _ := newCapabilityContext(e, http.MethodPost, "/api/admin/capabilities/capability.manage/disable", `{"reason":`)
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")
	handler, fake := newCapabilityHandlerForTest(t)

	err := handler.Disable(c)
	if err == nil {
		t.Fatal("Disable() error = nil, want error")
	}
	assertHTTPError(t, err, http.StatusBadRequest)
	assertEqual(t, fake.disableCalls, 0, "disable calls")
}
