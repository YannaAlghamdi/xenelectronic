package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/xenlectronic/core/internal/entrypoint/api/rest/example"
)

type (
	// Config contains the entire configuration of the rest api
	Config struct {
		Host    string
		Port    int
		Spec    string
		Version string
		Cors    CORSConfig
	}

	// CORSConfig contains CORS related configuration
	CORSConfig struct {
		AllowedOrigins []string
		AllowedHeaders []string
		AllowedMethods []string
	}

	// API is the top level struct of the REST API implementation
	API struct {
		config            *Config
		router            *mux.Router
		exampleController *example.Controller
	}
)

// NewRestAPI will create and return and instance of the API struct
func NewRestAPI(config *Config, router *mux.Router, exampleController *example.Controller) *API {
	return &API{
		config:            config,
		router:            router,
		exampleController: exampleController,
	}
}

// Run will configure, start and serve the REST API and it the services it depends on.
func (api *API) Run() error {
	api.router = api.router.PathPrefix("/api/v1").Subrouter()
	api.registerHandlers()
	api.exposeSwagger()
	api.exposeVersion()
	api.exposeHealth()
	api.addMiddlewares()
	api.enableCORS()
	return http.ListenAndServe(api.address(), api.router)
}

func (api *API) address() string {
	return fmt.Sprintf("%s:%d", api.config.Host, api.config.Port)
}

func (api *API) exposeSwagger() {
	api.router.HandleFunc("/spec", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, api.config.Spec)
	})
	log.Infof("OpenAPI Spec accessible at http://%s/api/v1/spec", api.address())
}

func (api *API) exposeVersion() {
	api.router.HandleFunc("/version", func(res http.ResponseWriter, req *http.Request) {
		responses.WriteOKWithEntity(res, api.config.Version)
	})
}

func (api *API) exposeHealth() {
	api.router.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		responses.WriteOK(res)
	})
}

func (api *API) enableCORS() {
	cors := handlers.CORS(
		handlers.AllowedOrigins(api.config.Cors.AllowedOrigins),
		handlers.AllowedHeaders(api.config.Cors.AllowedHeaders),
		handlers.AllowedMethods(api.config.Cors.AllowedMethods),
	)
	api.router.Use(cors)
}

func (api *API) addMiddlewares() {
	logger := middlewares.Logger(log.StandardLogger())
	api.router.Use(logger)
	log.Info("Logger filter enabled")
	// TODO add JWT filter here
}

func (api *API) registerHandlers() {

	api.exampleController.Register(api.router)

}
