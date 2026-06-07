package postgres

import (
	"context"
	"errors"

	"pos-go/internal/modules/servicecatalog/domain"

	"github.com/jackc/pgx/v5"
)

func (r *ServiceCatalogRepository) Create(ctx context.Context, item domain.ServiceCatalogItem) error {
	sql := `
		INSERT INTO service_catalog_items (
			id, name, normalized_name, default_price_rupiah,
			is_active, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.exec(ctx, sql, serviceCatalogItemArgs(item)...)
	return err
}

func (r *ServiceCatalogRepository) Update(ctx context.Context, item domain.ServiceCatalogItem) error {
	sql := `
		UPDATE service_catalog_items
		SET
			name = $2,
			normalized_name = $3,
			default_price_rupiah = $4,
			is_active = $5,
			updated_at = $7
		WHERE id = $1
	`

	_, err := r.exec(ctx, sql, serviceCatalogItemArgs(item)...)
	return err
}

func (r *ServiceCatalogRepository) SetActive(
	ctx context.Context,
	id domain.ServiceCatalogItemID,
	active bool,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE id = $1
		FOR UPDATE
	`, string(id))

	item, err := scanServiceCatalogItem(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ServiceCatalogItem{}, false, nil
	}
	if err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	if active {
		item.Activate(nowUTC())
	} else {
		item.Deactivate(nowUTC())
	}

	if err := r.Update(ctx, item); err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	return item, true, nil
}
