package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	tUser = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type PostgresClient struct {
	Connection *sqlx.DB
}

func NewPostgresDB(cfg Config) (*PostgresClient, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}

	return &PostgresClient{Connection: db}, nil
}
