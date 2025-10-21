package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryDatabase(ctx context.Context, pool *pgxpool.Pool, query string, args ...interface{}) (pgx.Rows, error) {
	response, err := pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error while querying database: %s", err)
	}
	return response, nil

}
