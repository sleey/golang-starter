package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/sleey/golang-starter/model"
)

type MainDB struct {
	DB *sqlx.DB
}

func NewMainDB(db *sqlx.DB) *MainDB {
	return &MainDB{DB: db}
}

func (m *MainDB) GetUsers(ctx context.Context) (result []model.User, err error) {
	err = m.DB.SelectContext(ctx, &result, `
	SELECT 1 as user_id, 'test_username1' as username
	UNION ALL
	SELECT 2 as user_id, 'test_username2' as username
	UNION ALL
	SELECT 3 as user_id, 'test_username3' as username;
	`)

	return
}

func (m *MainDB) GetUser(ctx context.Context, id int64) (result model.User, err error) {
	err = m.DB.QueryRowxContext(ctx, `
	SELECT user_id, username
	FROM (
		SELECT 1 as user_id, 'test_username1' as username
		UNION ALL
		SELECT 2 as user_id, 'test_username2' as username
		UNION ALL
		SELECT 3 as user_id, 'test_username3' as username
	) AS mock_data
	WHERE user_id = $1
	`, id).StructScan(&result)

	log.Info().Msgf("GetUser: %v", result)
	return
}
