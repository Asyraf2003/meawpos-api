// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"

	rootdomain "pos-go/internal/core/root/domain"
	rootports "pos-go/internal/core/root/ports"
)

type ListRoots struct{ store rootports.Store }

func NewListRoots(store rootports.Store) *ListRoots { return &ListRoots{store: store} }

func (uc *ListRoots) Execute(ctx context.Context, accountID string) ([]rootdomain.Root, error) {
	return uc.store.ListRootsForAccount(ctx, accountID)
}
