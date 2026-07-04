package auth

import (
	"context"
	"database/sql"
	"errors"

	"cipicung.id/be/pkg/models"
	"github.com/jmoiron/sqlx"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	Register(ctx context.Context, payload RegisterPayload) error
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
	InsertSessionLog(ctx context.Context, userID uint, accessToken, refreshToken string) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Register(ctx context.Context, payload RegisterPayload) error {
	role := "admin"

	query := `
		INSERT INTO users (name, username, password, role)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, payload.Name, payload.Username, payload.Password, role)
	return err
}

func (r *repository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT id, COALESCE(role, '') AS role, username, password, COALESCE(is_active, TRUE) AS is_active
		FROM users
		WHERE username = $1
	`

	var user models.User
	if err := r.db.GetContext(ctx, &user, query, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) UpdateLastLogin(ctx context.Context, userID uint) error {
	query := `
		UPDATE users
		SET last_login = NOW()
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

func (r *repository) InsertSessionLog(ctx context.Context, userID uint, accessToken, refreshToken string) error {
	query := `
		INSERT INTO user_session_log (user_id, access_token, refresh_token)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(ctx, query, userID, accessToken, refreshToken)
	return err
}

