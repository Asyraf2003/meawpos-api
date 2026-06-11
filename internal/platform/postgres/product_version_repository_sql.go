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

func productVersionInsertSQL() string {
	return `
		INSERT INTO product_versions (
			id,
			product_id,
			revision_no,
			event_name,
			changed_by_actor_id,
			change_reason,
			changed_at,
			snapshot_json
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, '{}'::jsonb)
	`
}

func productVersionListSQL() string {
	return `
		SELECT
			product_id,
			revision_no,
			event_name,
			COALESCE(changed_by_actor_id, ''),
			COALESCE(change_reason, ''),
			changed_at
		FROM product_versions
		WHERE product_id = $1
		ORDER BY
			revision_no ASC,
			changed_at ASC,
			id ASC
	`
}
