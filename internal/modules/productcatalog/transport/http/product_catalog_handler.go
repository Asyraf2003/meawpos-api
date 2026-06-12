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

type ProductCatalogHandler struct {
	list       listProducts
	lookup     lookupProducts
	show       getProductDetail
	create     createProduct
	update     updateProduct
	softDelete softDeleteProduct
	restore    restoreProduct
	versions   listProductVersions
}

func NewProductCatalogHandler(
	list listProducts,
	lookup lookupProducts,
	show getProductDetail,
	create createProduct,
	update updateProduct,
	softDelete softDeleteProduct,
	restore restoreProduct,
	versions listProductVersions,
) ProductCatalogHandler {
	return ProductCatalogHandler{
		list:       list,
		lookup:     lookup,
		show:       show,
		create:     create,
		update:     update,
		softDelete: softDelete,
		restore:    restore,
		versions:   versions,
	}
}

func (h ProductCatalogHandler) RegisterList(group *echo.Group) {
	group.GET("/products", h.List)
}

func (h ProductCatalogHandler) RegisterCreate(group *echo.Group) {
	group.POST("/products", h.Create)
}

func (h ProductCatalogHandler) RegisterLookup(group *echo.Group) {
	group.GET("/products/lookup", h.Lookup)
}

func (h ProductCatalogHandler) RegisterVersions(group *echo.Group) {
	group.GET("/products/:id/versions", h.Versions)
}

func (h ProductCatalogHandler) RegisterRestore(group *echo.Group) {
	group.PATCH("/products/:id/restore", h.Restore)
}

func (h ProductCatalogHandler) RegisterShow(group *echo.Group) {
	group.GET("/products/:id", h.Show)
}

func (h ProductCatalogHandler) RegisterUpdate(group *echo.Group) {
	group.PUT("/products/:id", h.Update)
}

func (h ProductCatalogHandler) RegisterDelete(group *echo.Group) {
	group.DELETE("/products/:id", h.Delete)
}
