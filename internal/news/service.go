package news

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddNewsPayload) error
	List(ctx context.Context, payload ListNewsPayload) ([]NewsResponse, error)
	FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsResponse, error)
	Update(ctx context.Context, payload EditNewsPayload) error
	Delete(ctx context.Context, payload NewsPayload) error
	FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddNewsPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Description = strings.TrimSpace(payload.Description)

	if payload.UploadedBy == 0 {
		return fmt.Errorf("%w: uploaded_by must be greater than 0", utils.ErrInvalidPayload)
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

	media, err := utils.PrepareMedia(payload.ImgID, "uploads/news", payload.UploadedBy)
	if err != nil {
		return err
	}

	return s.repository.Create(ctx, payload, media)
}

func (s *service) List(ctx context.Context, payload ListNewsPayload) ([]NewsResponse, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsResponse, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: news id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditNewsPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Description = strings.TrimSpace(payload.Description)

	if payload.ID == 0 {
		return fmt.Errorf("%w: news id must be greater than 0", utils.ErrInvalidPayload)
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

	media, err := utils.PrepareMedia(payload.ImgID, "uploads/news", 0)
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, payload, media)
}

func (s *service) Delete(ctx context.Context, payload NewsPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: news id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func (s *service) FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsResponse, error) {
	payload.Date = strings.TrimSpace(payload.Date)
	if payload.Date == "" {
		return nil, fmt.Errorf("%w: date is required", utils.ErrInvalidPayload)
	}
	if _, err := time.Parse("2006-01-02", payload.Date); err != nil {
		return nil, fmt.Errorf("%w: invalid date format", utils.ErrInvalidPayload)
	}
	return s.repository.FindByDate(ctx, payload)
}
