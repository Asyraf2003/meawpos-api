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
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestRestoreProductReturnsNotFound(t *testing.T) {
	usecase := NewRestoreProduct(
		&fakeProductRepository{},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		func() time.Time { return time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC) },
	)

	_, err := usecase.Execute(context.Background(), RestoreProductCommand{
		ID:      "product-404",
		ActorID: "actor-1",
		Reason:  "restore missing product",
	})

	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("expected product not found, got %v", err)
	}
}
