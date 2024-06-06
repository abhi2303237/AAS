package conceptDescription

import (
	"net/http"

	"github.com/gofiber/fiber/v2/log"

	"github.com/abhi2303237/AAS/backend"
	"github.com/abhi2303237/AAS/backend/memmory"
	"github.com/abhi2303237/AAS/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
)

type ConceptDescriptionRepo struct {
	CrudRep backend.ICrudRepo[ConceptDescription]
}

func (repo *ConceptDescriptionRepo) GetAllConceptDescriptions(c *fiber.Ctx, params GetAllConceptDescriptionsParams) error {
	conceptDescriptions, err := repo.CrudRep.Get()
	if err != nil {
		utils.SendConceptDescriptionError(c, http.StatusConflict, "Id already exist")
	}
	return c.Status(http.StatusOK).JSON(conceptDescriptions)
}

func (repo *ConceptDescriptionRepo) PostConceptDescription(c *fiber.Ctx) error {
	var conceptDescription PostConceptDescriptionJSONRequestBody
	if err := c.BodyParser(&conceptDescription); err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusBadRequest, "Invalid format for New ConceptDescription")
	}
	cd, err := repo.CrudRep.Put(&conceptDescription, *conceptDescription.Id)
	if err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusConflict, "Id already exist")
	}
	return c.Status(http.StatusOK).JSON(cd)
}

func (repo *ConceptDescriptionRepo) DeleteConceptDescriptionById(c *fiber.Ctx, cdIdentifier string) error {
	cd, err := repo.CrudRep.Delete(cdIdentifier)
	if err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusNotFound, "ConceptDescription not found")
	}
	return c.Status(http.StatusOK).JSON(cd)
}

func (repo *ConceptDescriptionRepo) GetConceptDescriptionById(c *fiber.Ctx, cdIdentifier string) error {

	cd, err := repo.CrudRep.GetById(cdIdentifier)
	if err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusNotFound, "ConceptDescription not found")
	}
	return c.Status(http.StatusOK).JSON(cd)
}

func (repo *ConceptDescriptionRepo) PutConceptDescriptionById(c *fiber.Ctx, cdIdentifier string) error {
	var conceptDescription PutConceptDescriptionByIdJSONRequestBody
	if err := c.BodyParser(&conceptDescription); err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusBadRequest, "Invalid format for New ConceptDescription")
	}
	cd, err := repo.CrudRep.Patch(&conceptDescription, cdIdentifier)
	if err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusNotFound, "ConceptDescription not found")
	}
	return c.Status(http.StatusNoContent).JSON(cd)
}

func (repo *ConceptDescriptionRepo) GetDescription(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func NewConceptDescriptionRepo() ServerInterface {
	var inMemmoryConceptDescriptionCrudRepo backend.ICrudRepo[ConceptDescription]
	if err := container.Resolve(&inMemmoryConceptDescriptionCrudRepo); err != nil {
		log.Fatal(err)
	}

	return &ConceptDescriptionRepo{
		CrudRep: inMemmoryConceptDescriptionCrudRepo,
	}
}

func NewInMemmoryConceptDescriptionCrudRepo() backend.ICrudRepo[ConceptDescription] {
	return &memmory.InMemmoryCrudRepo[ConceptDescription]{
		Data: make(map[string]ConceptDescription),
	}
}
