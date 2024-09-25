package util

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func InitializeDB(dbURL string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	return db
}
