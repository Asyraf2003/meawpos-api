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

package middleware

import (
	"errors"
	"net/http"
	"testing"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/labstack/echo/v4"
)

func TestRequireAuth_SetsPrincipalOnSuccess(t *testing.T) {
	c, rec := newAuthnTestContext("Bearer token-123")

	verifier := &fakeAccessTokenVerifier{
		claims: ports.AccessTokenClaims{
			AccountID:  "acc-123",
			SessionID:  "sess-123",
			TrustLevel: "aal1",
		},
	}
	resolver := &fakePrincipalResolver{
		principal: domain.Principal{
			AccountID:   "acc-123",
			SessionID:   "sess-123",
			Roles:       []string{"base"},
			Permissions: []string{"profile.self.read"},
			TrustLevel:  "aal1",
		},
	}
	sessionChecker := &fakeSessionStatusChecker{active: true}

	called := false
	handler := RequireAuth(verifier, resolver, sessionChecker)(func(c echo.Context) error {
		called = true
		assertPrincipalInContext(t, c.Request().Context(), "acc-123", "profile.self.read")
		return c.NoContent(http.StatusNoContent)
	})

	if err := handler(c); err != nil {
		t.Fatalf("handler() error = %v", err)
	}
	if !called {
		t.Fatal("next handler was not called")
	}
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", rec.Code)
	}

	assertResolverInput(t, resolver.lastInput, "acc-123", "sess-123", "aal1")
	assertSessionCheckerSessionID(t, sessionChecker.lastSessionID, "sess-123")
}

func TestRequireAuth_RejectsMissingBearerToken(t *testing.T) {
	c, _ := newAuthnTestContext("")

	handler := RequireAuth(
		&fakeAccessTokenVerifier{},
		&fakePrincipalResolver{},
		&fakeSessionStatusChecker{active: true},
	)(failIfCalledHandler(t))

	assertHTTPErrorCode(t, handler(c), http.StatusUnauthorized)
}

func TestRequireAuth_RejectsInvalidAccessToken(t *testing.T) {
	c, _ := newAuthnTestContext("Bearer token-123")

	handler := RequireAuth(
		&fakeAccessTokenVerifier{err: errors.New("bad token")},
		&fakePrincipalResolver{},
		&fakeSessionStatusChecker{active: true},
	)(failIfCalledHandler(t))

	assertHTTPErrorCode(t, handler(c), http.StatusUnauthorized)
}

func TestRequireAuth_RejectsInactiveSession(t *testing.T) {
	c, _ := newAuthnTestContext("Bearer token-123")

	handler := RequireAuth(
		&fakeAccessTokenVerifier{
			claims: ports.AccessTokenClaims{
				AccountID:  "acc-123",
				SessionID:  "sess-123",
				TrustLevel: "aal1",
			},
		},
		&fakePrincipalResolver{},
		&fakeSessionStatusChecker{active: false},
	)(failIfCalledHandler(t))

	assertHTTPErrorCode(t, handler(c), http.StatusUnauthorized)
}
