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

package http

import (
	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type fakeGoogleFlow struct {
	startInput    authusecase.GoogleStartInput
	startOutput   authusecase.GoogleStartOutput
	startErr      error
	callbackInput authusecase.GoogleCallbackInput
	callbackOut   authusecase.GoogleCallbackOutput
	callbackErr   error
}

func (f *fakeGoogleFlow) GoogleStart(ctx echo.Context, in authusecase.GoogleStartInput) (authusecase.GoogleStartOutput, error) {
	_ = ctx
	f.startInput = in
	return f.startOutput, f.startErr
}

func (f *fakeGoogleFlow) GoogleCallback(ctx echo.Context, in authusecase.GoogleCallbackInput) (authusecase.GoogleCallbackOutput, error) {
	_ = ctx
	f.callbackInput = in
	return f.callbackOut, f.callbackErr
}
