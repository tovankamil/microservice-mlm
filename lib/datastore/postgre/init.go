package postgre

import (
	"e-wallet-microservices/lib/config"
	"fmt"

	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	cfg      *config.Config
	DBClient DBInterface
}

func InitDbConnection(cfg *config.Config) *Database {
	return &Database{
		cfg: cfg,
	}
}

func (db *Database) InitiateConnection() (DBInterface, error) {
	cfg := db.cfg

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgre.Address,
		cfg.Postgre.Ports,
		cfg.Postgre.Username,
		cfg.Postgre.Password,
		cfg.Postgre.Database,
		cfg.Postgre.SSLMode)

	return db.Connect(connectionStr)
}

func (db *Database) Connect(connStr string) (DBInterface, error) {
	cfg := db.cfg.Postgre
	client, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}
	if cfg.MaxOpenConnections != -1 {
		client.SetMaxOpenConns(cfg.MaxOpenConnections)
	}
	if cfg.ConnMaxLifetime != -1 {
		client.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime))
	}
	if cfg.ConnMaxIdleTime != -1 {
		client.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime))
	}
	return client, nil
}
