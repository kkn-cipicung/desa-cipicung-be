package potential

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddPotentialPayload) error
	List(ctx context.Context, payload ListPotentialPayload) ([]PotentialResponse, error)
	FindByID(ctx context.Context, payload PotentialPayload) (*PotentialResponse, error)
	Update(ctx context.Context, payload EditPotentialPayload) error
	Delete(ctx context.Context, payload PotentialPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddPotentialPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Subtitle = strings.TrimSpace(payload.Subtitle)
	payload.Slug = strings.TrimSpace(payload.Slug)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.OwnerName = strings.TrimSpace(payload.OwnerName)
	payload.OwnerMsisdn = strings.TrimSpace(payload.OwnerMsisdn)

	if payload.UploadedBy == 0 {
		return fmt.Errorf("%w: uploaded_by must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.CategoryID == 0 {
		return fmt.Errorf("%w: category_id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.Title == "" {
		return fmt.Errorf("%w: title is required", utils.ErrInvalidPayload)
	}
	if payload.Slug == "" {
		return fmt.Errorf("%w: slug is required", utils.ErrInvalidPayload)
	}
	if payload.Description == "" {
		return fmt.Errorf("%w: description is required", utils.ErrInvalidPayload)
	}
	if payload.Latitude < -90 || payload.Latitude > 90 {
		return fmt.Errorf("%w: latitude must be between -90 and 90", utils.ErrInvalidPayload)
	}
	if payload.Longitude < -180 || payload.Longitude > 180 {
		return fmt.Errorf("%w: longitude must be between -180 and 180", utils.ErrInvalidPayload)
	}
	if payload.OwnerName == "" {
		return fmt.Errorf("%w: owner_name is required", utils.ErrInvalidPayload)
	}

	media, err := utils.PrepareMedia(payload.ImgID, "uploads/potentials", payload.UploadedBy)
	if err != nil {
		return err
	}

	return s.repository.Create(ctx, payload, media)
}

func (s *service) List(ctx context.Context, payload ListPotentialPayload) ([]PotentialResponse, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload PotentialPayload) (*PotentialResponse, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: potential id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditPotentialPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Subtitle = strings.TrimSpace(payload.Subtitle)
	payload.Slug = strings.TrimSpace(payload.Slug)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.OwnerName = strings.TrimSpace(payload.OwnerName)
	payload.OwnerMsisdn = strings.TrimSpace(payload.OwnerMsisdn)

	if payload.ID == 0 {
		return fmt.Errorf("%w: potential id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.CategoryID == 0 {
		return fmt.Errorf("%w: category_id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.Title == "" {
		return fmt.Errorf("%w: title is required", utils.ErrInvalidPayload)
	}
	if payload.Slug == "" {
		return fmt.Errorf("%w: slug is required", utils.ErrInvalidPayload)
	}
	if payload.Description == "" {
		return fmt.Errorf("%w: description is required", utils.ErrInvalidPayload)
	}
	if payload.Latitude < -90 || payload.Latitude > 90 {
		return fmt.Errorf("%w: latitude must be between -90 and 90", utils.ErrInvalidPayload)
	}
	if payload.Longitude < -180 || payload.Longitude > 180 {
		return fmt.Errorf("%w: longitude must be between -180 and 180", utils.ErrInvalidPayload)
	}
	if payload.OwnerName == "" {
		return fmt.Errorf("%w: owner_name is required", utils.ErrInvalidPayload)
	}

	media, err := utils.PrepareMedia(payload.ImgID, "uploads/potentials", 0)
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, payload, media)
}

func (s *service) Delete(ctx context.Context, payload PotentialPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: potential id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}
