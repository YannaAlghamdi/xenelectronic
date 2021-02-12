package example

import (
	"github.com/gorilla/mux"

	"github.com/xenelectronic/core/internal/component/example"
)

// Controller is the REST controller
type Controller struct {
	exampleService example.Service
}

// NewController creates a new Controller
func NewController(exampleService example.Service) *Controller {
	return &Controller{
		exampleService,
	}
}

// Register registers the endpoints to the router
func (c *Controller) Register(router *mux.Router) {
	example := router.PathPrefix("/accounts").Subrouter()
	exampleService := c.exampleService

	// No auth
	example.
		Path("/no-auth").
		HandlerFunc(getNoAuthHandler(exampleService))
}
