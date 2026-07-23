package dashboard

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddDashboardPayload) error
	List(ctx context.Context, payload ListDashboardPayload) ([]DashboardOutput, error)
	Detail(ctx context.Context) (*DashboardOutput, error)
	FindActive(ctx context.Context) (*DashboardOutput, error)
	FindOverview(ctx context.Context) (*DashboardOverviewOutput, error)
	Update(ctx context.Context, payload EditDashboardPayload) error
	Activate(ctx context.Context, payload DashboardPayload) error
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

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/dashboard", payload.CreatedBy)
	if err != nil {
		return err
	}

	if err := s.repository.Create(ctx, payload, media); err != nil {
		_ = utils.RemovePreparedMedia(media)
		return err
	}
	return nil
}

func (s *service) List(ctx context.Context, payload ListDashboardPayload) ([]DashboardOutput, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	items, err := s.repository.List(ctx, payload)
	if err != nil {
		return nil, err
	}

	outputs := make([]DashboardOutput, 0, len(items))
	for _, item := range items {
		outputs = append(outputs, mapDashboardOutput(item))
	}
	return outputs, nil
}

func (s *service) Detail(ctx context.Context) (*DashboardOutput, error) {
	item, err := s.repository.Detail(ctx)
	if err != nil {
		return nil, err
	}
	output := mapDashboardOutput(*item)
	return &output, nil
}

func (s *service) FindActive(ctx context.Context) (*DashboardOutput, error) {
	item, err := s.repository.FindActive(ctx)
	if err != nil {
		return nil, err
	}
	output := mapDashboardOutput(*item)
	return &output, nil
}

func (s *service) FindOverview(ctx context.Context) (*DashboardOverviewOutput, error) {
	return s.repository.FindOverview(ctx)
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
	if payload.UpdatedBy == 0 {
		return fmt.Errorf("%w: updated_by must be greater than 0", utils.ErrInvalidPayload)
	}

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/dashboard", payload.UpdatedBy)
	if err != nil {
		return err
	}

	if err := s.repository.Update(ctx, payload, media); err != nil {
		_ = utils.RemovePreparedMedia(media)
		return err
	}

	return nil
}

func (s *service) Activate(ctx context.Context, payload DashboardPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: dashboard id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Activate(ctx, payload)
}

func (s *service) Delete(ctx context.Context, payload DashboardPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: dashboard id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func mapDashboardOutput(item DashboardResponse) DashboardOutput {
	return DashboardOutput{
		ID: item.ID,
		Creator: DashboardRef{
			ID:   item.CreatedBy,
			Name: item.CreatorName,
		},
		Category: DashboardRef{
			ID:   item.CategoryID,
			Name: item.CategoryName,
		},
		Title:       item.Title,
		Description: item.Description,
		Media:       item.Media,
		IsActive:    item.IsActive,
		CreatedAt:   utils.FormatTimestamp(item.CreatedAt),
	}
}
