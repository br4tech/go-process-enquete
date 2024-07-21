package main

import (
	"github.com/br4tech/go-process-enquete/config"
	"github.com/br4tech/go-process-enquete/internal/adapter"
	"github.com/br4tech/go-process-enquete/internal/model"
)

func main() {
	cfg := config.GetConfig()
	adapter := adapter.NewPostgresAdapter(&cfg)

	EnqueteMigration(adapter)
}

func EnqueteMigration(db *adapter.PostgresAdapter) {
	db.Db.Migrator().CreateTable(
		&model.Poll{},
		&model.Option{},
	)
}
