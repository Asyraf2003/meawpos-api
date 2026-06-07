//go:build integration

package postgres

import (
	"context"
	"testing"
)

func TestSeededExistingProtectedCapabilities(t *testing.T) {
	ctx := context.Background()
	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	repository := NewCapabilityRepository(pool)

	expected := map[string]struct {
		method             string
		path               string
		requiredPermission string
	}{
		"profile.self.show": {
			method:             "GET",
			path:               "/api/me",
			requiredPermission: "profile.self.read",
		},
		"authz.profile.self.show": {
			method:             "GET",
			path:               "/api/authz/me",
			requiredPermission: "profile.self.read",
		},
		"auth.session.logout": {
			method:             "POST",
			path:               "/api/auth/logout",
			requiredPermission: "auth.session.logout",
		},
		"account.role.assign": {
			method:             "POST",
			path:               "/api/admin/accounts/:account_id/roles",
			requiredPermission: "account.role.assign",
		},
		"account.role.remove": {
			method:             "DELETE",
			path:               "/api/admin/accounts/:account_id/roles/:role_key",
			requiredPermission: "account.role.assign",
		},
	}

	for key, want := range expected {
		got, err := repository.Get(ctx, key)
		if err != nil {
			t.Fatalf("expected seeded capability %q: %v", key, err)
		}
		if got.Method != want.method {
			t.Fatalf("%s method = %q, want %q", key, got.Method, want.method)
		}
		if got.Path != want.path {
			t.Fatalf("%s path = %q, want %q", key, got.Path, want.path)
		}
		if got.RequiredPermission != want.requiredPermission {
			t.Fatalf("%s required permission = %q, want %q", key, got.RequiredPermission, want.requiredPermission)
		}
		if !got.DefaultEnabled {
			t.Fatalf("%s default enabled = false, want true", key)
		}
		if !got.Enabled {
			t.Fatalf("%s enabled = false, want true", key)
		}
	}
}
