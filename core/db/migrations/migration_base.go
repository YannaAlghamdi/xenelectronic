package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var migrationBase = &gormigrate.Migration{
	ID:       "base",
	Migrate:  func(tx *gorm.DB) error { return nil },
	Rollback: func(tx *gorm.DB) error { return nil },
}
