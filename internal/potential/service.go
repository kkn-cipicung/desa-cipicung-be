package potential

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddPotentialPayload) error
	List(ctx context.Context, payload ListPotentialPayload) ([]PotentialOutput, error)
	FindByID(ctx context.Context, payload PotentialPayload) (*PotentialOutput, error)
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
	payload.Description = strings.TrimSpace(payload.Description)
	payload.LocationID = normalizeOptionalID(payload.LocationID)
	payload.Location = normalizeLocationInput(payload.Location)
	if payload.Location != nil && payload.Location.ID != nil {
		payload.LocationID = nil
	}

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
	if err := validateLocation(payload.Location); err != nil {
		return err
	}
	payload.Slug = utils.GenerateSlug(payload.Title)

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/potentials", payload.UploadedBy)
	if err != nil {
		return err
	}

	if err := s.repository.Create(ctx, payload, media); err != nil {
		_ = utils.RemovePreparedMedia(media)
		return err
	}
	return nil
}

func (s *service) List(ctx context.Context, payload ListPotentialPayload) ([]PotentialOutput, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	items, err := s.repository.List(ctx, payload)
	if err != nil {
		return nil, err
	}
	return mapPotentialOutputs(items), nil
}

func (s *service) FindByID(ctx context.Context, payload PotentialPayload) (*PotentialOutput, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: potential id must be greater than 0", utils.ErrInvalidPayload)
	}
	item, err := s.repository.FindByID(ctx, payload)
	if err != nil {
		return nil, err
	}
	output := mapPotentialOutput(*item)
	return &output, nil
}

func (s *service) Update(ctx context.Context, payload EditPotentialPayload) error {
	payload.Title = strings.TrimSpace(payload.Title)
	payload.Subtitle = strings.TrimSpace(payload.Subtitle)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.LocationID = normalizeOptionalID(payload.LocationID)
	payload.Location = normalizeLocationInput(payload.Location)
	if payload.Location != nil && payload.Location.ID != nil {
		payload.LocationID = nil
	}

	if payload.ID == 0 {
		return fmt.Errorf("%w: potential id must be greater than 0", utils.ErrInvalidPayload)
	}
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
	if err := validateLocation(payload.Location); err != nil {
		return err
	}
	payload.Slug = utils.GenerateSlug(payload.Title)

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/potentials", payload.UploadedBy)
	if err != nil {
		return err
	}

	if err := s.repository.Update(ctx, payload, media); err != nil {
		_ = utils.RemovePreparedMedia(media)
		return err
	}
	return nil
}

func normalizeOptionalID(id *uint) *uint {
	if id != nil && *id == 0 {
		return nil
	}
	return id
}

func normalizeLocationInput(location *PotentialLocationInput) *PotentialLocationInput {
	if location == nil {
		return nil
	}
	location.ID = normalizeOptionalID(location.ID)
	location.Title = strings.TrimSpace(location.Title)
	location.Description = strings.TrimSpace(location.Description)
	if location.ID == nil && location.Latitude == 0 && location.Longitude == 0 && location.Title == "" && location.Description == "" {
		return nil
	}
	return location
}

func validateLocation(location *PotentialLocationInput) error {
	if location == nil {
		return nil
	}
	if location.Latitude < -90 || location.Latitude > 90 {
		return fmt.Errorf("%w: latitude must be between -90 and 90", utils.ErrInvalidPayload)
	}
	if location.Longitude < -180 || location.Longitude > 180 {
		return fmt.Errorf("%w: longitude must be between -180 and 180", utils.ErrInvalidPayload)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, payload PotentialPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: potential id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func mapPotentialOutputs(items []PotentialResponse) []PotentialOutput {
	outputs := make([]PotentialOutput, 0, len(items))
	for _, item := range items {
		outputs = append(outputs, mapPotentialOutput(item))
	}
	return outputs
}

func mapPotentialOutput(item PotentialResponse) PotentialOutput {
	createdAt := ""
	if item.CreatedAt != nil {
		createdAt = utils.FormatTimestamp(*item.CreatedAt)
	}

	var location *PotentialLocation
	if item.LocationID != nil {
		location = &PotentialLocation{ID: *item.LocationID}
	}

	return PotentialOutput{
		ID: item.ID,
		Category: PotentialRef{
			ID:   item.CategoryID,
			Name: item.CategoryName,
		},
		Title:       item.Title,
		Subtitle:    item.Subtitle,
		Slug:        item.Slug,
		Description: item.Description,
		Location:    location,
		Media:       item.Media,
		CreatedAt:   createdAt,
	}
}
