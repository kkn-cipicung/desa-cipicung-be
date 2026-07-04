package category

import (
	"context"
	"strings"

	"cipicung.id/be/pkg/models"
)

const (
	defaultCategoryLimit = 10
	maxCategoryLimit     = 100
)

type Service interface {
	Create(ctx context.Context, payload AddCategoryPayload) error
	List(ctx context.Context, payload ListCategoryPayload) ([]models.Category, error)
	FindByID(ctx context.Context, payload CategoryPayload) (*models.Category, error)
	Update(ctx context.Context, payload EditCategoryPayload) error
	Delete(ctx context.Context, payload CategoryPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddCategoryPayload) error {
	slug := generateSlug(payload.Name)
	return s.repository.Create(ctx, payload, slug)
}

func (s *service) List(ctx context.Context, payload ListCategoryPayload) ([]models.Category, error) {
	payload = normalizeListCategoryPayload(payload)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload CategoryPayload) (*models.Category, error) {
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditCategoryPayload) error {
	slug := generateSlug(payload.Name)
	return s.repository.Update(ctx, payload, slug)
}

func (s *service) Delete(ctx context.Context, payload CategoryPayload) error {
	return s.repository.Delete(ctx, payload)
}

func normalizeListCategoryPayload(payload ListCategoryPayload) ListCategoryPayload {
	if payload.Limit <= 0 {
		payload.Limit = defaultCategoryLimit
	}
	if payload.Limit > maxCategoryLimit {
		payload.Limit = maxCategoryLimit
	}
	if payload.Index < 0 {
		payload.Index = 0
	}
	return payload
}

func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
