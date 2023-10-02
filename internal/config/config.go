package config

import (
	"context"
	"fmt"
	"github.com/RoyMusthang/backend/internal/api"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Config struct {
	DB *pgxpool.Pool
}

var a api.Api

func (c *Config) Initialize(host, user, password, dbname string) {
	dbConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, dbname))
	if err != nil {
		log.Panic(err)
	}

	dbConfig.MaxConns = 200
	dbConfig.MinConns = 100

	c.DB, err = pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Panic(err)
	}
}
