// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

package usecase

import "time"

type GoogleStartInput struct {
	Purpose     string
	RedirectURL string
}

type GoogleStartOutput struct {
	RedirectTo string `json:"redirect_to"`
	State      string `json:"state"`
}

type ClientInfo struct {
	UserAgent string
	IP        string
}

type GoogleCallbackInput struct {
	Code        string
	State       string
	RedirectURL string
	Client      ClientInfo
}

type GoogleCallbackOutput struct {
	AccessToken    string    `json:"access_token"`
	AccessExp      time.Time `json:"access_exp"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshExp     time.Time `json:"refresh_exp"`
	TrustLevel     string    `json:"trust_level"`
	StepUpRequired bool      `json:"step_up_required"`
}
