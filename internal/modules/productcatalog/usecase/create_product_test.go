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

	"pos-go/internal/modules/productcatalog/domain"
)

func TestCreateProductSuccessPersistsProductAndChecksDuplicateCandidate(t *testing.T) {
	fixedNow := time.Date(2026, 6, 9, 12, 0, 0, 0, time.UTC)
	repository := &fakeProductRepository{}
	duplicateChecker := &fakeProductDuplicateChecker{}
	versionRepository := &fakeProductVersionRepository{}
	auditRecorder := &fakeProductAuditRecorder{}

	uc := NewCreateProduct(
		repository,
		duplicateChecker,
		versionRepository,
		auditRecorder,
		fakeProductIDGenerator{id: "prod_001"},
		func() time.Time { return fixedNow },
	)

	result, err := uc.Execute(context.Background(), CreateProductCommand{
		Code:            "  PRD-001  ",
		Name:            "  Oli   Mesin ",
		Brand:           " Yamaha   Genuine ",
		Size:            domain.IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.created == nil {
		t.Fatalf("repository.Create was not called")
	}

	assertCreateProductDuplicateCandidate(t, duplicateChecker)
	assertCreateProductResult(t, result, fixedNow)
	assertCreateProductSideEffects(t, versionRepository, auditRecorder, fixedNow)

}

func TestCreateProductReturnsDuplicateCheckerError(t *testing.T) {
	duplicateErr := errors.New("duplicate failure")

	uc := NewCreateProduct(
		&fakeProductRepository{},
		&fakeProductDuplicateChecker{err: duplicateErr},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		fakeProductIDGenerator{id: "prod_002"},
		time.Now,
	)

	_, err := uc.Execute(context.Background(), CreateProductCommand{
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})
	if !errors.Is(err, duplicateErr) {
		t.Fatalf("Execute() error = %v, want %v", err, duplicateErr)
	}
}
