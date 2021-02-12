package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.AutomaticEnv()
	// To get the value from the config file using key
	// viper package read .env
	user := viper.Get("datasource.username")
	password := viper.Get("datasource.password")
	database := viper.Get("datasource.database")
	host := viper.Get("datasource.host")
	port := viper.Get("datasource.port")
	// https://gobyexample.com/string-formatting
	conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, database, password)
	fmt.Println("conname is\t\t", conname)
	db, err := gorm.Open("postgres", conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Product{})
	// Initialise value
	u1 := uuid.Must(uuid.NewV4(), nil)
	m := Product{ID: u1, Name: "ipad", Description: "test"}
	db.Create(&m)
	return db
}
