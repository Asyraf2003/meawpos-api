package domain

import (
	"strings"
	"time"
)

type ServiceCatalogItemID string

type ServiceCatalogItemStatus string

const (
	ServiceCatalogItemStatusActive   ServiceCatalogItemStatus = "active"
	ServiceCatalogItemStatusInactive ServiceCatalogItemStatus = "inactive"
)

type NormalizedName string

type MoneyRupiah int64

type ServiceCatalogItem struct {
	id                 ServiceCatalogItemID
	name               string
	normalizedName     NormalizedName
	defaultPriceRupiah MoneyRupiah
	isActive           bool
	createdAt          time.Time
	updatedAt          time.Time
}

func NewServiceCatalogItem(
	id ServiceCatalogItemID,
	name string,
	defaultPriceRupiah MoneyRupiah,
	now time.Time,
) (ServiceCatalogItem, error) {
	id = ServiceCatalogItemID(strings.TrimSpace(string(id)))

	if err := ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItem{}, err
	}

	if err := ValidateServiceCatalogItemName(name); err != nil {
		return ServiceCatalogItem{}, err
	}

	if err := ValidateServiceCatalogItemDefaultPrice(defaultPriceRupiah); err != nil {
		return ServiceCatalogItem{}, err
	}

	return ServiceCatalogItem{
		id:                 id,
		name:               normalizeDisplayName(name),
		normalizedName:     NormalizeName(name),
		defaultPriceRupiah: defaultPriceRupiah,
		isActive:           true,
		createdAt:          now,
		updatedAt:          now,
	}, nil
}

func (i ServiceCatalogItem) ID() ServiceCatalogItemID {
	return i.id
}

func (i ServiceCatalogItem) Name() string {
	return i.name
}

func (i ServiceCatalogItem) NormalizedName() NormalizedName {
	return i.normalizedName
}

func (i ServiceCatalogItem) DefaultPriceRupiah() MoneyRupiah {
	return i.defaultPriceRupiah
}

func (i ServiceCatalogItem) IsActive() bool {
	return i.isActive
}

func (i ServiceCatalogItem) Status() ServiceCatalogItemStatus {
	if i.isActive {
		return ServiceCatalogItemStatusActive
	}

	return ServiceCatalogItemStatusInactive
}

func (i ServiceCatalogItem) CreatedAt() time.Time {
	return i.createdAt
}

func (i ServiceCatalogItem) UpdatedAt() time.Time {
	return i.updatedAt
}

func (i *ServiceCatalogItem) Update(
	name string,
	defaultPriceRupiah MoneyRupiah,
	now time.Time,
) error {
	if err := ValidateServiceCatalogItemName(name); err != nil {
		return err
	}

	if err := ValidateServiceCatalogItemDefaultPrice(defaultPriceRupiah); err != nil {
		return err
	}

	i.name = normalizeDisplayName(name)
	i.normalizedName = NormalizeName(name)
	i.defaultPriceRupiah = defaultPriceRupiah
	i.updatedAt = now

	return nil
}

func (i *ServiceCatalogItem) Activate(now time.Time) {
	i.isActive = true
	i.updatedAt = now
}

func (i *ServiceCatalogItem) Deactivate(now time.Time) {
	i.isActive = false
	i.updatedAt = now
}
