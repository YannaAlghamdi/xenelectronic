//+build wireinject

package main

import (
	"github.com/YannaAlghamdi/xenelectronic/core/config"
	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"github.com/YannaAlghamdi/xenelectronic/core/db/migrations"
	handler "github.com/YannaAlghamdi/xenelectronic/core/handlers"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/services"
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	CreateServer,
	models.NewProductRespository,
	models.NewCategoryRepository,
	models.NewCartRepository,
	models.NewOrderRepository,
	services.NewCartService,
	services.NewProductService,
	services.NewCategoryService,
	services.NewOrderService,
	handler.NewOrderHandler,
	handler.NewCategoryHandler,
	handler.NewProductHandler,
	handler.NewCartHandler,
	// Change these next two lines to use a different DB client
	db.NewPostgresClient,
	wire.Bind(new(db.DBClient), new(*db.PostgresClient)),
)

var MigrationSet = wire.NewSet(
	migrations.NewMigration,
	// Change these next two lines to use a different DB client
	db.NewPostgresClient,
	wire.Bind(new(db.DBClient), new(*db.PostgresClient)),
)

func buildServer(cfg *config.Config) (Server, error) {
	wire.Build(DependencySet)
	return nil, nil
}

func initializeMigration(cfg *config.Config) (migrations.Migration, error) {
	panic(wire.Build(MigrationSet))
}
