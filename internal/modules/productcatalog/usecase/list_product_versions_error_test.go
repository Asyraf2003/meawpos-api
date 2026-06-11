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

	"pos-go/internal/modules/productcatalog/ports"
)

func TestListProductVersionsPropagatesRepositoryError(t *testing.T) {
	listErr := errors.New("list product versions failed")
	usecase := NewListProductVersions(&listProductVersionsRepositoryDouble{
		listErr: listErr,
	})

	_, err := usecase.Execute(context.Background(), ListProductVersionsQuery{
		ProductID: "prod_001",
	})

	if !errors.Is(err, listErr) {
		t.Fatalf("Execute() error = %v, want %v", err, listErr)
	}
}

type listProductVersionsRepositoryDouble struct {
	capturedProductID string
	records           []ports.ProductVersionRecord
	listErr           error
	appended          []ports.ProductVersionRecord
}

func (d *listProductVersionsRepositoryDouble) Append(
	_ context.Context,
	version ports.ProductVersionRecord,
) error {
	d.appended = append(d.appended, version)

	return nil
}

func (d *listProductVersionsRepositoryDouble) ListByProductID(
	_ context.Context,
	productID string,
) ([]ports.ProductVersionRecord, error) {
	d.capturedProductID = productID
	if d.listErr != nil {
		return nil, d.listErr
	}

	return d.records, nil
}
