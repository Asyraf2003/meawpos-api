// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package money

import "math"

func Add(left, right Money) (Money, error) {
	if right.amount > 0 && left.amount > math.MaxInt64-right.amount {
		return Money{}, ErrOverflow
	}
	if right.amount < 0 && left.amount < math.MinInt64-right.amount {
		return Money{}, ErrOverflow
	}
	return IDR(left.amount + right.amount), nil
}

func Subtract(left, right Money) (Money, error) {
	if right.amount == math.MinInt64 {
		if left.amount >= 0 {
			return Money{}, ErrOverflow
		}
		return IDR(left.amount - right.amount), nil
	}
	return Add(left, IDR(-right.amount))
}

func Multiply(value Money, quantity int64) (Money, error) {
	if value.amount == 0 || quantity == 0 {
		return IDR(0), nil
	}
	if value.amount == math.MinInt64 && quantity == -1 ||
		quantity == math.MinInt64 && value.amount == -1 {
		return Money{}, ErrOverflow
	}
	result := value.amount * quantity
	if result/quantity != value.amount {
		return Money{}, ErrOverflow
	}
	return IDR(result), nil
}
