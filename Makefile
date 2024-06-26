generate:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --package api --config=tools/config.yaml ConceptDescriptionRepository.yaml

generate registry:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --package AASRegistry --config=AASRegistry/config.yaml AASRegistry/AASRegistry.yaml

generate submode-registry:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --package SubmodelRegistry --config=SubmodelRegistry/config.yaml SubmodelRegistry/SubmodelRegistry.yaml