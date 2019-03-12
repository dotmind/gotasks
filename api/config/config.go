package config

import "github.com/tkanos/gonfig"

var configFile string = "./config/config.json"

type Configuration struct {
	Port     string `json:"port"`
	WifiName string `json:"wifi-name"`
}

func Load() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(configFile, &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
