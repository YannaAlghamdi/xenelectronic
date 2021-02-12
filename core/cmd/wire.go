//+build wireinject

package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	// REST endpoints
	"github.com/xenelectronic/core/internal/entrypoint/api/rest"

	//service components
	"github.com/xenelectronic/core/internal/component/example"

	//controllers
	exampleHandlers "github.com/xenelectronic/core/internal/entrypoint/api/rest/example"
)

type Keys struct {
}

func createRestAPI() *rest.API {
	wire.Build(
		// services
		example.NewDefaultExampleService,

		// handlers
		exampleHandlers.NewController,

		// rest
		mux.NewRouter,
		ProvideRestAPIConfig,
		rest.NewRestAPI,

		// Allow configration of credentials from file
		ProvideKeysFromFile,
		// Put other factory method references here
	)
	return &rest.API{}
}

func ProvideRestAPIConfig(keys *Keys) *rest.Config {
	var config rest.Config
	err := viper.UnmarshalKey("api.rest", &config)
	if err != nil {
		log.WithError(err).Error("unable to read RestAPIConfig")
		os.Exit(1)
	}

	config.Version = root.Version

	log.Info("========================================")
	log.Info("API Configuration")
	log.Info("========================================")
	log.Info("Host:    ", config.Host)
	log.Info("Port:    ", config.Port)
	log.Info("Spec:    ", config.Spec)
	log.Info("Version: ", config.Port)

	log.Debugf("API Config: %+v", config)
	return &config
}

func ProvideKeysFromFile() *Keys {
	file := viper.GetString("secrets.file")
	jsonFile, err := os.Open(file)
	if err != nil {
		log.WithError(err).Error("unable to read security keys file")
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var secrets Keys
	json.Unmarshal(byteValue, &secrets)

	return &secrets
}

// Put other factory function declarations here
