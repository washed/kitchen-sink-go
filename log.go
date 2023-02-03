package kitchen_sink

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DefaultLogLevel string = "info"

type LogConfig struct {
	LogLevel   string `yaml:"logLevel"`
	TimeFormat string `yaml:"timeFormat"`
	LogJSON    bool   `yaml:"logJSON"`
}

func InitLogger(logConfig LogConfig) error {
	timeFormat := logConfig.TimeFormat
	if timeFormat == "" {
		timeFormat = time.RFC3339Nano
	}

	zerolog.TimeFieldFormat = timeFormat

	// set log level to trace during init
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	if logConfig.LogJSON {
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat},
		)
	}

	if logConfig.LogLevel == "" {
		logConfig.LogLevel = DefaultLogLevel
		log.Error().
			Str("DefaultLogLevel", DefaultLogLevel).
			Msg("using default log level")
	}

	logLevel, err := zerolog.ParseLevel(logConfig.LogLevel)
	if err != nil {
		log.Error().
			Err(err).
			Str("logConfig.LogLevel", logConfig.LogLevel).
			Msg("error configuring log level")
		return err
	}
	zerolog.SetGlobalLevel(logLevel)
	log.Error().
		Str("logConfig.LogLevel", logConfig.LogLevel).
		Int("logLevel", int(logLevel)).
		Msg("set log level")

	return nil
}
