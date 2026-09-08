// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package root

import (
	"errors"
	"net/http"

	rootdomain "pos-go/internal/core/root/domain"
	httpresponse "pos-go/internal/transport/http/response"
)

func present(root rootdomain.Root) rootResponse {
	return rootResponse{ID: root.ID, Name: root.Name, PrimaryOwnerAccountID: root.PrimaryOwnerAccountID, CreatedAt: root.CreatedAt.Format("2006-01-02T15:04:05.000000Z07:00")}
}

func mapError(err error) error {
	if errors.Is(err, rootdomain.ErrRootNameRequired) {
		return httpresponse.NewHTTPError(http.StatusBadRequest, "root_name_required", err.Error())
	}
	return err
}
