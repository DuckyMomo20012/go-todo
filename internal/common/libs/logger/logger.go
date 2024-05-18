package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
)

// Ref: https://betterstack.com/community/guides/logging/zerolog/
var (
	once sync.Once
	log  zerolog.Logger
)

func Get() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		var logLevel int
		if viper.IsSet("LOG_LEVEL") {
			logLevel = viper.GetInt("LOG_LEVEL")
		} else {
			logLevel = int(zerolog.InfoLevel)
		}

		var output io.Writer = &zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if viper.IsSet("APP_ENV") && viper.GetString("APP_ENV") == "production" {
			output = os.Stdout
			logLevel = int(zerolog.InfoLevel)
		}

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Logger()
	})

	return log
}

func UpdateContext(update func(c zerolog.Context) zerolog.Context) {
	log.UpdateContext(update)
}

func WithService(serviceName string) zerolog.Logger {
	log = log.With().Str("service", serviceName).Logger()

	return log
}

func WithBasicSampler() zerolog.Logger {
	samplingRate := viper.GetUint32("LOG_SAMPLE_RATE")

	log = log.Sample(&zerolog.BasicSampler{
		N: samplingRate,
	})

	return log
}
