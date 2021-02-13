package models

// import (
// 	"log"
// 	"os"
// 	"sync"
// 	"testing"

// 	"gitlab.com/talino-labs/crowd-connect/tcc-core.git/config"
// 	"gitlab.com/talino-labs/crowd-connect/tcc-core.git/db"
// )

// var Tester = struct {
// 	dbClient db.DBClient
// }{}

// var once sync.Once

// func TestMain(m *testing.M) {
// 	// Setup dependecies prior to running all tests in models package

// 	cfg, err := config.NewConfig("../../config/config.yaml")

// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}

// 	Tester.dbClient = db.NewPostgresClient(cfg)

// 	exitVal := m.Run()
// 	os.Exit(exitVal)
// }
