// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package bootstrap

import (
	"context"
	"pos-go/internal/config"
	"testing"
)

func TestNew_RegistersGoogleAuthRoutesWhenConfigured(t *testing.T) {
	cfg := config.Config{
		AppEnv:      "test",
		HTTPPort:    "8080",
		DatabaseURL: testDatabaseURL(t),
		Auth: config.AuthConfig{
			Google: config.GoogleConfig{
				Issuer:       "https://accounts.google.com",
				ClientID:     "client-id",
				ClientSecret: "client-secret",
				RedirectURL:  "https://localhost:4173/api/auth/google/callback",
			},
			JWT: config.JWTConfig{
				Issuer: "pos-go",
				Aud:    "pos-go-client",
				Kid:    "local-dev-key",
				Secret: "test-secret-123",
				TTL:    15,
			},
			StateTTL:   10,
			SessionTTL: 720,
		},
	}

	app, err := New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer app.DB.Close()

	if !hasRoute(app, "GET", "/api/auth/google/start") {
		t.Fatal("expected GET /api/auth/google/start to be registered")
	}

	if !hasRoute(app, "GET", "/api/auth/google/callback") {
		t.Fatal("expected GET /api/auth/google/callback to be registered")
	}
	for _, path := range []string{"/api/auth/browser/google/start", "/api/auth/browser/google/callback"} {
		if !hasRoute(app, "GET", path) {
			t.Fatalf("Google browser route missing: %s", path)
		}
	}
	for _, path := range []string{"/api/auth/manual/login", "/api/auth/browser/manual/login"} {
		if hasRoute(app, "POST", path) {
			t.Fatalf("debug login registered without debug configuration: %s", path)
		}
	}

}
