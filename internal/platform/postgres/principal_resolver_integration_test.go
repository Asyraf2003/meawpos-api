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

//go:build integration

package postgres

import (
	"context"
	"reflect"
	"testing"

	"pos-go/internal/modules/auth/ports"
)

func TestPrincipalResolver_Resolve(t *testing.T) {
	ctx := context.Background()

	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	tx := mustBeginIntegrationTx(t, ctx, pool)
	defer tx.Rollback(ctx)

	txCtx := contextWithTx(ctx, tx)
	accountID := mustInsertPrincipalResolverFixture(t, ctx, tx)

	resolver := NewPrincipalResolver(pool)

	principal, err := resolver.Resolve(txCtx, ports.ResolvePrincipalInput{
		AccountID:  accountID,
		SessionID:  "session-123",
		TrustLevel: "aal1",
	})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	if principal.AccountID != accountID {
		t.Fatalf("account_id = %q, want %q", principal.AccountID, accountID)
	}
	if principal.SessionID != "session-123" {
		t.Fatalf("session_id = %q, want session-123", principal.SessionID)
	}
	if principal.TrustLevel != "aal1" {
		t.Fatalf("trust_level = %q, want aal1", principal.TrustLevel)
	}

	wantRoles := []string{"base", "cashier"}
	if !reflect.DeepEqual(principal.Roles, wantRoles) {
		t.Fatalf("roles = %#v, want %#v", principal.Roles, wantRoles)
	}

	wantPermissions := []string{
		"auth.session.logout",
		"auth.session.refresh",
		"payment.create",
		"profile.self.read",
		"sale.order.create",
		"sale.order.read",
	}
	if !reflect.DeepEqual(principal.Permissions, wantPermissions) {
		t.Fatalf("permissions = %#v, want %#v", principal.Permissions, wantPermissions)
	}

	if !principal.HasPermission("sale.order.create") {
		t.Fatal("expected permission sale.order.create")
	}
	if principal.HasPermission("inventory.manage") {
		t.Fatal("did not expect permission inventory.manage")
	}
}
