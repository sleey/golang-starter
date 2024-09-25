package users_handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/rs/zerolog/log"
	"github.com/sleey/golang-starter/internal/datastore/db"
	"github.com/sleey/golang-starter/model"
)

type UserHandler struct {
	db *db.MainDB
}

func NewUserHandler(db *db.MainDB) *UserHandler {
	return &UserHandler{db: db}
}

type GetUsersListRequest struct {
}

type GetUsersListResponse struct {
	Body []model.User
}

// GetUserList returns a list of users
func (n UserHandler) GetUserList(ctx context.Context, input *GetUsersListRequest) (*GetUsersListResponse, error) {
	data, err := n.db.GetUsers(ctx)
	if err != nil {
		log.Err(err).Send()
		return nil, huma.Error500InternalServerError("Failed to get user", err)
	}

	return &GetUsersListResponse{Body: data}, nil
}

type GetUserRequest struct {
	ID int64 `path:"id" example:"1" doc:"ID of the user " format:"int64" required:"true"`
}

type GetUserResponse struct {
	Body model.User
}

func (n UserHandler) GetUser(ctx context.Context, input *GetUserRequest) (*GetUserResponse, error) {
	res, err := n.db.GetUser(ctx, input.ID)
	if err != nil {
		return nil, huma.Error500InternalServerError("Failed to get user", err)
	}

	return &GetUserResponse{Body: res}, nil
}
