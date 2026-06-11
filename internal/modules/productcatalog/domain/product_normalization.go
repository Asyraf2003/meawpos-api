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

import "strings"

func normalizeCode(code string) *string {
	trimmed := strings.TrimSpace(code)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

func normalizeDisplayText(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}

func normalizeSearchText(value string) string {
	return strings.ToLower(normalizeDisplayText(value))
}

func validateThreshold(reorderPointQty, criticalThresholdQty *int) error {
	if (reorderPointQty == nil) != (criticalThresholdQty == nil) {
		return ErrProductThresholdPairRequired
	}

	if reorderPointQty == nil {
		return nil
	}

	if *reorderPointQty < 0 || *criticalThresholdQty < 0 {
		return ErrProductThresholdNegative
	}

	if *criticalThresholdQty > *reorderPointQty {
		return ErrProductCriticalAboveReorder
	}

	return nil
}

func copyIntPtr(value *int) *int {
	if value == nil {
		return nil
	}
	copied := *value
	return &copied
}
