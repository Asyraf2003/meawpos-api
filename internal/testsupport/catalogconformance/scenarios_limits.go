// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package catalogconformance

import (
	"context"
	"math"
	"sync"
	"testing"

	catalogusecase "pos-go/internal/modules/catalog/usecase"
)

func scenarioMaximumPrice(t *testing.T, f Fixture) {
	id, amount := newID(), int64(math.MaxInt64)
	mustCreate(t, f, id, "Maximum", &amount)
	read, err := f.Store.GetItem(context.Background(), f.RootID, id)
	if err != nil || read.Price == nil || read.Price.Amount.AmountRupiah() != math.MaxInt64 {
		t.Fatalf("read=%#v err=%v", read, err)
	}
}

func scenarioConcurrentID(t *testing.T, f Fixture) {
	id, amount := newID(), int64(33)
	start := make(chan struct{})
	errs := make(chan error, 2)
	var ready sync.WaitGroup
	ready.Add(2)
	for range 2 {
		go func() {
			ready.Done()
			<-start
			_, err := newCreate(f, id).Execute(context.Background(), catalogusecase.CreateItemCommand{
				RootID: f.RootID, Name: "Concurrent", PriceRupiah: &amount,
			})
			errs <- err
		}()
	}
	ready.Wait()
	close(start)
	var successes int
	for range 2 {
		if err := <-errs; err == nil {
			successes++
		}
	}
	if successes != 1 {
		t.Fatalf("successful concurrent creates=%d, want 1", successes)
	}
	mustCounts(t, f, id, 1, 1)
	read, err := f.Store.GetItem(context.Background(), f.RootID, id)
	if err != nil || read.Price == nil || read.Price.Amount.AmountRupiah() != amount {
		t.Fatalf("final read=%#v err=%v", read, err)
	}
}
