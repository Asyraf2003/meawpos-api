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

package bootstrap

import (
	"os"
	"strings"
	"testing"
)

func TestNew_WiresCapabilityGuardsForProtectedRoutes(t *testing.T) {
	source, err := os.ReadFile("app.go")
	if err != nil {
		t.Fatalf("ReadFile(app.go) error = %v", err)
	}

	required := []string{
		`RequireCapability("profile.self.show"`,
		`RequireCapability("authz.profile.self.show"`,
		`RequireCapability("auth.session.logout"`,
		`RequireCapability("account.role.assign"`,
		`RequireCapability("account.role.remove"`,
		`RequireCapability("capability.manage"`,
	}

	for _, needle := range required {
		if !strings.Contains(string(source), needle) {
			t.Fatalf("missing bootstrap capability guard: %s", needle)
		}
	}
}
