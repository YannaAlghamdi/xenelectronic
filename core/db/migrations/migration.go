package migrations

import (
	"log"
	"sync"

	"github.com/YannaAlghamdi/xenelectronic/core/config"
	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"gopkg.in/gormigrate.v1"
)

type migration struct {
	config        *config.Config
	dbClient      db.DBClient
	gormigrations []*gormigrate.Migration
}

type Migration interface {
	Migrate(string) error
	Rollback(string) error
}

var singleton *migration
var once sync.Once

func NewMigration(config *config.Config, dbClient db.DBClient) Migration {
	once.Do(func() {
		singleton = &migration{
			config:   config,
			dbClient: dbClient,
			gormigrations: []*gormigrate.Migration{
				migrationBase,
				//migration202102031101.Migration,
				migrationSeed,
			},
		}
	})
	return singleton
}

func (migration *migration) Migrate(migrationID string) error {
	var err error
	db, close := migration.dbClient.Open()
	defer close()

	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, migration.gormigrations)

	log.Printf("Migration ID %s", migrationID)

	if migrationID != "" {
		err = m.MigrateTo(migrationID)
	} else {
		err = m.Migrate()
	}

	if err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration ran successfully")

	return err
}

func (migration *migration) Rollback(migrationID string) error {
	var err error
	db, close := migration.dbClient.Open()
	defer close()

	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, migration.gormigrations)

	err = m.RollbackTo(migrationID)

	if err != nil {
		log.Fatalf("Could not rollback to migration: %v", err)
	}

	log.Printf("Migration rollback ran successfully")

	return err
}
