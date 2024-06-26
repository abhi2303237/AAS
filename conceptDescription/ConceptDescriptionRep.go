package conceptDescription

import (
	"net/http"

	"github.com/gofiber/fiber/v2/log"

	"github.com/abhi2303237/AAS/backend/types"
	"github.com/abhi2303237/AAS/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
)

type ConceptDescriptionRepo struct {
	CrudRep types.ICrudRepo[ConceptDescription]
}

func (repo *ConceptDescriptionRepo) GetAllConceptDescriptions(c *fiber.Ctx, params GetAllConceptDescriptionsParams) error {
	conceptDescriptions, err := repo.CrudRep.Get(c.Context())
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
	cd, err := repo.CrudRep.Put(c.Context(), &conceptDescription, *conceptDescription.Id)
	if err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusConflict, "Id already exist")
	}
	return c.Status(http.StatusOK).JSON(cd)
}

func (repo *ConceptDescriptionRepo) DeleteConceptDescriptionById(c *fiber.Ctx, cdIdentifier string) error {
	cd, err := repo.CrudRep.Delete(c.Context(), cdIdentifier)
	if err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusNotFound, "ConceptDescription not found")
	}
	return c.Status(http.StatusOK).JSON(cd)
}

func (repo *ConceptDescriptionRepo) GetConceptDescriptionById(c *fiber.Ctx, cdIdentifier string) error {

	cd, err := repo.CrudRep.GetById(c.Context(), cdIdentifier)
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
	cd, err := repo.CrudRep.Patch(c.Context(), &conceptDescription, cdIdentifier)
	if err != nil {
		return utils.SendConceptDescriptionError(c, http.StatusNotFound, "ConceptDescription not found")
	}
	return c.Status(http.StatusNoContent).JSON(cd)
}

func (repo *ConceptDescriptionRepo) GetDescription(c *fiber.Ctx) error {
	description := "This project is a comprehensive implementation of the DotAAS Part 2 HTTP/REST Submodel Service Specification, which is part of the Specification of the Asset Administration Shell. It provides a set of APIs for managing Concept Descriptions, including operations to get, post, delete, and update Concept Descriptions. The project also includes a self-describing endpoint that returns the service's capabilities and profiles. The implementation follows the guidelines and standards set by the Industrial Digital Twin Association (IDTA) to ensure interoperability and consistency across different systems."
	return c.Status(http.StatusOK).SendString(description)
}

func NewConceptDescriptionRepo() ServerInterface {
	var crudRepo types.ICrudRepo[ConceptDescription]
	if err := container.Resolve(&crudRepo); err != nil {
		log.Fatal(err)
	}

	return &ConceptDescriptionRepo{
		CrudRep: crudRepo,
	}
}
