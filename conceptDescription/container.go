package conceptDescription

import (
	"github.com/abhi2303237/AAS/app"
	"github.com/abhi2303237/AAS/backend/types"
	"github.com/abhi2303237/AAS/conceptDescription/config"
	"github.com/golobby/container/v3"
)

func InitializeIocContainer() {
	container.Singleton(func() config.Config { return app.GetConfig[config.Config]() })
	container.Singleton(func() types.ICrudRepo[ConceptDescription] { return NewConceptDescriptionCrudRepo() })
}
