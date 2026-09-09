// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func browserSessionRequest(
	t *testing.T,
	app *App,
	method, path, accessToken string,
	sessionCookie *http.Cookie,
	body any,
) (int, map[string]any, *http.Cookie) {
	t.Helper()

	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		reader = bytes.NewReader(encoded)
	}

	req := httptest.NewRequest(method, path, reader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
	if sessionCookie != nil {
		req.AddCookie(&http.Cookie{Name: sessionCookie.Name, Value: sessionCookie.Value})
	}

	rec := httptest.NewRecorder()
	app.Echo.ServeHTTP(rec, req)
	response := map[string]any{}
	if rec.Body.Len() > 0 {
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatalf("decode %s: %v body=%s", path, err, rec.Body.String())
		}
	}

	for _, cookie := range rec.Result().Cookies() {
		if cookie.Name == "miawpos_session" {
			return rec.Code, response, cookie
		}
	}
	return rec.Code, response, nil
}
