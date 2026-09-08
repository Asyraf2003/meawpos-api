// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package authorization

import "testing"

func TestAuthorityCombinesMultipleRoleGrants(t *testing.T) {
	authority := Authority{
		Roles:       []string{"cashier", "reader"},
		Permissions: []string{PermissionSaleOrderCreate, PermissionPaymentCreate, PermissionSaleOrderRead},
	}
	if !authority.HasPermissions(PermissionSaleOrderCreate, PermissionPaymentCreate) {
		t.Fatal("expected union of role grants to satisfy both permissions")
	}
	if authority.HasPermissions(PermissionPaymentRefund) {
		t.Fatal("unexpected permission")
	}
}

func TestAuthorityMembershipWithoutGrantsHasNoOperationAuthority(t *testing.T) {
	authority := Authority{AccountID: "member"}
	if authority.HasPermissions(PermissionCatalogItemRead) {
		t.Fatal("membership without a relevant grant must not authorize an operation")
	}
}

func TestAuthorityPrimaryOwnerHasFullRootAuthority(t *testing.T) {
	authority := Authority{AccountID: "owner", PrimaryOwner: true}
	if !authority.HasPermissions(
		PermissionCatalogItemRead,
		PermissionCatalogItemCreate,
		PermissionSaleOrderCreate,
		PermissionSaleOrderRead,
		PermissionSaleOrderReverse,
		PermissionPaymentCreate,
		PermissionPaymentRefund,
	) {
		t.Fatal("primary owner must resolve to full ROOT authority")
	}
}
