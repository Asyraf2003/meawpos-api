// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"testing"

	"pos-go/internal/core/authorization"

	"github.com/google/uuid"
)

func TestRootAuthority_MultiRootMultiRoleUnionAndIsolation(t *testing.T) {
	f := newR4Fixture(t)
	ctx := context.Background()
	membershipID := uuid.NewString()
	_, err := f.pool.Exec(ctx, `INSERT INTO root_memberships (id,root_id,account_id,created_at) VALUES ($1,$2,$3,now())`, membershipID, f.rootID, f.otherAccountID)
	if err != nil {
		t.Fatal(err)
	}
	roleA, roleB := uuid.NewString(), uuid.NewString()
	_, err = f.pool.Exec(ctx, `INSERT INTO root_roles (id,root_id,key,name,created_at) VALUES ($1,$3,'cashier','Cashier',now()),($2,$3,'reader','Reader',now())`, roleA, roleB, f.rootID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.pool.Exec(ctx, `INSERT INTO root_role_permissions (root_id,role_id,permission_key,created_at) VALUES ($1,$2,$4,now()),($1,$2,$5,now()),($1,$3,$6,now())`, f.rootID, roleA, roleB, authorization.PermissionSaleOrderCreate, authorization.PermissionPaymentCreate, authorization.PermissionSaleOrderRead)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.pool.Exec(ctx, `INSERT INTO root_membership_roles (root_id,membership_id,role_id,created_at) VALUES ($1,$2,$3,now()),($1,$2,$4,now())`, f.rootID, membershipID, roleA, roleB)
	if err != nil {
		t.Fatal(err)
	}
	resolver := NewRootAuthorityResolver(f.pool)
	authority, err := resolver.ResolveRootAuthority(ctx, f.rootID, f.otherAccountID)
	if err != nil || len(authority.Roles) != 2 || !authority.HasPermissions(authorization.PermissionSaleOrderCreate, authorization.PermissionPaymentCreate, authorization.PermissionSaleOrderRead) {
		t.Fatalf("authority=%#v err=%v", authority, err)
	}
	owner, err := resolver.ResolveRootAuthority(ctx, f.otherRootID, f.otherAccountID)
	if err != nil || !owner.PrimaryOwner {
		t.Fatalf("owner=%#v err=%v", owner, err)
	}
	if _, err := resolver.ResolveRootAuthority(ctx, f.otherRootID, f.accountID); err == nil {
		t.Fatal("expected cross-root access denial")
	}
}

func TestRootAuthority_MembershipAloneDoesNotGrantOperationAuthority(t *testing.T) {
	f := newR4Fixture(t)
	ctx := context.Background()
	_, err := f.pool.Exec(ctx, `INSERT INTO root_memberships (id,root_id,account_id,created_at) VALUES ($1,$2,$3,now())`, uuid.NewString(), f.rootID, f.otherAccountID)
	if err != nil {
		t.Fatal(err)
	}
	authority, err := NewRootAuthorityResolver(f.pool).ResolveRootAuthority(ctx, f.rootID, f.otherAccountID)
	if err != nil {
		t.Fatal(err)
	}
	for _, permission := range []string{
		authorization.PermissionCatalogItemRead,
		authorization.PermissionCatalogItemCreate,
		authorization.PermissionSaleOrderCreate,
		authorization.PermissionSaleOrderRead,
		authorization.PermissionSaleOrderReverse,
		authorization.PermissionPaymentCreate,
		authorization.PermissionPaymentRefund,
	} {
		if authority.HasPermissions(permission) {
			t.Fatalf("membership unexpectedly has %s authority: %#v", permission, authority)
		}
	}
	if authority.PrimaryOwner || len(authority.Permissions) != 0 {
		t.Fatalf("membership authority=%#v", authority)
	}
}

func TestRootAuthority_RejectsCrossRootRoleAssignment(t *testing.T) {
	f := newR4Fixture(t)
	ctx := context.Background()
	roleID := uuid.NewString()
	_, err := f.pool.Exec(ctx, `INSERT INTO root_roles (id,root_id,key,name,created_at) VALUES ($1,$2,'role','Role',now())`, roleID, f.rootID)
	if err != nil {
		t.Fatal(err)
	}
	var membershipID string
	if err := f.pool.QueryRow(ctx, `SELECT id FROM root_memberships WHERE root_id=$1`, f.otherRootID).Scan(&membershipID); err != nil {
		t.Fatal(err)
	}
	if _, err := f.pool.Exec(ctx, `INSERT INTO root_membership_roles (root_id,membership_id,role_id,created_at) VALUES ($1,$2,$3,now())`, f.rootID, membershipID, roleID); err == nil {
		t.Fatal("expected cross-root constraint failure")
	}
}
