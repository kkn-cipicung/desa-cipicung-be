package category

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/pkg/models"
	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddCategoryPayload) error
	List(ctx context.Context, payload ListCategoryPayload) ([]CategoryResponse, error)
	FindByID(ctx context.Context, payload CategoryPayload) (*CategoryResponse, error)
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
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Type = strings.TrimSpace(payload.Type)

	if payload.Name == "" {
		return fmt.Errorf("%w: name is required", utils.ErrInvalidPayload)
	}
	if payload.Type == "" {
		return fmt.Errorf("%w: type is required", utils.ErrInvalidPayload)
	}

	slug := utils.GenerateSlug(payload.Name)
	return s.repository.Create(ctx, payload, slug)
}

func (s *service) List(ctx context.Context, payload ListCategoryPayload) ([]CategoryResponse, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	items, err := s.repository.List(ctx, payload)
	if err != nil {
		return nil, err
	}
	return mapCategoryResponses(items), nil
}

func (s *service) FindByID(ctx context.Context, payload CategoryPayload) (*CategoryResponse, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: category id must be greater than 0", utils.ErrInvalidPayload)
	}
	item, err := s.repository.FindByID(ctx, payload)
	if err != nil {
		return nil, err
	}
	output := mapCategoryResponse(*item)
	return &output, nil
}

func (s *service) Update(ctx context.Context, payload EditCategoryPayload) error {
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Type = strings.TrimSpace(payload.Type)

	if payload.ID == 0 {
		return fmt.Errorf("%w: category id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.Name == "" {
		return fmt.Errorf("%w: name is required", utils.ErrInvalidPayload)
	}
	if payload.Type == "" {
		return fmt.Errorf("%w: type is required", utils.ErrInvalidPayload)
	}

	slug := utils.GenerateSlug(payload.Name)
	return s.repository.Update(ctx, payload, slug)
}

func (s *service) Delete(ctx context.Context, payload CategoryPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: category id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func mapCategoryResponses(items []models.Category) []CategoryResponse {
	outputs := make([]CategoryResponse, 0, len(items))
	for _, item := range items {
		outputs = append(outputs, mapCategoryResponse(item))
	}
	return outputs
}

func mapCategoryResponse(item models.Category) CategoryResponse {
	return CategoryResponse{
		ID:        item.ID,
		Name:      item.Name,
		Slug:      item.Slug,
		Type:      item.Type,
		CreatedAt: utils.FormatTimestamp(item.CreatedAt),
	}
}
