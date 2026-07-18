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
	payload.Slug = strings.TrimSpace(payload.Slug)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.OwnerName = strings.TrimSpace(payload.OwnerName)
	payload.OwnerMsisdn = strings.TrimSpace(payload.OwnerMsisdn)
	if payload.Location != nil {
		payload.Location.Title = strings.TrimSpace(payload.Location.Title)
		payload.Location.Description = strings.TrimSpace(payload.Location.Description)
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
	if payload.Slug == "" {
		return fmt.Errorf("%w: slug is required", utils.ErrInvalidPayload)
	}
	if payload.Description == "" {
		return fmt.Errorf("%w: description is required", utils.ErrInvalidPayload)
	}
	if payload.OwnerName == "" {
		return fmt.Errorf("%w: owner_name is required", utils.ErrInvalidPayload)
	}
	if err := validateLocation(payload.Location); err != nil {
		return err
	}

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/potentials", payload.UploadedBy)
	if err != nil {
		return err
	}

	return s.repository.Create(ctx, payload, media)
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
	payload.Slug = strings.TrimSpace(payload.Slug)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.OwnerName = strings.TrimSpace(payload.OwnerName)
	payload.OwnerMsisdn = strings.TrimSpace(payload.OwnerMsisdn)
	if payload.Location != nil {
		payload.Location.Title = strings.TrimSpace(payload.Location.Title)
		payload.Location.Description = strings.TrimSpace(payload.Location.Description)
	}

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
	if payload.OwnerName == "" {
		return fmt.Errorf("%w: owner_name is required", utils.ErrInvalidPayload)
	}
	if err := validateLocation(payload.Location); err != nil {
		return err
	}

	media, err := utils.PrepareMedia(payload.MediaID, "uploads/potentials", 0)
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, payload, media)
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
		Location: PotentialLocation{
			ID: item.LocationID,
		},
		Owner: PotentialOwner{
			Name:   item.OwnerName,
			Msisdn: item.OwnerMsisdn,
		},
		MediaID:   item.MediaID,
		CreatedAt: createdAt,
	}
}
