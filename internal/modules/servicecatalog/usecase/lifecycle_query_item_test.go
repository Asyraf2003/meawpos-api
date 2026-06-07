package usecase

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/servicecatalog/ports"
)

func TestActivateServiceCatalogItemMarksInactiveItemActive(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)

	if _, _, err := repo.SetActive(ctx, "svc_1", false); err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}

	uc := NewActivateServiceCatalogItem(repo)

	got, err := uc.Execute(ctx, ActivateServiceCatalogItemCommand{ID: "svc_1"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if !got.IsActive {
		t.Fatal("activated item should be active")
	}
}

func TestDeactivateServiceCatalogItemMarksActiveItemInactive(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)

	uc := NewDeactivateServiceCatalogItem(repo)

	got, err := uc.Execute(ctx, DeactivateServiceCatalogItemCommand{ID: "svc_1"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if got.IsActive {
		t.Fatal("deactivated item should be inactive")
	}
}

func TestShowServiceCatalogItemMissingItemReturnsNotFound(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	uc := NewShowServiceCatalogItem(repo)

	_, err := uc.Execute(ctx, ShowServiceCatalogItemCommand{ID: "svc_missing"})
	if !errors.Is(err, ErrServiceCatalogItemNotFound) {
		t.Fatalf("error = %v, want %v", err, ErrServiceCatalogItemNotFound)
	}
}

func TestListServiceCatalogItemsFiltersActiveInactiveAndAll(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)
	seedServiceCatalogItem(t, repo, "svc_2", "Cuci Motor", 15000)

	if _, _, err := repo.SetActive(ctx, "svc_2", false); err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}

	uc := NewListServiceCatalogItems(repo)

	active, err := uc.Execute(ctx, ListServiceCatalogItemsCommand{Status: ports.ListStatusActive})
	if err != nil {
		t.Fatalf("active Execute() error = %v", err)
	}

	if len(active) != 1 || active[0].ID != "svc_1" {
		t.Fatalf("active result = %+v, want only svc_1", active)
	}

	inactive, err := uc.Execute(ctx, ListServiceCatalogItemsCommand{Status: ports.ListStatusInactive})
	if err != nil {
		t.Fatalf("inactive Execute() error = %v", err)
	}

	if len(inactive) != 1 || inactive[0].ID != "svc_2" {
		t.Fatalf("inactive result = %+v, want only svc_2", inactive)
	}

	all, err := uc.Execute(ctx, ListServiceCatalogItemsCommand{Status: ports.ListStatusAll})
	if err != nil {
		t.Fatalf("all Execute() error = %v", err)
	}

	if len(all) != 2 {
		t.Fatalf("all result count = %d, want %d", len(all), 2)
	}
}

func TestLookupServiceCatalogItemsExcludesInactiveByDefault(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Cuci Motor", 15000)
	seedServiceCatalogItem(t, repo, "svc_2", "Cuci Mobil", 25000)

	if _, _, err := repo.SetActive(ctx, "svc_2", false); err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}

	uc := NewLookupServiceCatalogItems(repo)

	got, err := uc.Execute(ctx, LookupServiceCatalogItemsCommand{
		Query: "cuci",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(got) != 1 || got[0].ID != "svc_1" {
		t.Fatalf("lookup result = %+v, want only active svc_1", got)
	}
}

func TestLookupServiceCatalogItemsEnforcesMaxLimit(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	uc := NewLookupServiceCatalogItems(repo)

	_, err := uc.Execute(ctx, LookupServiceCatalogItemsCommand{
		Limit: 51,
	})
	if !errors.Is(err, ErrInvalidLookupLimit) {
		t.Fatalf("error = %v, want %v", err, ErrInvalidLookupLimit)
	}
}
