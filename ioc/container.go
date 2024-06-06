package ioc

import (
	"github.com/abhi2303237/AAS/backend"
	"github.com/abhi2303237/AAS/conceptDescription"
	"github.com/golobby/container/v3"
)

func InitializeIocContainer() {
	container.Singleton(func() backend.ICrudRepo[conceptDescription.ConceptDescription] {
		return conceptDescription.NewInMemmoryConceptDescriptionCrudRepo()
	})
}
