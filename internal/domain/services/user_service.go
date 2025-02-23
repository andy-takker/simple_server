package services

import (
	"context"

	entities "github.com/andy-takker/simple_server/internal/domain/entities"
	repositories "github.com/andy-takker/simple_server/internal/domain/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewService(
	userRepository repositories.UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, userData *entities.CreateUser) (*entities.User, error) {
	var id = uuid.New().String()
	return s.userRepository.CreateUser(ctx, &entities.CreateUserWithID{
		ID:        id,
		Username:  userData.Username,
		Email:     userData.Email,
		Phone:     userData.Phone,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	})
}

func (s *UserService) FetchUserByID(ctx context.Context, userID string) (*entities.User, error) {
	return s.userRepository.FetchUserByID(ctx, userID)
}

func (s *UserService) UpdateUserByID(ctx context.Context, userData *entities.UpdateUser) (*entities.User, error) {
	return s.userRepository.UpdateUserByID(ctx, userData)
}

func (s *UserService) DeleteUserByID(ctx context.Context, userID string) error {
	return s.userRepository.DeleteUserByID(ctx, userID)
}

func (s *UserService) FetchUserList(ctx context.Context, params *entities.UserListParams) (*entities.UserList, error) {
	users, err := s.userRepository.FetchUserList(ctx, params)
	if err != nil {
		return nil, err
	}

	count, err := s.userRepository.CountUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	return &entities.UserList{
		Total: count,
		Items: *users,
	}, nil
}
