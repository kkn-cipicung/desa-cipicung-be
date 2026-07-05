package dashboard

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddDashboardPayload) error
	List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error)
	FindByID(ctx context.Context, payload DashboardPayload) (*DashboardResponse, error)
	Update(ctx context.Context, payload EditDashboardPayload) error
	Delete(ctx context.Context, payload DashboardPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddDashboardPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Description = strings.TrimSpace(payload.Description)

	if payload.CreatedBy == 0 {
		return fmt.Errorf("%w: created_by must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.CategoryID == 0 {
		return fmt.Errorf("%w: category_id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.Title == "" {
		return fmt.Errorf("%w: title is required", utils.ErrInvalidPayload)
	}
	if payload.Description == "" {
		return fmt.Errorf("%w: description is required", utils.ErrInvalidPayload)
	}

	media, err := utils.PrepareMedia(payload.ImgID, "uploads/dashboard", payload.CreatedBy)
	if err != nil {
		return err
	}

	return s.repository.Create(ctx, payload, media)
}

func (s *service) List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload DashboardPayload) (*DashboardResponse, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: dashboard id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditDashboardPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Description = strings.TrimSpace(payload.Description)

	if payload.ID == 0 {
		return fmt.Errorf("%w: dashboard id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.CategoryID == 0 {
		return fmt.Errorf("%w: category_id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.Title == "" {
		return fmt.Errorf("%w: title is required", utils.ErrInvalidPayload)
	}
	if payload.Description == "" {
		return fmt.Errorf("%w: description is required", utils.ErrInvalidPayload)
	}

	media, err := utils.PrepareMedia(payload.ImgID, "uploads/dashboard", 0)
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, payload, media)
}

func (s *service) Delete(ctx context.Context, payload DashboardPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: dashboard id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}
