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

package http

import (
	"context"

	authusecase "pos-go/internal/modules/auth/usecase"
)

type fakeAssignAccountRoleUsecase struct {
	lastAccountID string
	lastRoleKey   string
	calls         int
	err           error
}

func (f *fakeAssignAccountRoleUsecase) Execute(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.calls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return f.err
}

type fakeRemoveAccountRoleUsecase struct {
	lastAccountID string
	lastRoleKey   string
	calls         int
	err           error
}

func (f *fakeRemoveAccountRoleUsecase) Execute(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.calls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return f.err
}

func newAccountRoleHandlerForTest(
	assignErr error,
	removeErr error,
) (*AccountRoleHandler, *fakeAssignAccountRoleUsecase, *fakeRemoveAccountRoleUsecase) {
	assignUsecase := &fakeAssignAccountRoleUsecase{err: assignErr}
	removeUsecase := &fakeRemoveAccountRoleUsecase{err: removeErr}

	handler := NewAccountRoleHandler(assignUsecase, removeUsecase)

	if removeErr == authusecase.ErrBaseRoleRemovalNotAllowed {
		return handler, assignUsecase, removeUsecase
	}

	return handler, assignUsecase, removeUsecase
}
