package business

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/pkg/models"
	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddBusinessPayload) error
	List(ctx context.Context, payload ListBusinessPayload) ([]models.Business, error)
	FindByID(ctx context.Context, payload BusinessPayload) (*models.Business, error)
	Update(ctx context.Context, payload EditBusinessPayload) error
	Delete(ctx context.Context, payload BusinessPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddBusinessPayload) error {
	payload.OwnerName = strings.TrimSpace(payload.OwnerName)
	payload.BusinessName = strings.TrimSpace(payload.BusinessName)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.Phone = strings.TrimSpace(payload.Phone)
	payload.Address = strings.TrimSpace(payload.Address)

	if payload.CategoryID == nil || *payload.CategoryID == 0 {
		return fmt.Errorf("%w: category_id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.OwnerName == "" {
		return fmt.Errorf("%w: owner_name is required", utils.ErrInvalidPayload)
	}
	if payload.BusinessName == "" {
		return fmt.Errorf("%w: business_name is required", utils.ErrInvalidPayload)
	}
	if payload.Description == "" {
		return fmt.Errorf("%w: description is required", utils.ErrInvalidPayload)
	}
	if payload.Phone == "" {
		return fmt.Errorf("%w: phone is required", utils.ErrInvalidPayload)
	}
	if payload.Address == "" {
		return fmt.Errorf("%w: address is required", utils.ErrInvalidPayload)
	}

	if payload.Instagram != nil {
		trimmed := strings.TrimSpace(*payload.Instagram)
		payload.Instagram = &trimmed
	}
	if payload.Facebook != nil {
		trimmed := strings.TrimSpace(*payload.Facebook)
		payload.Facebook = &trimmed
	}

	return s.repository.Create(ctx, payload)
}

func (s *service) List(ctx context.Context, payload ListBusinessPayload) ([]models.Business, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	payload.Type = strings.TrimSpace(payload.Type)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload BusinessPayload) (*models.Business, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: business id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditBusinessPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: business id must be greater than 0", utils.ErrInvalidPayload)
	}

	if payload.CategoryID == nil || *payload.CategoryID == 0 {
		return fmt.Errorf("%w: category_id must be greater than 0", utils.ErrInvalidPayload)
	}

	if payload.OwnerName != nil {
		trimmed := strings.TrimSpace(*payload.OwnerName)
		if trimmed == "" {
			return fmt.Errorf("%w: owner_name cannot be empty", utils.ErrInvalidPayload)
		}
		payload.OwnerName = &trimmed
	}

	if payload.BusinessName != nil {
		trimmed := strings.TrimSpace(*payload.BusinessName)
		if trimmed == "" {
			return fmt.Errorf("%w: business_name cannot be empty", utils.ErrInvalidPayload)
		}
		payload.BusinessName = &trimmed
	}

	if payload.Description != nil {
		trimmed := strings.TrimSpace(*payload.Description)
		if trimmed == "" {
			return fmt.Errorf("%w: description cannot be empty", utils.ErrInvalidPayload)
		}
		payload.Description = &trimmed
	}

	if payload.Phone != nil {
		trimmed := strings.TrimSpace(*payload.Phone)
		if trimmed == "" {
			return fmt.Errorf("%w: phone cannot be empty", utils.ErrInvalidPayload)
		}
		payload.Phone = &trimmed
	}

	if payload.Address != nil {
		trimmed := strings.TrimSpace(*payload.Address)
		if trimmed == "" {
			return fmt.Errorf("%w: address cannot be empty", utils.ErrInvalidPayload)
		}
		payload.Address = &trimmed
	}

	if payload.Instagram != nil {
		trimmed := strings.TrimSpace(*payload.Instagram)
		payload.Instagram = &trimmed
	}
	if payload.Facebook != nil {
		trimmed := strings.TrimSpace(*payload.Facebook)
		payload.Facebook = &trimmed
	}

	return s.repository.Update(ctx, payload)
}

func (s *service) Delete(ctx context.Context, payload BusinessPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: business id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}
