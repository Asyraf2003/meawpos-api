package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

type CreateProductCommand struct {
	Code                 string
	Name                 string
	Brand                string
	Size                 *int
	SalePriceRupiah      int64
	ReorderPointQty      *int
	CriticalThresholdQty *int
	ActorID              string
	Reason               string
}

type CreateProductResult struct {
	ID                   string
	Code                 *string
	Name                 string
	NormalizedName       string
	Brand                string
	NormalizedBrand      string
	Size                 *int
	SalePriceRupiah      int64
	ReorderPointQty      *int
	CriticalThresholdQty *int
	Status               string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type ProductIDGenerator interface {
	NewProductID() (string, error)
}

type CreateProduct struct {
	repository       ports.ProductRepository
	duplicateChecker ports.ProductDuplicateChecker
	idGenerator      ProductIDGenerator
	now              func() time.Time
}

func NewCreateProduct(
	repository ports.ProductRepository,
	duplicateChecker ports.ProductDuplicateChecker,
	idGenerator ProductIDGenerator,
	now func() time.Time,
) *CreateProduct {
	return &CreateProduct{
		repository:       repository,
		duplicateChecker: duplicateChecker,
		idGenerator:      idGenerator,
		now:              now,
	}
}

func (uc *CreateProduct) Execute(
	ctx context.Context,
	cmd CreateProductCommand,
) (CreateProductResult, error) {
	id, err := uc.idGenerator.NewProductID()
	if err != nil {
		return CreateProductResult{}, err
	}

	product, err := domain.NewProduct(domain.ProductInput{
		ID:                   id,
		Code:                 cmd.Code,
		Name:                 cmd.Name,
		Brand:                cmd.Brand,
		Size:                 cmd.Size,
		SalePriceRupiah:      cmd.SalePriceRupiah,
		ReorderPointQty:      cmd.ReorderPointQty,
		CriticalThresholdQty: cmd.CriticalThresholdQty,
	})
	if err != nil {
		return CreateProductResult{}, err
	}

	if err := uc.duplicateChecker.CheckCreateDuplicate(ctx, ports.ProductDuplicateCandidate{
		Code:            product.Code(),
		NormalizedName:  product.NormalizedName(),
		NormalizedBrand: product.NormalizedBrand(),
		Size:            product.Size(),
	}); err != nil {
		return CreateProductResult{}, err
	}

	if err := uc.repository.Create(ctx, product); err != nil {
		return CreateProductResult{}, err
	}

	now := uc.now()

	return CreateProductResult{
		ID:                   product.ID(),
		Code:                 product.Code(),
		Name:                 product.Name(),
		NormalizedName:       product.NormalizedName(),
		Brand:                product.Brand(),
		NormalizedBrand:      product.NormalizedBrand(),
		Size:                 product.Size(),
		SalePriceRupiah:      product.SalePriceRupiah(),
		ReorderPointQty:      product.ReorderPointQty(),
		CriticalThresholdQty: product.CriticalThresholdQty(),
		Status:               string(product.Status()),
		CreatedAt:            now,
		UpdatedAt:            now,
	}, nil
}
