package schemas

import (
	"time"

	entities "github.com/andy-takker/simple_server/internal/domain/entities"
)

type CreateUserSchema struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type UserSchema struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ConvertToUserSchema(user *entities.User) *UserSchema {
	return &UserSchema{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

type UserListSchema struct {
	Items []UserSchema `json:"items"`
	Total int64        `json:"total"`
}

func ConvertToUserListSchema(users *entities.UserList) *UserListSchema {
	items := make([]UserSchema, len(users.Items))
	for i, user := range users.Items {
		items[i] = *ConvertToUserSchema(&user)
	}
	return &UserListSchema{
		Items: items,
		Total: users.Total,
	}
}
