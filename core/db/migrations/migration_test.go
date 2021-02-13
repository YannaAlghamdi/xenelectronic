package migrations

import (
	"log"
	"testing"

	"github.com/YannaAlghamdi/xenelectronic/core/config"
	"github.com/YannaAlghamdi/xenelectronic/core/db"
)

func TestMigration(t *testing.T) {
	config, err := config.NewConfig("../../config/config.yaml")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	postgresClient := db.NewPostgresClient(config)
	migration := migrations.NewMigration(config, postgresClient)
	migration.Migrate()
}
