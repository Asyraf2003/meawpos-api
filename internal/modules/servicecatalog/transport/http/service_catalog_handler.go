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

import "github.com/labstack/echo/v4"

type ServiceCatalogHandler struct {
	list       listServiceCatalogItems
	lookup     lookupServiceCatalogItems
	show       showServiceCatalogItem
	create     createServiceCatalogItem
	update     updateServiceCatalogItem
	activate   activateServiceCatalogItem
	deactivate deactivateServiceCatalogItem
}

func NewServiceCatalogHandler(
	list listServiceCatalogItems,
	lookup lookupServiceCatalogItems,
	show showServiceCatalogItem,
	create createServiceCatalogItem,
	update updateServiceCatalogItem,
	activate activateServiceCatalogItem,
	deactivate deactivateServiceCatalogItem,
) ServiceCatalogHandler {
	return ServiceCatalogHandler{
		list:       list,
		lookup:     lookup,
		show:       show,
		create:     create,
		update:     update,
		activate:   activate,
		deactivate: deactivate,
	}
}

func (h ServiceCatalogHandler) RegisterList(group *echo.Group) {
	group.GET("/items", h.List)
}

func (h ServiceCatalogHandler) RegisterCreate(group *echo.Group) {
	group.POST("/items", h.Create)
}

func (h ServiceCatalogHandler) RegisterLookup(group *echo.Group) {
	group.GET("/items/lookup", h.Lookup)
}

func (h ServiceCatalogHandler) RegisterShow(group *echo.Group) {
	group.GET("/items/:id", h.Show)
}

func (h ServiceCatalogHandler) RegisterUpdate(group *echo.Group) {
	group.PUT("/items/:id", h.Update)
}

func (h ServiceCatalogHandler) RegisterActivate(group *echo.Group) {
	group.POST("/items/:id/activate", h.Activate)
}

func (h ServiceCatalogHandler) RegisterDeactivate(group *echo.Group) {
	group.POST("/items/:id/deactivate", h.Deactivate)
}
