package config

import (
	"github.com/tkanos/gonfig"
)

//Configuration struct ...
type Configuration struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

// GetConfig ...
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.json", &configuration)
	return configuration
}
