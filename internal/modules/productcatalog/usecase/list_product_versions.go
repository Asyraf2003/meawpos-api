package usecase

import "pos-go/internal/modules/productcatalog/ports"

type ListProductVersions struct {
	versions ports.ProductVersionRepository
}

func NewListProductVersions(versions ports.ProductVersionRepository) *ListProductVersions {
	return &ListProductVersions{
		versions: versions,
	}
}
