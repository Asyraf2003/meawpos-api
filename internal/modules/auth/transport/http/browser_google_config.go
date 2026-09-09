// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"errors"
	"net/url"
	"strings"
)

const browserGoogleCallbackPath = "/api/auth/browser/google/callback"

// The existing server-owned Google redirect configuration owns the browser origin.
// Request Host, forwarded headers and return URLs never select a redirect target.
func browserGoogleCallbackURL(configured string, secure bool) (string, error) {
	u, err := url.Parse(strings.TrimSpace(configured))
	if err != nil || u.Host == "" || u.Hostname() == "" || u.User != nil || u.Opaque != "" ||
		u.RawQuery != "" || u.ForceQuery || u.Fragment != "" ||
		(u.Scheme != "https" && (secure || u.Scheme != "http")) {
		return "", errors.New("invalid Google browser redirect configuration: HTTPS origin required outside local mode")
	}
	return (&url.URL{Scheme: u.Scheme, Host: u.Host, Path: browserGoogleCallbackPath}).String(), nil
}
