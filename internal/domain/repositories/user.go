package repository

import (
	"context"

	entities "github.com/andy-takker/simple_server/internal/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userData *entities.CreateUserWithID) (*entities.User, error)
	FetchUserByID(ctx context.Context, userID string) (*entities.User, error)
	UpdateUserByID(ctx context.Context, userData *entities.UpdateUser) (*entities.User, error)
	DeleteUserByID(ctx context.Context, userID string) error
	FetchUserList(ctx context.Context, params *entities.UserListParams) (*[]entities.User, error)
	CountUsers(ctx context.Context, params *entities.UserListParams) (int64, error)
}
