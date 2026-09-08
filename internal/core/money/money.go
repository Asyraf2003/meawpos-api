// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.

package money

import "errors"

const CurrencyIDR = "IDR"

var ErrOverflow = errors.New("money overflow")

type Money struct {
	amount int64
}

func IDR(amountRupiah int64) Money {
	return Money{amount: amountRupiah}
}

func (m Money) AmountRupiah() int64 {
	return m.amount
}

func (m Money) Currency() string {
	return CurrencyIDR
}
