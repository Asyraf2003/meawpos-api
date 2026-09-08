// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import "context"

func (r *RootAuthorityResolver) loadRootGrants(ctx context.Context, rootID, accountID string) ([]string, []string, bool, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT rr.key, rp.permission_key
		FROM root_memberships m
		LEFT JOIN root_membership_roles mr ON mr.root_id=m.root_id AND mr.membership_id=m.id
		LEFT JOIN root_roles rr ON rr.root_id=mr.root_id AND rr.id=mr.role_id
		LEFT JOIN root_role_permissions rp ON rp.root_id=rr.root_id AND rp.role_id=rr.id
		WHERE m.root_id=$1 AND m.account_id=$2
		ORDER BY rr.key,rp.permission_key`, rootID, accountID)
	if err != nil {
		return nil, nil, false, err
	}
	defer rows.Close()
	roles, permissions := []string{}, []string{}
	roleSeen, permissionSeen, member := map[string]bool{}, map[string]bool{}, false
	for rows.Next() {
		member = true
		var role, permission *string
		if err := rows.Scan(&role, &permission); err != nil {
			return nil, nil, false, err
		}
		if role != nil && !roleSeen[*role] {
			roles = append(roles, *role)
			roleSeen[*role] = true
		}
		if permission != nil && !permissionSeen[*permission] {
			permissions = append(permissions, *permission)
			permissionSeen[*permission] = true
		}
	}
	return roles, permissions, member, rows.Err()
}
