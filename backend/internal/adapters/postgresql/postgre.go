package postgresql

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Util787/mws-content-registry/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Db  *pgxpool.Pool
	log *slog.Logger
}

func NewStorage(db *pgxpool.Pool, log *slog.Logger) *Storage {
	return &Storage{Db: db, log: log}
}

func ConnectPostgreSQL(cfg config.PostgesConfig) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgresql://%s:%s@postgres:5432/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Name)

	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	if err := pingDB(pool); err != nil {
		return nil, err
	}
	return pool, nil
}

func pingDB(pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Posgtresql not connected: ", err)
		return err
	}
	fmt.Println("Posgtresql connected")
	return nil
}
