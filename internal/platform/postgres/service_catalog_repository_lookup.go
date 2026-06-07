package postgres

import (
	"context"
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

func (r *ServiceCatalogRepository) Lookup(
	ctx context.Context,
	filter ports.LookupServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	args := []any{}
	conditions := []string{}
	nextArg := 1

	if strings.TrimSpace(filter.Query) != "" {
		conditions = append(conditions, "normalized_name LIKE $"+itoa(nextArg))
		args = append(args, "%"+domain.NormalizeName(filter.Query)+"%")
		nextArg++
	}

	if filter.ActiveOnly {
		conditions = append(conditions, "is_active = true")
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 20
	}

	args = append(args, limit)

	sql := serviceCatalogItemSelectSQL()
	if len(conditions) > 0 {
		sql += "\n\t\tWHERE " + strings.Join(conditions, " AND ")
	}
	sql += "\n\t\tORDER BY normalized_name, id"
	sql += "\n\t\tLIMIT $" + itoa(nextArg)

	return r.findMany(ctx, sql, args...)
}
