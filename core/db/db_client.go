package db

import (
	"github.com/YannaAlghamdi/xenelectronic/core/config"
	"github.com/jinzhu/gorm"
)

type dbClient struct {
	config *config.Config
	db     *gorm.DB
	// DBClient
}

type DBClient interface {
	Open() (db *gorm.DB, close func())
	Close()
}

func newDbClient(config *config.Config) *dbClient {
	return &dbClient{
		config: config,
	}
}
