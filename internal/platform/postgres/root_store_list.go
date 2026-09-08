// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"

	rootdomain "pos-go/internal/core/root/domain"
)

func (s *RootStore) ListRootsForAccount(ctx context.Context, accountID string) ([]rootdomain.Root, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT r.id,r.name,r.primary_owner_account_id,r.created_at,r.updated_at
		FROM roots r JOIN root_memberships m ON m.root_id=r.id
		WHERE m.account_id=$1 ORDER BY r.created_at,r.id`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]rootdomain.Root, 0)
	for rows.Next() {
		var root rootdomain.Root
		if err := rows.Scan(&root.ID, &root.Name, &root.PrimaryOwnerAccountID, &root.CreatedAt, &root.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, root)
	}
	return result, rows.Err()
}
