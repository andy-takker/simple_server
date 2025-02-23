package user

import (
	"context"
	"strings"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	entities "github.com/andy-takker/simple_server/internal/domain/entities"
	def "github.com/andy-takker/simple_server/internal/domain/repositories"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	pool *pgxpool.Pool
	m    sync.RWMutex
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) Close() {
	r.pool.Close()
}

func (r *repository) CreateUser(ctx context.Context, userData *entities.CreateUserWithID) (*entities.User, error) {
	r.m.Lock()
	defer r.m.Unlock()

	var user entities.User

	err := r.pool.QueryRow(
		ctx,
		`
			INSERT INTO users (id, username, email, phone, first_name, last_name)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, username, email, phone, first_name, last_name, created_at, updated_at
		`,
		userData.ID,
		userData.Username,
		userData.Email,
		userData.Phone,
		userData.FirstName,
		userData.LastName,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, entities.ErrorUserAlreadyExists
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) FetchUserByID(ctx context.Context, userID string) (*entities.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var user entities.User

	err := r.pool.QueryRow(
		ctx,
		`
			SELECT 
				id, 
				username,
				email,
				phone,
				first_name,
				last_name,
				created_at,
				updated_at
			FROM users
			WHERE id = $1 AND deleted_at IS NULL
		`,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, entities.ErrorUserNotFound
	}

	return &user, nil
}

func (r *repository) UpdateUserByID(ctx context.Context, userData *entities.UpdateUser) (*entities.User, error) {
	r.m.Lock()
	defer r.m.Unlock()

	var user entities.User
	err := r.pool.QueryRow(
		ctx,
		`
			UPDATE users
			SET 
				first_name = $1,
				last_name = $2,
				email = $3,
				phone = $4,
				username = $5,
				updated_at = now()
			WHERE id = $6
			RETURNING id, username, email, phone, first_name, last_name, created_at, updated_at
		`,
		userData.FirstName,
		userData.LastName,
		userData.Email,
		userData.Phone,
		userData.Username,
		userData.ID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) DeleteUserByID(ctx context.Context, userID string) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, err := r.pool.Exec(
		ctx,
		`
			UPDATE users
			SET deleted_at = now()
			WHERE id = $1
		`,
		userID,
	)

	return err
}

func (r *repository) FetchUserList(ctx context.Context, params *entities.UserListParams) (*[]entities.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	rows, err := r.pool.Query(
		ctx,
		`
			SELECT 
				id, 
				username,
				email,
				phone,
				first_name,
				last_name,
				created_at,
				updated_at
			FROM users
			WHERE deleted_at IS NULL
			ORDER BY created_at
			LIMIT $1
			OFFSET $2
		`,
		params.Limit,
		params.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &users, nil
}

func (r *repository) CountUsers(ctx context.Context, params *entities.UserListParams) (int64, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var count int64
	err := r.pool.QueryRow(
		ctx,
		`
			SELECT COUNT(*)
			FROM users
			WHERE deleted_at IS NULL
		`,
	).Scan(&count)
	return count, err
}
