package db

import (
	"context"

	"github.com/DuckyMomo20012/go-todo/internal/common/libs/logger"
	pgxLogger "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

// Ref: https://github.com/jackc/pgx/issues/1582#issue-1683304180
func NewTracerLogger(logger *zerolog.Logger) pgx.QueryTracer {
	return &tracelog.TraceLog{
		Logger:   pgxLogger.NewLogger(*logger),
		LogLevel: tracelog.LogLevelTrace,
	}
}

func NewDb(connString string) *pgxpool.Pool {
	log := logger.Get()

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Panic().Err(err).Msg("failed to parse db connection string")
	}

	config.ConnConfig.Tracer = NewTracerLogger(&log)

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create db pool")
	}

	return pool
}
