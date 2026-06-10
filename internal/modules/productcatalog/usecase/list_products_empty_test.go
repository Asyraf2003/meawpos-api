package usecase

import (
	"context"
	"testing"
)

func TestListProductsReturnsEmptyItemsForEmptyReaderList(t *testing.T) {
	usecase := NewListProducts(&listProductsReaderDouble{})

	result, err := usecase.Execute(context.Background(), ListProductsQuery{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 0 {
		t.Fatalf("result Items length = %d, want %d", len(result.Items), 0)
	}
}
