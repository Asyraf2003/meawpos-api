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

package usecase

import "context"

type fakeManualAccountRepository struct {
	accountID string
	lastEmail string
}

func (f *fakeManualAccountRepository) ResolveOrCreateManualAccount(ctx context.Context, email string) (string, error) {
	_ = ctx
	f.lastEmail = email
	return f.accountID, nil
}

type fakeManualRoleAssigner struct {
	ensureCalls   int
	lastAccountID string
	lastRoleKey   string
}

func (f *fakeManualRoleAssigner) EnsureRole(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.ensureCalls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return nil
}
