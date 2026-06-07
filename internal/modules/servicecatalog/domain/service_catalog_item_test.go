package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNormalizeNameTrimsWhitespace(t *testing.T) {
	got := NormalizeName("  Potong Rambut  ")
	want := NormalizedName("potong rambut")

	if got != want {
		t.Fatalf("NormalizeName() = %q, want %q", got, want)
	}
}

func TestNormalizeNameCompactsRepeatedWhitespace(t *testing.T) {
	got := NormalizeName("Potong   \t\n  Rambut")
	want := NormalizedName("potong rambut")

	if got != want {
		t.Fatalf("NormalizeName() = %q, want %q", got, want)
	}
}

func TestNormalizeNameLowercases(t *testing.T) {
	got := NormalizeName("PoToNg RAMBUT")
	want := NormalizedName("potong rambut")

	if got != want {
		t.Fatalf("NormalizeName() = %q, want %q", got, want)
	}
}

func TestNewServiceCatalogItemRejectsBlankName(t *testing.T) {
	_, err := NewServiceCatalogItem("svc_1", "   ", 10000, time.Now())
	if !errors.Is(err, ErrBlankServiceCatalogItemName) {
		t.Fatalf("error = %v, want %v", err, ErrBlankServiceCatalogItemName)
	}
}

func TestNewServiceCatalogItemRejectsZeroDefaultPrice(t *testing.T) {
	_, err := NewServiceCatalogItem("svc_1", "Potong Rambut", 0, time.Now())
	if !errors.Is(err, ErrInvalidServiceCatalogItemDefaultPrice) {
		t.Fatalf("error = %v, want %v", err, ErrInvalidServiceCatalogItemDefaultPrice)
	}
}

func TestNewServiceCatalogItemRejectsNegativeDefaultPrice(t *testing.T) {
	_, err := NewServiceCatalogItem("svc_1", "Potong Rambut", -1, time.Now())
	if !errors.Is(err, ErrInvalidServiceCatalogItemDefaultPrice) {
		t.Fatalf("error = %v, want %v", err, ErrInvalidServiceCatalogItemDefaultPrice)
	}
}

func TestNewServiceCatalogItemCreatesActiveItemByDefault(t *testing.T) {
	now := time.Date(2026, 6, 8, 10, 0, 0, 0, time.UTC)

	item, err := NewServiceCatalogItem("svc_1", "  Potong Rambut  ", 10000, now)
	if err != nil {
		t.Fatalf("NewServiceCatalogItem() error = %v", err)
	}

	if !item.IsActive() {
		t.Fatal("new service catalog item should be active by default")
	}

	if item.Status() != ServiceCatalogItemStatusActive {
		t.Fatalf("Status() = %q, want %q", item.Status(), ServiceCatalogItemStatusActive)
	}

	if item.Name() != "Potong Rambut" {
		t.Fatalf("Name() = %q, want %q", item.Name(), "Potong Rambut")
	}

	if item.NormalizedName() != "potong rambut" {
		t.Fatalf("NormalizedName() = %q, want %q", item.NormalizedName(), "potong rambut")
	}

	if item.DefaultPriceRupiah() != 10000 {
		t.Fatalf("DefaultPriceRupiah() = %d, want %d", item.DefaultPriceRupiah(), 10000)
	}
}
