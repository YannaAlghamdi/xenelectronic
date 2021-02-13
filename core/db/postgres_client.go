package db

import (
	"fmt"
	"log"
	"time"

	"github.com/YannaAlghamdi/xenelectronic/core/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresClient struct {
	*dbClient
}

func NewPostgresClient(config *config.Config) *PostgresClient {
	return &PostgresClient{newDbClient(config)}
}

func (client *PostgresClient) Open() (db *gorm.DB, close func()) {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		client.config.Postgres.Host,
		client.config.Postgres.Port,
		client.config.Postgres.User,
		client.config.Postgres.Db,
		client.config.Postgres.Password,
		client.config.Postgres.SSLMode))

	if err != nil {
		log.Panic(err)
	}

	client.db = db

	return db, func() {
		client.Close()
	}
}

func (client *PostgresClient) Close() {
	client.db.Close()
}
