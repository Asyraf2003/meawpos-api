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

func TestCapabilityHandler_EnableSuccess(t *testing.T) {
	e := echo.New()
	c, rec := newCapabilityContext(e, http.MethodPost, "/api/admin/capabilities/capability.manage/enable", "")
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")
	handler, fake := newCapabilityHandlerForTest(t)
	fake.capabilities["capability.manage"] = fake.capabilities["capability.manage"].Disable("maintenance")

	if err := handler.Enable(c); err != nil {
		t.Fatalf("Enable() error = %v", err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertEqual(t, fake.enableCalls, 1, "enable calls")
	data := decodeCapability(t, decodeEnvelope(t, rec).Data)
	if !data.Enabled {
		t.Fatal("enabled = false, want true")
	}
	assertEqual(t, data.DisabledReason, "", "disabled_reason")
}

func TestCapabilityHandler_DisableSuccessWithReason(t *testing.T) {
	e := echo.New()
	c, rec := newCapabilityContext(e, http.MethodPost, "/api/admin/capabilities/capability.manage/disable", `{"reason":"maintenance"}`)
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")
	handler, fake := newCapabilityHandlerForTest(t)

	if err := handler.Disable(c); err != nil {
		t.Fatalf("Disable() error = %v", err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertEqual(t, fake.disableCalls, 1, "disable calls")
	assertEqual(t, fake.lastDisableReason, "maintenance", "reason")
	data := decodeCapability(t, decodeEnvelope(t, rec).Data)
	if data.Enabled {
		t.Fatal("enabled = true, want false")
	}
	assertEqual(t, data.DisabledReason, "maintenance", "disabled_reason")
}
