// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package main

import (
	"flag"
	"net"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golobby/container/v3"

	"github.com/abhi2303237/AAS/app"
	"github.com/abhi2303237/AAS/conceptDescription"
	"github.com/abhi2303237/AAS/conceptDescription/config"
	"github.com/abhi2303237/AAS/utils"
	"github.com/gofiber/fiber/v2"

	middleware "github.com/oapi-codegen/fiber-middleware"
)

func main() {
	conceptDescription.InitializeIocContainer()
	var config config.Config
	utils.HandleFatal(container.Resolve(&config))
	port := flag.String("port", config.Port, "Port for test HTTP server")
	flag.Parse()

	// Create an instance of our handler which satisfies the generated interface
	conceptDescriptionRepo := conceptDescription.NewConceptDescriptionRepo()

	s := NewFiberConceptDescriptionServer(&conceptDescriptionRepo, config)

	// And we serve HTTP until the world ends.
	log.Fatal(s.Listen(net.JoinHostPort("0.0.0.0", *port)))
}

func NewFiberConceptDescriptionServer(conceptDescriptionRepo *conceptDescription.ServerInterface, config config.Config) *fiber.App {

	swagger, err := conceptDescription.GetSwagger()
	utils.HandleFatal(err)

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic fiber router
	app := app.NewFiberApp(fiber.Config{}, app.Config{
		EnableMetricsUi: config.EnableMetricsUi,
		AppTitle:        config.ApplicationName,
		ContextPath:     config.ContextPath,
		LogLevel:        config.LogLevel,
	})

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	app.Use(middleware.OapiRequestValidator(swagger))
	// We now register our petStore above as the handler for the interface
	conceptDescription.RegisterHandlers(app, *conceptDescriptionRepo)

	return app
}
