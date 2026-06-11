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

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestLookupProductsPropagatesReaderLookupError(t *testing.T) {
	lookupErr := errors.New("lookup products failed")
	usecase := NewLookupProducts(&lookupProductsReaderDouble{
		lookupErr: lookupErr,
	})

	_, err := usecase.Execute(context.Background(), LookupProductsQuery{})

	if !errors.Is(err, lookupErr) {
		t.Fatalf("Execute() error = %v, want %v", err, lookupErr)
	}
}

type lookupProductsReaderDouble struct {
	capturedQuery ports.ProductLookupQuery
	lookupItems   []ports.ProductLookupItem
	lookupErr     error
}

func (d *lookupProductsReaderDouble) GetByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	return nil, nil
}

func (d *lookupProductsReaderDouble) List(
	_ context.Context,
	_ ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	return nil, nil
}

func (d *lookupProductsReaderDouble) Lookup(
	_ context.Context,
	query ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	d.capturedQuery = query
	if d.lookupErr != nil {
		return nil, d.lookupErr
	}

	return d.lookupItems, nil
}
