//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
)


//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=tools/config.yaml ConceptDescriptionRepository.yaml