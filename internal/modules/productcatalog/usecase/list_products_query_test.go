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

func TestListProductsForwardsQueryToReader(t *testing.T) {
	reader := &listProductsReaderDouble{}
	usecase := NewListProducts(reader)

	_, err := usecase.Execute(context.Background(), ListProductsQuery{
		Search:  "kopi",
		Status:  "active",
		Page:    2,
		PerPage: 25,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if reader.capturedQuery.Search != "kopi" {
		t.Fatalf("captured query Search = %q, want %q", reader.capturedQuery.Search, "kopi")
	}

	if reader.capturedQuery.Status != "active" {
		t.Fatalf("captured query Status = %q, want %q", reader.capturedQuery.Status, "active")
	}

	if reader.capturedQuery.Page != 2 {
		t.Fatalf("captured query Page = %d, want %d", reader.capturedQuery.Page, 2)
	}

	if reader.capturedQuery.PerPage != 25 {
		t.Fatalf("captured query PerPage = %d, want %d", reader.capturedQuery.PerPage, 25)
	}
}
