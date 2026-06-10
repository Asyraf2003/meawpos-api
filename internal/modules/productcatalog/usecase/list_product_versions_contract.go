package usecase

import "time"

type ListProductVersionsQuery struct {
	ProductID string
}

type ListProductVersionsResult struct {
	Items []ListProductVersionItem
}

type ListProductVersionItem struct {
	ProductID        string
	RevisionNo       int
	EventName        string
	ChangedByActorID string
	ChangeReason     string
	ChangedAt        time.Time
}
