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

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

type fakeProductIDGenerator struct {
	id  string
	err error
}

func (f fakeProductIDGenerator) NewProductID() (string, error) {
	if f.err != nil {
		return "", f.err
	}

	return f.id, nil
}

type fakeProductRepository struct {
	created *domain.Product
	err     error
}

func (f *fakeProductRepository) Create(_ context.Context, product *domain.Product) error {
	f.created = product

	return f.err
}

func (f *fakeProductRepository) Update(_ context.Context, _ *domain.Product) error {
	return nil
}

func (f *fakeProductRepository) FindByID(_ context.Context, _ string) (*domain.Product, error) {
	return nil, ports.ErrProductNotFound
}

type fakeProductDuplicateChecker struct {
	createCalled bool
	candidate    ports.ProductDuplicateCandidate
	err          error
}

func (f *fakeProductDuplicateChecker) CheckCreateDuplicate(
	_ context.Context,
	candidate ports.ProductDuplicateCandidate,
) error {
	f.createCalled = true
	f.candidate = candidate

	return f.err
}

func (f *fakeProductDuplicateChecker) CheckUpdateDuplicate(
	_ context.Context,
	_ string,
	_ ports.ProductDuplicateCandidate,
) error {
	return nil
}

type fakeProductVersionRepository struct {
	appended []ports.ProductVersionRecord
	err      error
}

func (f *fakeProductVersionRepository) Append(
	_ context.Context,
	version ports.ProductVersionRecord,
) error {
	f.appended = append(f.appended, version)

	return f.err
}

func (f *fakeProductVersionRepository) ListByProductID(
	_ context.Context,
	_ string,
) ([]ports.ProductVersionRecord, error) {
	return nil, nil
}

type fakeProductAuditRecorder struct {
	records []ports.ProductAuditRecord
	err     error
}

func (f *fakeProductAuditRecorder) RecordProductAudit(
	_ context.Context,
	record ports.ProductAuditRecord,
) error {
	f.records = append(f.records, record)

	return f.err
}
