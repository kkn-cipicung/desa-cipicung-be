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
	List(ctx context.Context, payload ListNewsPayload) ([]NewsOutput, error)
	FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsOutput, error)
	Update(ctx context.Context, payload EditNewsPayload) error
	Delete(ctx context.Context, payload NewsPayload) error
	FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsOutput, error)
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

	if payload.MediaID == nil {
		return s.repository.Create(ctx, payload, nil)
	}

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/news", payload.UploadedBy)
	if err != nil {
		return err
	}

	return s.repository.Create(ctx, payload, media)
}

func (s *service) List(ctx context.Context, payload ListNewsPayload) ([]NewsOutput, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	items, err := s.repository.List(ctx, payload)
	if err != nil {
		return nil, err
	}
	return mapNewsOutputs(items), nil
}

func (s *service) FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsOutput, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: news id must be greater than 0", utils.ErrInvalidPayload)
	}
	item, err := s.repository.FindByID(ctx, payload)
	if err != nil {
		return nil, err
	}
	output := mapNewsOutput(*item)
	return &output, nil
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

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/news", 0)
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

func (s *service) FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsOutput, error) {
	payload.Date = strings.TrimSpace(payload.Date)
	if payload.Date == "" {
		return nil, fmt.Errorf("%w: date is required", utils.ErrInvalidPayload)
	}
	if _, err := time.Parse("2006-01-02", payload.Date); err != nil {
		return nil, fmt.Errorf("%w: invalid date format", utils.ErrInvalidPayload)
	}
	items, err := s.repository.FindByDate(ctx, payload)
	if err != nil {
		return nil, err
	}
	return mapNewsOutputs(items), nil
}

func mapNewsOutputs(items []NewsResponse) []NewsOutput {
	outputs := make([]NewsOutput, 0, len(items))
	for _, item := range items {
		outputs = append(outputs, mapNewsOutput(item))
	}
	return outputs
}

func mapNewsOutput(item NewsResponse) NewsOutput {
	return NewsOutput{
		ID: item.ID,
		Category: NewsRef{
			ID:   item.CategoryID,
			Name: item.CategoryName,
		},
		Uploader: NewsRef{
			ID:   item.UploadedBy,
			Name: item.UploaderName,
		},
		Title:       item.Title,
		Description: item.Description,
		MediaID:     item.MediaID,
		CreatedAt:   utils.FormatTimestamp(item.CreatedAt),
	}
}
