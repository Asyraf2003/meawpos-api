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

package config

import (
	"os"
	"testing"
)

func TestLoad_DebugDisabledByDefault(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("AUTH_DEBUG_ENABLED", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Auth.Debug.Enabled {
		t.Fatal("Auth.Debug.Enabled = true, want false")
	}
}

func TestLoad_DebugEnabledFromEnv(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("AUTH_DEBUG_ENABLED", "true")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if !cfg.Auth.Debug.Enabled {
		t.Fatal("Auth.Debug.Enabled = false, want true")
	}
}

func TestLoad_InvalidDebugBool(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("AUTH_DEBUG_ENABLED", "not-bool")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want error")
	}
}

func TestMain(m *testing.M) {
	_ = os.Unsetenv("AUTH_DEBUG_ENABLED")
	_ = os.Unsetenv("DATABASE_URL")
	os.Exit(m.Run())
}
