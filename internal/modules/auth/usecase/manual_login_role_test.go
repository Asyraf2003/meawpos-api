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

package usecase

import (
	"context"
	"testing"
	"time"
)

func TestManualLogin_KasirAssignsCashierRole(t *testing.T) {
	roles := &fakeManualRoleAssigner{}
	usecase := NewManualLogin(
		&fakeManualAccountRepository{accountID: "acc-kasir"},
		roles,
		&fakeSessionStore{},
		&fakeTokenIssuer{token: "access-token", exp: time.Now().Add(15 * time.Minute)},
		&fakeTransactor{},
		30*24*time.Hour,
	)

	_, err := usecase.Execute(context.Background(), ManualLoginInput{
		Email:    "kasir@example.com",
		Password: "12345678",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if roles.lastRoleKey != "cashier" {
		t.Fatalf("role key = %q, want cashier", roles.lastRoleKey)
	}
}

func TestManualLogin_RejectsUnsupportedEmail(t *testing.T) {
	usecase := NewManualLogin(
		&fakeManualAccountRepository{accountID: "acc"},
		&fakeManualRoleAssigner{},
		&fakeSessionStore{},
		&fakeTokenIssuer{},
		&fakeTransactor{},
		30*24*time.Hour,
	)

	_, err := usecase.Execute(context.Background(), ManualLoginInput{
		Email:    "owner@example.com",
		Password: "12345678",
	})
	if err != ErrManualLoginInvalidCredentials {
		t.Fatalf("error = %v, want ErrManualLoginInvalidCredentials", err)
	}
}

func TestManualLogin_RejectsInvalidPassword(t *testing.T) {
	usecase := NewManualLogin(
		&fakeManualAccountRepository{accountID: "acc"},
		&fakeManualRoleAssigner{},
		&fakeSessionStore{},
		&fakeTokenIssuer{},
		&fakeTransactor{},
		30*24*time.Hour,
	)

	_, err := usecase.Execute(context.Background(), ManualLoginInput{
		Email:    "admin@example.com",
		Password: "wrong-password",
	})
	if err != ErrManualLoginInvalidCredentials {
		t.Fatalf("error = %v, want ErrManualLoginInvalidCredentials", err)
	}
}
