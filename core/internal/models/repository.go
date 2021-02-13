package models

import (
	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"github.com/jinzhu/gorm"
)

type repository struct {
	dbClient db.DBClient
}

func (repository *repository) Connect() (db *gorm.DB, close func()) {
	return repository.dbClient.Open()
}

func newRepository(dbClient db.DBClient) *repository {
	return &repository{
		dbClient: dbClient}
}
