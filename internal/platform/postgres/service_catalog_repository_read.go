package postgres

import (
	"context"
	"errors"

	"pos-go/internal/modules/servicecatalog/domain"

	"github.com/jackc/pgx/v5"
)

func (r *ServiceCatalogRepository) FindByID(
	ctx context.Context,
	id domain.ServiceCatalogItemID,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE id = $1
	`, string(id))

	return scanOptionalServiceCatalogItem(row)
}

func (r *ServiceCatalogRepository) FindByNormalizedName(
	ctx context.Context,
	normalizedName domain.NormalizedName,
) (domain.ServiceCatalogItem, bool, error) {
	row := r.queryRow(ctx, serviceCatalogItemSelectSQL()+`
		WHERE normalized_name = $1
	`, string(normalizedName))

	return scanOptionalServiceCatalogItem(row)
}

func scanOptionalServiceCatalogItem(
	row serviceCatalogItemScanner,
) (domain.ServiceCatalogItem, bool, error) {
	item, err := scanServiceCatalogItem(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ServiceCatalogItem{}, false, nil
	}
	if err != nil {
		return domain.ServiceCatalogItem{}, false, err
	}

	return item, true, nil
}

func (r *ServiceCatalogRepository) findMany(
	ctx context.Context,
	sql string,
	args ...any,
) ([]domain.ServiceCatalogItem, error) {
	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []domain.ServiceCatalogItem{}
	for rows.Next() {
		item, err := scanServiceCatalogItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}
