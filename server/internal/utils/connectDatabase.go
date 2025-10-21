package utils

import (
	"context"
	"log"

	"github.com/hrpofficial736/uplift/server/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabase(ctx context.Context) *pgxpool.Pool {
	cfg := config.ConfigLoad()
	databaseUrl := cfg.DatabaseUrl

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Fatalf("error connecting to the database: %s", err)
	}
	return pool
}
