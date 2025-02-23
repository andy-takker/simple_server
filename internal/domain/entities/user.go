package model

import (
	"time"
)

type CreateUser struct {
	Username  string
	Email     string
	Phone     string
	FirstName string
	LastName  string
}

type CreateUserWithID struct {
	ID        string
	Username  string
	Email     string
	Phone     string
	LastName  string
	FirstName string
}

type User struct {
	ID        string
	Username  string
	Email     string
	Phone     string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserListParams struct {
	Limit  int64
	Offset int64
}

type UserList struct {
	Total int64
	Items []User
}

type UpdateUser struct {
	ID        string
	Username  string
	Email     string
	Phone     string
	FirstName string
	LastName  string
}
