// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"pos-go/internal/config"
)

func newR4HTTPApp(t *testing.T) *App {
	t.Helper()
	cfg := config.Config{DatabaseURL: testDatabaseURL(t), Auth: config.AuthConfig{Debug: config.DebugConfig{Enabled: true}, JWT: config.JWTConfig{Issuer: "pos-go", Aud: "pos-go-client", Kid: "test", Secret: "test-secret-123", TTL: 15 * time.Minute}, SessionTTL: 720 * time.Hour}, Components: config.ComponentConfig{Business: []string{"catalog.core", "catalog.pricing", "sales", "payment.cash"}}}
	app, err := New(t.Context(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(app.DB.Close)
	return app
}

func r4Request(t *testing.T, app *App, method, path, token, key string, body any) (int, map[string]any) {
	t.Helper()
	encoded, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(method, path, bytes.NewReader(encoded))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	if key != "" {
		req.Header.Set("Idempotency-Key", key)
	}
	rec := httptest.NewRecorder()
	app.Echo.ServeHTTP(rec, req)
	var response map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode %s: %v body=%s", path, err, rec.Body.String())
	}
	return rec.Code, response
}

func r4Login(t *testing.T, app *App, email string) string {
	t.Helper()
	status, body := r4Request(t, app, http.MethodPost, "/api/auth/manual/login", "", "", map[string]any{"email": email, "password": "12345678"})
	if status != http.StatusOK {
		t.Fatalf("login status=%d body=%v", status, body)
	}
	token, _ := body["access_token"].(string)
	if token == "" {
		t.Fatalf("login body=%v", body)
	}
	return token
}

func dataMap(t *testing.T, body map[string]any) map[string]any {
	t.Helper()
	data, ok := body["data"].(map[string]any)
	if !ok {
		t.Fatalf("body=%v", body)
	}
	return data
}
