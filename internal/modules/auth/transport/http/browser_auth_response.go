// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"time"

	authusecase "pos-go/internal/modules/auth/usecase"
)

type browserAuthResponse struct {
	AccessToken    string    `json:"access_token"`
	AccessExp      time.Time `json:"access_exp"`
	SessionExp     time.Time `json:"session_exp"`
	TrustLevel     string    `json:"trust_level"`
	StepUpRequired bool      `json:"step_up_required"`
}

func browserResponseFromManual(out authusecase.ManualLoginOutput) browserAuthResponse {
	return browserAuthResponse{
		AccessToken:    out.AccessToken,
		AccessExp:      out.AccessExp,
		SessionExp:     out.RefreshExp,
		TrustLevel:     out.TrustLevel,
		StepUpRequired: out.StepUpRequired,
	}
}

func browserResponseFromRefresh(out authusecase.RefreshTokenOutput) browserAuthResponse {
	return browserAuthResponse{
		AccessToken:    out.AccessToken,
		AccessExp:      out.AccessExp,
		SessionExp:     out.RefreshExp,
		TrustLevel:     out.TrustLevel,
		StepUpRequired: out.StepUpRequired,
	}
}
