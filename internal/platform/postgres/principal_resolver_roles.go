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

package postgres

import "context"

func (r *PrincipalResolver) loadRoles(ctx context.Context, accountID string) ([]string, error) {
	rows, err := r.query(ctx, `
		SELECT r.key
		FROM account_roles ar
		JOIN roles r ON r.id = ar.role_id
		WHERE ar.account_id = $1
		ORDER BY r.key
	`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var roleKey string
		if err := rows.Scan(&roleKey); err != nil {
			return nil, err
		}
		roles = append(roles, roleKey)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
