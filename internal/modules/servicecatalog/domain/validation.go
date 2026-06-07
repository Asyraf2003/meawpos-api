package domain

import "strings"

func ValidateServiceCatalogItemID(id ServiceCatalogItemID) error {
	if strings.TrimSpace(string(id)) == "" {
		return ErrInvalidServiceCatalogItemID
	}

	return nil
}

func ValidateServiceCatalogItemName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrBlankServiceCatalogItemName
	}

	return nil
}

func ValidateServiceCatalogItemDefaultPrice(price MoneyRupiah) error {
	if price <= 0 {
		return ErrInvalidServiceCatalogItemDefaultPrice
	}

	return nil
}
