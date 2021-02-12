package cmd

import (
	"github.com/spf13/viper"
)

var defaults = map[string]interface{}{
	"log.level": "debug",

	"api.rest.host":                "0.0.0.0",
	"api.rest.port":                8080,
	"api.rest.spec":                "./openapi.yaml",
	"api.rest.cors.allowedOrigins": []string{"*"},
	"api.rest.cors.allowedHeaders": []string{
		"Content-Type",
		"Sec-Fetch-Dest",
		"Referer",
		"accept",
		"Sec-Fetch-Mode",
		"Sec-Fetch-Site",
		"User-Agent",
		"API-KEY",
		"Authorization",
	},
	"api.rest.cors.allowedMethods": []string{
		"OPTIONS",
		"GET",
		"POST",
		"PUT",
		"DELETE",
	},

	// Secrets file for setting credentials in environments
	"secrets.file": "config/secrets.json",
        
	// Put custom configuration here
}

func init() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}
