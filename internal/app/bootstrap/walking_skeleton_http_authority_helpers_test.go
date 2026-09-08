// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"testing"

	"github.com/google/uuid"
)

func grantRootRole(t *testing.T, app *App, rootID, email, roleKey string, permissions ...string) {
	t.Helper()
	ctx := t.Context()
	var accountID string
	if err := app.DB.QueryRow(ctx, `SELECT id FROM accounts WHERE email=$1`, email).Scan(&accountID); err != nil {
		t.Fatal(err)
	}
	membershipID := uuid.NewString()
	if err := app.DB.QueryRow(ctx, `
		INSERT INTO root_memberships (id,root_id,account_id,created_at)
		VALUES ($1,$2,$3,now())
		ON CONFLICT (root_id,account_id) DO UPDATE SET account_id=EXCLUDED.account_id
		RETURNING id`, membershipID, rootID, accountID).Scan(&membershipID); err != nil {
		t.Fatal(err)
	}
	roleID := uuid.NewString()
	if _, err := app.DB.Exec(ctx, `INSERT INTO root_roles (id,root_id,key,name,created_at) VALUES ($1,$2,$3,$3,now())`, roleID, rootID, roleKey); err != nil {
		t.Fatal(err)
	}
	for _, permission := range permissions {
		if _, err := app.DB.Exec(ctx, `INSERT INTO root_role_permissions (root_id,role_id,permission_key,created_at) VALUES ($1,$2,$3,now())`, rootID, roleID, permission); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := app.DB.Exec(ctx, `INSERT INTO root_membership_roles (root_id,membership_id,role_id,created_at) VALUES ($1,$2,$3,now())`, rootID, membershipID, roleID); err != nil {
		t.Fatal(err)
	}
}

func addRootMembership(t *testing.T, app *App, rootID, email string) {
	t.Helper()
	grantRootRole(t, app, rootID, email, "no-authority-"+uuid.NewString())
}
