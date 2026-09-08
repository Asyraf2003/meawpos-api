// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package bootstrap

import (
	"fmt"
)

var defaultBusinessComponents = []string{"catalog.core", "catalog.pricing", "sales", "payment.cash"}
var componentRequirements = map[string][]string{
	"catalog.core":    {},
	"catalog.pricing": {"catalog.core"},
	"payment.cash":    {},
	"sales":           {"catalog.core", "catalog.pricing", "payment.cash"},
}

type componentSet map[string]bool

func newComponentSet(configured []string) (componentSet, error) {
	if configured == nil {
		configured = defaultBusinessComponents
	}
	set := componentSet{}
	for _, name := range configured {
		if _, known := componentRequirements[name]; !known {
			return nil, fmt.Errorf("unknown business component %q", name)
		}
		set[name] = true
	}
	for name := range set {
		for _, required := range componentRequirements[name] {
			if !set[required] {
				return nil, fmt.Errorf("business component %q requires %q", name, required)
			}
		}
	}
	return set, nil
}
