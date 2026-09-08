// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package bootstrap

import (
	"context"
	"testing"

	"pos-go/internal/config"
)

func TestNew_DoesNotRegisterInactiveBusinessRoutes(t *testing.T) {
	cfg := config.Config{
		DatabaseURL: testDatabaseURL(t),
		Auth: config.AuthConfig{Debug: config.DebugConfig{Enabled: true}, JWT: config.JWTConfig{
			Issuer: "pos-go", Aud: "client", Kid: "kid", Secret: "test-secret-123", TTL: 15,
		}, SessionTTL: 720},
		Components: config.ComponentConfig{Business: []string{}},
	}
	app, err := New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer app.DB.Close()
	if !hasRoute(app, "POST", "/api/roots") {
		t.Fatal("trusted ROOT route should be registered")
	}
	if hasRoute(app, "POST", "/api/roots/:root_id/catalog/items") {
		t.Fatal("inactive catalog route registered")
	}
	if hasRoute(app, "POST", "/api/roots/:root_id/sales") {
		t.Fatal("inactive sales route registered")
	}
}
