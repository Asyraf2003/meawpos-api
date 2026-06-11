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

package domain

import "testing"

func TestNewProductTrimsAndNormalizesNameBrandAndCode(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_001",
		Code:            "  BRG-001  ",
		Name:            "  Oli   Mesin  ",
		Brand:           "  Yamaha   Genuine ",
		Size:            IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	if product.Code() == nil || *product.Code() != "BRG-001" {
		t.Fatalf("Code() = %v, want BRG-001", product.Code())
	}
	if product.Name() != "Oli Mesin" {
		t.Fatalf("Name() = %q, want Oli Mesin", product.Name())
	}
	if product.NormalizedName() != "oli mesin" {
		t.Fatalf("NormalizedName() = %q, want oli mesin", product.NormalizedName())
	}
	if product.Brand() != "Yamaha Genuine" {
		t.Fatalf("Brand() = %q, want Yamaha Genuine", product.Brand())
	}
	if product.NormalizedBrand() != "yamaha genuine" {
		t.Fatalf("NormalizedBrand() = %q, want yamaha genuine", product.NormalizedBrand())
	}
}

func TestNewProductConvertsBlankCodeToNil(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_002",
		Code:            "   ",
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	if product.Code() != nil {
		t.Fatalf("Code() = %v, want nil", product.Code())
	}
}
