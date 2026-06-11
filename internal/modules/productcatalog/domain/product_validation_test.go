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

func TestNewProductRejectsInvalidNameBrandAndPrice(t *testing.T) {
	tests := []struct {
		name  string
		input ProductInput
	}{
		{name: "blank name", input: ProductInput{ID: "prod_blank_name", Name: " ", Brand: "NGK", SalePriceRupiah: 25000}},
		{name: "blank brand", input: ProductInput{ID: "prod_blank_brand", Name: "Busi", Brand: " ", SalePriceRupiah: 25000}},
		{name: "zero sale price", input: ProductInput{ID: "prod_zero_price", Name: "Busi", Brand: "NGK", SalePriceRupiah: 0}},
		{name: "negative sale price", input: ProductInput{ID: "prod_negative_price", Name: "Busi", Brand: "NGK", SalePriceRupiah: -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewProduct(tt.input); err == nil {
				t.Fatalf("NewProduct() error = nil, want error")
			}
		})
	}
}

func TestNewProductValidatesThresholdPair(t *testing.T) {
	tests := []struct {
		name  string
		input ProductInput
	}{
		{
			name: "reorder without critical",
			input: ProductInput{ID: "prod_reorder_only", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, ReorderPointQty: IntPtr(10)},
		},
		{
			name: "critical without reorder",
			input: ProductInput{ID: "prod_critical_only", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, CriticalThresholdQty: IntPtr(3)},
		},
		{
			name: "negative reorder",
			input: ProductInput{ID: "prod_negative_reorder", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, ReorderPointQty: IntPtr(-1), CriticalThresholdQty: IntPtr(0)},
		},
		{
			name: "critical greater than reorder",
			input: ProductInput{ID: "prod_bad_threshold_order", Name: "Kampas Rem", Brand: "Honda",
				SalePriceRupiah: 40000, ReorderPointQty: IntPtr(5), CriticalThresholdQty: IntPtr(6)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewProduct(tt.input); err == nil {
				t.Fatalf("NewProduct() error = nil, want error")
			}
		})
	}
}
