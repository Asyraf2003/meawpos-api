// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	rootusecase "pos-go/internal/core/root/usecase"

	"github.com/google/uuid"
)

func TestRootCreation_CommitsOwnerMembershipAndAudit(t *testing.T) {
	ctx := context.Background()
	pool := mustOpenIntegrationPool(t, ctx)
	accountID, sessionID := uuid.NewString(), uuid.NewString()
	if _, err := pool.Exec(ctx, `INSERT INTO accounts (id,email) VALUES ($1,$2)`, accountID, accountID+"@test.invalid"); err != nil {
		t.Fatal(err)
	}
	var rootID string
	t.Cleanup(func() {
		if rootID != "" {
			_, _ = pool.Exec(ctx, `DELETE FROM audit_events WHERE root_id=$1`, rootID)
			_, _ = pool.Exec(ctx, `DELETE FROM roots WHERE id=$1`, rootID)
		}
		_, _ = pool.Exec(ctx, `DELETE FROM accounts WHERE id=$1`, accountID)
		pool.Close()
	})
	uc := rootusecase.NewCreateRoot(NewRootStore(pool), NewAuditWriter(pool), NewTransactor(pool), uuid.NewString, time.Now)
	root, err := uc.Execute(ctx, rootusecase.CreateRootCommand{Name: "Root", AccountID: accountID, SessionID: sessionID, RequestID: uuid.NewString()})
	rootID = root.ID
	if err != nil {
		t.Fatal(err)
	}
	var memberships, audits int
	err = pool.QueryRow(ctx, `SELECT (SELECT count(*) FROM root_memberships WHERE root_id=$1 AND account_id=$2),(SELECT count(*) FROM audit_events WHERE root_id=$1 AND operation='root.create')`, rootID, accountID).Scan(&memberships, &audits)
	if err != nil || memberships != 1 || audits != 1 {
		t.Fatalf("counts=%d,%d err=%v", memberships, audits, err)
	}
}

func TestRootCreation_AuditFailureRollsBackRootAndMembership(t *testing.T) {
	ctx := context.Background()
	pool := mustOpenIntegrationPool(t, ctx)
	accountID, rootID := uuid.NewString(), uuid.NewString()
	if _, err := pool.Exec(ctx, `INSERT INTO accounts (id,email) VALUES ($1,$2)`, accountID, accountID+"@test.invalid"); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _, _ = pool.Exec(ctx, `DELETE FROM accounts WHERE id=$1`, accountID); pool.Close() })
	ids := []string{rootID, uuid.NewString(), uuid.NewString()}
	uc := rootusecase.NewCreateRoot(NewRootStore(pool), failingAuditWriter{}, NewTransactor(pool), func() string { id := ids[0]; ids = ids[1:]; return id }, time.Now)
	if _, err := uc.Execute(ctx, rootusecase.CreateRootCommand{Name: "Rollback", AccountID: accountID, SessionID: uuid.NewString(), RequestID: uuid.NewString()}); err == nil {
		t.Fatal("expected audit failure")
	}
	var roots, memberships int
	if err := pool.QueryRow(ctx, `SELECT (SELECT count(*) FROM roots WHERE id=$1),(SELECT count(*) FROM root_memberships WHERE root_id=$1)`, rootID).Scan(&roots, &memberships); err != nil || roots != 0 || memberships != 0 {
		t.Fatalf("counts=%d,%d err=%v", roots, memberships, err)
	}
}
