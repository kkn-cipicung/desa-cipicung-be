package auth

import (
	"context"
	"errors"

	"cipicung.id/be/utils"
	tokenutils "cipicung.id/be/utils/token"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidAuthPayload = utils.ErrInvalidAuthPayload
	ErrInvalidCredentials = utils.ErrInvalidCredentials
	ErrInactiveUser       = utils.ErrInactiveUser
	ErrUsernameTaken      = utils.ErrUsernameTaken
)

const defaultRegisterRole = "admin"

type Service interface {
	Register(ctx context.Context, payload RegisterPayload) error
	Login(ctx context.Context, payload LoginPayload) (*tokenutils.Pair, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Register(ctx context.Context, payload RegisterPayload) error {
	payload.Name = utils.NormalizeAuthName(payload.Name)
	payload.Username = utils.NormalizeAuthUsername(payload.Username)

	if err := utils.ValidateAuthRegisterPayload(payload.Name, payload.Username, payload.Password); err != nil {
		return err
	}

	if err := s.ensureUsernameAvailable(ctx, payload.Username); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	payload.Password = string(hashedPassword)

	return s.repository.Register(ctx, payload, defaultRegisterRole)
}

func (s *service) Login(ctx context.Context, payload LoginPayload) (*tokenutils.Pair, error) {
	payload.Username = utils.NormalizeAuthUsername(payload.Username)

	if err := utils.ValidateAuthLoginPayload(payload.Username, payload.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	user, err := s.repository.FindByUsername(ctx, payload.Username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if !user.IsActive {
		return nil, ErrInactiveUser
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := s.repository.UpdateLastLogin(ctx, user.ID); err != nil {
		return nil, err
	}

	tokens, err := tokenutils.GeneratePair(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	if err := s.repository.InsertSessionLog(ctx, user.ID, tokens.AccessToken, tokens.RefreshToken); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *service) ensureUsernameAvailable(ctx context.Context, username string) error {
	_, err := s.repository.FindByUsername(ctx, username)
	if err == nil {
		return ErrUsernameTaken
	}

	if errors.Is(err, ErrUserNotFound) {
		return nil
	}

	return err
}
