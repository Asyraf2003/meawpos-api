package usecase

import (
	"context"
	"testing"
)

func TestListProductsForwardsQueryToReader(t *testing.T) {
	reader := &listProductsReaderDouble{}
	usecase := NewListProducts(reader)

	_, err := usecase.Execute(context.Background(), ListProductsQuery{
		Search:  "kopi",
		Status:  "active",
		Page:    2,
		PerPage: 25,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if reader.capturedQuery.Search != "kopi" {
		t.Fatalf("captured query Search = %q, want %q", reader.capturedQuery.Search, "kopi")
	}

	if reader.capturedQuery.Status != "active" {
		t.Fatalf("captured query Status = %q, want %q", reader.capturedQuery.Status, "active")
	}

	if reader.capturedQuery.Page != 2 {
		t.Fatalf("captured query Page = %d, want %d", reader.capturedQuery.Page, 2)
	}

	if reader.capturedQuery.PerPage != 25 {
		t.Fatalf("captured query PerPage = %d, want %d", reader.capturedQuery.PerPage, 25)
	}
}
