package usecase

import "errors"

var (
	ErrServiceCatalogItemNotFound                = errors.New("service catalog item not found")
	ErrDuplicateServiceCatalogItemNormalizedName = errors.New("service catalog item normalized name already exists")
	ErrMissingServiceCatalogItemIDGenerator      = errors.New("service catalog item id generator is required")
	ErrInvalidLookupLimit                        = errors.New("service catalog item lookup limit must be between 1 and 50")
)
