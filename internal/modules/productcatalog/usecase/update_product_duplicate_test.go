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
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestUpdateProductChecksDuplicateCandidate(t *testing.T) {
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_001",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	checker := &fakeUpdateProductDuplicateChecker{}
	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		checker,
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_001",
		Code:            "  PRD-002  ",
		Name:            "  Oli   Mesin ",
		Brand:           " Yamaha   Genuine ",
		Size:            domain.IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if !checker.updateCalled {
		t.Fatalf("update duplicate checker was not called")
	}
	if checker.productID != "prod_001" {
		t.Fatalf("productID = %q, want prod_001", checker.productID)
	}
	if checker.candidate.Code == nil || *checker.candidate.Code != "PRD-002" {
		t.Fatalf("candidate.Code = %v, want PRD-002", checker.candidate.Code)
	}
	if checker.candidate.NormalizedName != "oli mesin" {
		t.Fatalf("candidate.NormalizedName = %q, want oli mesin", checker.candidate.NormalizedName)
	}
	if checker.candidate.NormalizedBrand != "yamaha genuine" {
		t.Fatalf("candidate.NormalizedBrand = %q, want yamaha genuine", checker.candidate.NormalizedBrand)
	}
}
