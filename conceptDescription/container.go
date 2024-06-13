package conceptDescription

import (
	"github.com/abhi2303237/AAS/backend/types"
	"github.com/golobby/container/v3"
)

func InitializeIocContainer() {
	container.Singleton(func() types.ICrudRepo[ConceptDescription] { return NewConceptDescriptionCrudRepo() })
}
