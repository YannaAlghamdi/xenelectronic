package models

import (
	"log"
	"os"
	"sync"
	"testing"

	"github.com/YannaAlghamdi/xenelectronic/core/config"
	"github.com/YannaAlghamdi/xenelectronic/core/db"
)

var Tester = struct {
	dbClient db.DBClient
}{}

var once sync.Once

func TestMain(m *testing.M) {
	// Setup dependecies prior to running all tests in models package

	cfg, err := config.NewConfig("../../config/config.yaml")

	if err != nil {
		log.Fatalf("%v", err)
	}

	Tester.dbClient = db.NewPostgresClient(cfg)

	exitVal := m.Run()
	os.Exit(exitVal)
}
