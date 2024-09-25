package migrations

import (
	"database/sql"
	"embed"

	goose "github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed sql/*.sql
var embedMigrations embed.FS

func Up(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)
	sqlFolder := "sql"

	if err := goose.SetDialect("postgres"); err != nil {
		log.Panic().Err(err).Msg("Failed to set dialect")
	}

	if err := goose.Up(db, sqlFolder); err != nil {
		log.Panic().Err(err).Msg("Failed to run goose up migration")
	}

	if err := goose.Status(db, sqlFolder); err != nil {
		log.Panic().Err(err).Msg("Failed to run goose status")
	}
}

func Down(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)
	sqlFolder := "sql"

	if err := goose.SetDialect("postgres"); err != nil {
		log.Panic().Err(err).Msg("Failed to set dialect")
	}

	if err := goose.Down(db, sqlFolder); err != nil {
		log.Panic().Err(err).Msg("Failed to run goose down migration")
	}
}
