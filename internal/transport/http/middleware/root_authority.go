// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package middleware

import (
	"context"
	"errors"
	"net/http"

	"pos-go/internal/core/authorization"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type rootAuthorityContextKey struct{}

func RootAuthorityFromContext(ctx context.Context) (authorization.Authority, bool) {
	authority, ok := ctx.Value(rootAuthorityContextKey{}).(authorization.Authority)
	return authority, ok
}

func RequireRootAuthority(resolver authorization.Resolver, required ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			principal, ok := PrincipalFromContext(c.Request().Context())
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "authentication required")
			}
			rootID := c.Param("root_id")
			if uuid.Validate(rootID) != nil {
				return httpresponse.NewHTTPError(http.StatusBadRequest, "invalid_root_id", "invalid root id")
			}
			authority, err := resolver.ResolveRootAuthority(c.Request().Context(), rootID, principal.AccountID)
			if errors.Is(err, authorization.ErrRootAccessDenied) {
				return httpresponse.NewHTTPError(http.StatusForbidden, "root_access_denied", "root access denied")
			}
			if err != nil {
				return err
			}
			if !authority.HasPermissions(required...) {
				return httpresponse.NewHTTPError(http.StatusForbidden, "root_access_denied", "root access denied")
			}
			ctx := context.WithValue(c.Request().Context(), rootAuthorityContextKey{}, authority)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
