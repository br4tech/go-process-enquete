package adapter

import (
	"fmt"

	"github.com/br4tech/go-process-enquete/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAdapter struct {
	Db *gorm.DB
}

func NewPostgresAdapter(cfg *config.Config) *PostgresAdapter {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s timezone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
		cfg.Db.SSLMode,
		cfg.Db.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	return &PostgresAdapter{Db: db}
}
