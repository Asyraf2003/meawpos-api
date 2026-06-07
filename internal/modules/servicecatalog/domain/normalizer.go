package domain

import "strings"

func NormalizeName(name string) NormalizedName {
	parts := strings.Fields(strings.TrimSpace(name))
	return NormalizedName(strings.ToLower(strings.Join(parts, " ")))
}

func normalizeDisplayName(name string) string {
	return strings.TrimSpace(name)
}
