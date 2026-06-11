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

//go:build integration

package postgres

import (
	"context"
	"testing"
)

func TestAccountRoleRemover_RemoveRole(t *testing.T) {
	ctx := context.Background()

	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	tx := mustBeginIntegrationTx(t, ctx, pool)
	defer tx.Rollback(ctx)

	txCtx := contextWithTx(ctx, tx)
	accountID, roleKey := mustInsertAccountRoleRemoverFixture(t, ctx, tx)

	countBefore := mustCountAccountRoleByKey(t, ctx, tx, accountID, roleKey)
	if countBefore != 1 {
		t.Fatalf("count before = %d, want 1", countBefore)
	}

	remover := NewAccountRoleRemover(pool)

	if err := remover.RemoveRole(txCtx, accountID, roleKey); err != nil {
		t.Fatalf("RemoveRole() first call error = %v", err)
	}
	if err := remover.RemoveRole(txCtx, accountID, roleKey); err != nil {
		t.Fatalf("RemoveRole() second call error = %v", err)
	}

	countAfter := mustCountAccountRoleByKey(t, ctx, tx, accountID, roleKey)
	if countAfter != 0 {
		t.Fatalf("count after = %d, want 0", countAfter)
	}
}
