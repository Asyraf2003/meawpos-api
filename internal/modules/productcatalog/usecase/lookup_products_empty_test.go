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

import (
	"context"
	"testing"
)

func TestLookupProductsReturnsEmptyItems(t *testing.T) {
	usecase := NewLookupProducts(&lookupProductsReaderDouble{})

	result, err := usecase.Execute(context.Background(), LookupProductsQuery{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 0 {
		t.Fatalf("Items length = %d, want 0", len(result.Items))
	}
	if result.Items == nil {
		t.Fatalf("Items = nil, want empty slice")
	}
}
