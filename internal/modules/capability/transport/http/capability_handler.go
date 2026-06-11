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
	"context"

	"pos-go/internal/modules/capability/domain"

	"github.com/labstack/echo/v4"
)

type ListCapabilitiesUsecase interface {
	Execute(ctx context.Context) ([]domain.Capability, error)
}

type ShowCapabilityUsecase interface {
	Execute(ctx context.Context, key string) (domain.Capability, error)
}

type EnableCapabilityUsecase interface {
	Execute(ctx context.Context, key string) error
}

type DisableCapabilityUsecase interface {
	Execute(ctx context.Context, key string, reason string) error
}

type CapabilityHandler struct {
	listUsecase    ListCapabilitiesUsecase
	showUsecase    ShowCapabilityUsecase
	enableUsecase  EnableCapabilityUsecase
	disableUsecase DisableCapabilityUsecase
}

func NewCapabilityHandler(
	listUsecase ListCapabilitiesUsecase,
	showUsecase ShowCapabilityUsecase,
	enableUsecase EnableCapabilityUsecase,
	disableUsecase DisableCapabilityUsecase,
) *CapabilityHandler {
	return &CapabilityHandler{
		listUsecase:    listUsecase,
		showUsecase:    showUsecase,
		enableUsecase:  enableUsecase,
		disableUsecase: disableUsecase,
	}
}

func (h *CapabilityHandler) Register(group *echo.Group) {
	group.GET("/capabilities", h.List)
	group.GET("/capabilities/:key", h.Show)
	group.POST("/capabilities/:key/enable", h.Enable)
	group.POST("/capabilities/:key/disable", h.Disable)
}
