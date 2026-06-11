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

	"pos-go/internal/modules/productcatalog/ports"
)

func TestLookupProductsForwardsQuery(t *testing.T) {
	reader := &lookupProductsReaderDouble{}
	usecase := NewLookupProducts(reader)

	_, err := usecase.Execute(context.Background(), LookupProductsQuery{
		Query:          "filter",
		Limit:          15,
		IncludeDeleted: true,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if reader.capturedQuery.Query != "filter" {
		t.Fatalf("Query = %q, want filter", reader.capturedQuery.Query)
	}
	if reader.capturedQuery.Limit != 15 {
		t.Fatalf("Limit = %d, want 15", reader.capturedQuery.Limit)
	}
	if !reader.capturedQuery.IncludeDeleted {
		t.Fatalf("IncludeDeleted = false, want true")
	}
}

var _ ports.ProductReader = (*lookupProductsReaderDouble)(nil)
