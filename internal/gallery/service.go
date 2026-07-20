package gallery

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddGalleryPayload) error
	List(ctx context.Context, payload ListGalleryPayload) ([]GalleryListOutput, error)
	FindByID(ctx context.Context, payload GalleryPayload) (*GalleryDetailOutput, error)
	Update(ctx context.Context, payload EditGalleryPayload) error
	Delete(ctx context.Context, payload GalleryPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddGalleryPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Description = strings.TrimSpace(payload.Description)

	if payload.CreatedBy == 0 {
		return fmt.Errorf("%w: created_by must be greater than 0", utils.ErrInvalidPayload)
	}
	if err := validateGalleryPayload(payload.CategoryID, payload.Title, payload.Description); err != nil {
		return err
	}

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/gallery", payload.CreatedBy)
	if err != nil {
		return err
	}
	if err := s.repository.Create(ctx, payload, media); err != nil {
		_ = utils.RemovePreparedMedia(media)
		return err
	}
	return nil
}

func (s *service) List(ctx context.Context, payload ListGalleryPayload) ([]GalleryListOutput, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	items, err := s.repository.List(ctx, payload)
	if err != nil {
		return nil, err
	}

	outputs := make([]GalleryListOutput, 0, len(items))
	for _, item := range items {
		outputs = append(outputs, GalleryListOutput{
			ID:    item.ID,
			Title: item.Title,
			Image: item.Image,
		})
	}
	return outputs, nil
}

func (s *service) FindByID(ctx context.Context, payload GalleryPayload) (*GalleryDetailOutput, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: gallery id must be greater than 0", utils.ErrInvalidPayload)
	}
	item, err := s.repository.FindByID(ctx, payload)
	if err != nil {
		return nil, err
	}

	output := GalleryDetailOutput{
		Title:       item.Title,
		Image:       item.Image,
		Description: item.Description,
		Category: []GalleryRef{
			{
				ID:   item.CategoryID,
				Name: item.CategoryName,
			},
		},
	}
	return &output, nil
}

func (s *service) Update(ctx context.Context, payload EditGalleryPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Description = strings.TrimSpace(payload.Description)

	if payload.ID == 0 {
		return fmt.Errorf("%w: gallery id must be greater than 0", utils.ErrInvalidPayload)
	}
	if payload.UpdatedBy == 0 {
		return fmt.Errorf("%w: updated_by must be greater than 0", utils.ErrInvalidPayload)
	}
	if err := validateGalleryPayload(payload.CategoryID, payload.Title, payload.Description); err != nil {
		return err
	}

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/gallery", payload.UpdatedBy)
	if err != nil {
		return err
	}
	if err := s.repository.Update(ctx, payload, media); err != nil {
		_ = utils.RemovePreparedMedia(media)
		return err
	}
	return nil
}

func (s *service) Delete(ctx context.Context, payload GalleryPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: gallery id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func validateGalleryPayload(categoryID uint, title string, description string) error {
	if categoryID == 0 {
		return fmt.Errorf("%w: category_id must be greater than 0", utils.ErrInvalidPayload)
	}
	if title == "" {
		return fmt.Errorf("%w: title is required", utils.ErrInvalidPayload)
	}
	if description == "" {
		return fmt.Errorf("%w: description is required", utils.ErrInvalidPayload)
	}
	return nil
}
