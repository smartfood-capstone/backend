package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/smartfood-capstone/backend/internal/config"
)

func New(cfg config.Config) *sqlx.DB {

	log.Println(cfg)
	connURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.PostgresDBUser, cfg.PostgresDBPassword, cfg.PostgresHost, cfg.PostgresDBPort, cfg.PostgresDBName, cfg.PostgresSSLMode)
	db, err := sqlx.Connect("postgres", connURL)
	if err != nil {
		log.Fatalf("error when creating database instance err: %s", err)
	}
	return db
}
