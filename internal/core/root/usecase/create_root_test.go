// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/core/audit"
	rootdomain "pos-go/internal/core/root/domain"
)

type rootStoreFake struct{ roots []rootdomain.Root }

func (f *rootStoreFake) CreateRoot(_ context.Context, root rootdomain.Root, _ string) error {
	f.roots = append(f.roots, root)
	return nil
}
func (f *rootStoreFake) ListRootsForAccount(context.Context, string) ([]rootdomain.Root, error) {
	return f.roots, nil
}

type auditFake struct {
	err    error
	events []audit.Event
}

func (f *auditFake) WriteAuditEvent(_ context.Context, event audit.Event) error {
	f.events = append(f.events, event)
	return f.err
}

type txFake struct{}

func (txFake) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestCreateRootCreatesOwnerMembershipAndAuditIntent(t *testing.T) {
	store, writer := &rootStoreFake{}, &auditFake{}
	ids := []string{"root-id", "membership-id", "audit-id"}
	uc := NewCreateRoot(store, writer, txFake{}, func() string { id := ids[0]; ids = ids[1:]; return id }, func() time.Time { return time.Unix(1, 0) })
	got, err := uc.Execute(context.Background(), CreateRootCommand{Name: "  Toko A  ", AccountID: "account", SessionID: "session", RequestID: "request"})
	if err != nil || got.Name != "Toko A" || len(writer.events) != 1 {
		t.Fatalf("Execute() = %#v, %v; audits=%d", got, err, len(writer.events))
	}
}

func TestCreateRootPropagatesAuditFailure(t *testing.T) {
	writer := &auditFake{err: errors.New("audit failed")}
	uc := NewCreateRoot(&rootStoreFake{}, writer, txFake{}, func() string { return "id" }, time.Now)
	_, err := uc.Execute(context.Background(), CreateRootCommand{Name: "Root", AccountID: "account"})
	if err == nil {
		t.Fatal("Execute() error = nil")
	}
}
