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

func TestListProductsReturnsEmptyItemsForEmptyReaderList(t *testing.T) {
	usecase := NewListProducts(&listProductsReaderDouble{})

	result, err := usecase.Execute(context.Background(), ListProductsQuery{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 0 {
		t.Fatalf("result Items length = %d, want %d", len(result.Items), 0)
	}
}
