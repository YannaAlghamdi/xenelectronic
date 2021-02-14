package main

import (
	"fmt"
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/config"
	handler "github.com/YannaAlghamdi/xenelectronic/core/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func init() {
	defer logger.Sync()
}

type Server interface {
	Start()
}

type server struct {
	config          *config.Config
	categoryHandler handler.CategoryHandler
	productHandler  handler.ProductHandler
	cartHandler     handler.CartHandler
}

func CreateServer(
	cfg *config.Config,
	categoryHandler handler.CategoryHandler,
	productHandler handler.ProductHandler,
	cartHandler handler.CartHandler,
) Server {
	return &server{
		config:          cfg,
		categoryHandler: categoryHandler,
		productHandler:  productHandler,
		cartHandler:     cartHandler,
	}
}

func (s *server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/categories/{id}/products", s.productHandler.ListProductsByCategoryId).Methods("GET")
	router.HandleFunc("/categories", s.categoryHandler.ListCategories).Methods("GET")
	router.HandleFunc("/carts", s.cartHandler.CreateCart).Methods("POST")
	router.HandleFunc("/carts", s.cartHandler.ListCarts).Methods("GET")
	router.HandleFunc("/carts", s.cartHandler.AddProductToCart).Methods("PUT")
	router.HandleFunc("/carts/{id}/products", s.cartHandler.ListProducts).Methods("GET")

	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	http.Handle("/", handlers.CORS(originsOK, headersOK, methodsOK)(router))
	logger.Info("Server started on " + s.config.HTTP.Port)
	fmt.Println(http.ListenAndServe(s.config.HTTP.Port, nil))
}
