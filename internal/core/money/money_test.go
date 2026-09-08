// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package money

import (
	"errors"
	"math"
	"testing"
)

func TestCheckedArithmetic(t *testing.T) {
	sum, err := Add(IDR(18_000), IDR(55_000))
	if err != nil || sum.AmountRupiah() != 73_000 {
		t.Fatalf("Add() = %d, %v", sum.AmountRupiah(), err)
	}
	product, err := Multiply(IDR(18_000), 2)
	if err != nil || product.AmountRupiah() != 36_000 {
		t.Fatalf("Multiply() = %d, %v", product.AmountRupiah(), err)
	}
	change, err := Subtract(IDR(100_000), IDR(73_000))
	if err != nil || change.AmountRupiah() != 27_000 {
		t.Fatalf("Subtract() = %d, %v", change.AmountRupiah(), err)
	}
}

func TestArithmeticRejectsOverflow(t *testing.T) {
	checks := []func() error{
		func() error { _, err := Add(IDR(math.MaxInt64), IDR(1)); return err },
		func() error { _, err := Subtract(IDR(0), IDR(math.MinInt64)); return err },
		func() error { _, err := Multiply(IDR(math.MaxInt64), 2); return err },
	}
	for _, check := range checks {
		if err := check(); !errors.Is(err, ErrOverflow) {
			t.Fatalf("error = %v, want ErrOverflow", err)
		}
	}
}
