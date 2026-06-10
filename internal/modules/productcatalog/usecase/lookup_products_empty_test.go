package usecase

import (
	"context"
	"testing"
)

func TestLookupProductsReturnsEmptyItems(t *testing.T) {
	usecase := NewLookupProducts(&lookupProductsReaderDouble{})

	result, err := usecase.Execute(context.Background(), LookupProductsQuery{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 0 {
		t.Fatalf("Items length = %d, want 0", len(result.Items))
	}
	if result.Items == nil {
		t.Fatalf("Items = nil, want empty slice")
	}
}
