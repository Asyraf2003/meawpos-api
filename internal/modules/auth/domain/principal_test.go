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

package domain

import "testing"

func TestPrincipalHasPermission(t *testing.T) {
	principal := Principal{
		AccountID: "acc-123",
		SessionID: "sess-123",
		Roles: []string{
			"base",
			"cashier",
		},
		Permissions: []string{
			"profile.self.read",
			"sale.order.create",
			"sale.order.read",
		},
		TrustLevel: "aal1",
	}

	if !principal.HasPermission("sale.order.create") {
		t.Fatal("expected permission sale.order.create to be present")
	}

	if principal.HasPermission("inventory.manage") {
		t.Fatal("did not expect permission inventory.manage to be present")
	}
}
