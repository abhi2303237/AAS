// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/gofiber/fiber/v2/log"

	"github.com/abhi2303237/AAS/conceptDescription"
	"github.com/gofiber/fiber/v2"

	"github.com/abhi2303237/AAS/ioc"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

func main() {
	ioc.InitializeIocContainer()
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	// Create an instance of our handler which satisfies the generated interface
	conceptDescriptionRepo := conceptDescription.NewConceptDescriptionRepo()

	s := NewFiberConceptDescriptionServer(&conceptDescriptionRepo)

	// And we serve HTTP until the world ends.
	log.Fatal(s.Listen(net.JoinHostPort("0.0.0.0", *port)))
}

func NewFiberConceptDescriptionServer(conceptDescriptionRepo *conceptDescription.ServerInterface) *fiber.App {

	swagger, err := conceptDescription.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic fiber router
	app := fiber.New()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	app.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	conceptDescription.RegisterHandlers(app, *conceptDescriptionRepo)

	return app
}
