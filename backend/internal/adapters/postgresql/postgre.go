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

const (
	defaultMaxConns        = 10
	defaultConnMaxLifetime = time.Hour
	defaultConnMaxIdleTime = time.Minute * 10
)

func NewStorage(db *pgxpool.Pool, log *slog.Logger) *Storage {
	return &Storage{Db: db, log: log}
}

func ConnectPostgreSQL(cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = defaultMaxConns
	poolConfig.MaxConnLifetime = defaultConnMaxLifetime
	poolConfig.MaxConnIdleTime = defaultConnMaxIdleTime

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
