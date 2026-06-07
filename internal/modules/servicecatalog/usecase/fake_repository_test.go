package usecase

import (
	"context"
	"sort"
	"strings"
	"testing"
	"time"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

type fakeServiceCatalogRepository struct {
	items map[domain.ServiceCatalogItemID]domain.ServiceCatalogItem
}

func newFakeServiceCatalogRepository() *fakeServiceCatalogRepository {
	return &fakeServiceCatalogRepository{
		items: make(map[domain.ServiceCatalogItemID]domain.ServiceCatalogItem),
	}
}

func (r *fakeServiceCatalogRepository) Create(_ context.Context, item domain.ServiceCatalogItem) error {
	r.items[item.ID()] = item
	return nil
}

func (r *fakeServiceCatalogRepository) Update(_ context.Context, item domain.ServiceCatalogItem) error {
	r.items[item.ID()] = item
	return nil
}

func (r *fakeServiceCatalogRepository) FindByID(
	_ context.Context,
	id domain.ServiceCatalogItemID,
) (domain.ServiceCatalogItem, bool, error) {
	item, found := r.items[id]
	return item, found, nil
}

func (r *fakeServiceCatalogRepository) FindByNormalizedName(
	_ context.Context,
	normalizedName domain.NormalizedName,
) (domain.ServiceCatalogItem, bool, error) {
	for _, item := range r.items {
		if item.NormalizedName() == normalizedName {
			return item, true, nil
		}
	}

	return domain.ServiceCatalogItem{}, false, nil
}

func (r *fakeServiceCatalogRepository) List(
	_ context.Context,
	filter ports.ListServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	items := make([]domain.ServiceCatalogItem, 0, len(r.items))
	query := string(domain.NormalizeName(filter.Query))

	for _, item := range r.items {
		if !matchesStatus(item, filter.Status) {
			continue
		}

		if query != "" && !strings.Contains(string(item.NormalizedName()), query) {
			continue
		}

		items = append(items, item)
	}

	sortItemsByNormalizedName(items)

	return items, nil
}

func (r *fakeServiceCatalogRepository) Lookup(
	_ context.Context,
	filter ports.LookupServiceCatalogItemsFilter,
) ([]domain.ServiceCatalogItem, error) {
	items := make([]domain.ServiceCatalogItem, 0, len(r.items))
	query := string(domain.NormalizeName(filter.Query))

	for _, item := range r.items {
		if filter.ActiveOnly && !item.IsActive() {
			continue
		}

		if query != "" && !strings.Contains(string(item.NormalizedName()), query) {
			continue
		}

		items = append(items, item)
	}

	sortItemsByNormalizedName(items)

	if filter.Limit > 0 && len(items) > filter.Limit {
		return items[:filter.Limit], nil
	}

	return items, nil
}

func (r *fakeServiceCatalogRepository) SetActive(
	_ context.Context,
	id domain.ServiceCatalogItemID,
	active bool,
) (domain.ServiceCatalogItem, bool, error) {
	item, found := r.items[id]
	if !found {
		return domain.ServiceCatalogItem{}, false, nil
	}

	if active {
		item.Activate(time.Now())
	} else {
		item.Deactivate(time.Now())
	}

	r.items[id] = item

	return item, true, nil
}

func matchesStatus(item domain.ServiceCatalogItem, status ports.ListStatusFilter) bool {
	switch status {
	case ports.ListStatusInactive:
		return !item.IsActive()
	case ports.ListStatusAll:
		return true
	case ports.ListStatusActive, "":
		return item.IsActive()
	default:
		return item.IsActive()
	}
}

func sortItemsByNormalizedName(items []domain.ServiceCatalogItem) {
	sort.Slice(items, func(left, right int) bool {
		return items[left].NormalizedName() < items[right].NormalizedName()
	})
}

func seedServiceCatalogItem(
	t *testing.T,
	repo *fakeServiceCatalogRepository,
	id domain.ServiceCatalogItemID,
	name string,
	price domain.MoneyRupiah,
) domain.ServiceCatalogItem {
	t.Helper()

	item, err := domain.NewServiceCatalogItem(id, name, price, fixedNow())
	if err != nil {
		t.Fatalf("NewServiceCatalogItem() error = %v", err)
	}

	if err := repo.Create(context.Background(), item); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	return item
}

func fixedIDGenerator(id domain.ServiceCatalogItemID) ServiceCatalogItemIDGenerator {
	return func() (domain.ServiceCatalogItemID, error) {
		return id, nil
	}
}

func fixedClock() time.Time {
	return time.Date(2026, 6, 8, 10, 0, 0, 0, time.UTC)
}

func fixedNow() time.Time {
	return fixedClock()
}
