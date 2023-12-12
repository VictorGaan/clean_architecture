package postgres

import (
	"clean_architecture/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.SslMode)
	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate()
	return db, nil
}
