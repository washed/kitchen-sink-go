package main

import (
	"github.com/rs/zerolog/log"

	ks "github.com/washed/kitchen-sink-go"
)

type Config struct {
	Log             ks.LogConfig `yaml:"log"`
	SomeConfigField string       `yaml:"someConfigField"`
}

func main() {
	config := Config{}
	log.Info().Interface("config", config).Msg("empty config")
	ks.ReadConfig(&config)
	log.Info().Interface("config", config).Msg("read config")
	ks.InitLogger(config.Log)

	log.Info().Interface("config", config).Msg("human readable log")

	config.Log.LogJSON = true
	ks.InitLogger(config.Log)

	log.Info().Interface("config", config).Msg("JSON log")

}
