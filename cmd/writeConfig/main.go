package main

import (
	"time"

	ks "github.com/washed/kitchen-sink-go"
)

type Config struct {
	Log             ks.LogConfig `yaml:"log"`
	SomeConfigField string       `yaml:"someConfigField"`
}

func main() {
	config := Config{
		Log: ks.LogConfig{
			LogLevel:   "info",
			TimeFormat: time.RFC3339Nano,
			LogJSON:    false,
		},
		SomeConfigField: "foobar",
	}

	ks.WriteConfigFile("config.yaml", &config)
}
