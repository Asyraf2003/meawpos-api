// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package ports

import (
	"context"

	rootdomain "pos-go/internal/core/root/domain"
)

type Store interface {
	CreateRoot(ctx context.Context, root rootdomain.Root, membershipID string) error
	ListRootsForAccount(ctx context.Context, accountID string) ([]rootdomain.Root, error)
}
