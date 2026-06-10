package usecase

import "time"

type RestoreProductCommand struct {
	ID      string
	ActorID string
	Reason  string
}

type RestoreProductResult struct {
	ID         string
	Status     string
	RestoredAt time.Time
	RevisionNo int
}
