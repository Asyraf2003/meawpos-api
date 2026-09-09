// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestRefreshToken_RepeatedRefreshKeepsAbsoluteSessionExpiry(t *testing.T) {
	originalExpiry := time.Now().Add(24 * time.Hour).Truncate(time.Second)
	repo := &fakeRefreshSessionRepository{session: ports.RefreshSession{
		SessionID: "sess-123", AccountID: "acc-123", ExpiresAt: originalExpiry,
	}}
	usecase := NewRefreshToken(repo, &fakeTokenIssuer{token: "access", exp: time.Now().Add(time.Minute)})

	first, err := usecase.Execute(context.Background(), RefreshTokenInput{RefreshToken: "first"})
	if err != nil {
		t.Fatalf("first refresh error = %v", err)
	}
	if !first.RefreshExp.Equal(originalExpiry) || !repo.lastNewExpiresAt.Equal(originalExpiry) {
		t.Fatalf("first expiry moved: output=%s stored=%s", first.RefreshExp, repo.lastNewExpiresAt)
	}

	second, err := usecase.Execute(context.Background(), RefreshTokenInput{RefreshToken: first.RefreshToken})
	if err != nil {
		t.Fatalf("second refresh error = %v", err)
	}
	if second.RefreshToken == first.RefreshToken {
		t.Fatal("second refresh token was not rotated")
	}
	if !second.RefreshExp.Equal(originalExpiry) || !repo.lastNewExpiresAt.Equal(originalExpiry) {
		t.Fatalf("second expiry moved: output=%s stored=%s", second.RefreshExp, repo.lastNewExpiresAt)
	}
	if repo.findCalls != 2 || repo.rotateCalls != 2 {
		t.Fatalf("find=%d rotate=%d, want 2 each", repo.findCalls, repo.rotateCalls)
	}
}
