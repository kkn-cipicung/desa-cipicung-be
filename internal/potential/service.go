package potential

import (
	"context"

	"cipicung.id/be/utils/file"
)

const (
	defaultPotentialLimit = 10
	maxPotentialLimit     = 100
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
	media, err := preparePotentialMedia(payload.ImgID, payload.UploadedBy)
	if err != nil {
		return err
	}

	return s.repository.Create(ctx, payload, media)
}

func (s *service) List(ctx context.Context, payload ListPotentialPayload) ([]PotentialResponse, error) {
	payload = normalizeListPotentialPayload(payload)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload PotentialPayload) (*PotentialResponse, error) {
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditPotentialPayload) error {
	media, err := preparePotentialMedia(payload.ImgID, 0)
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, payload, media)
}

func (s *service) Delete(ctx context.Context, payload PotentialPayload) error {
	return s.repository.Delete(ctx, payload)
}

func normalizeListPotentialPayload(payload ListPotentialPayload) ListPotentialPayload {
	if payload.Limit <= 0 {
		payload.Limit = defaultPotentialLimit
	}

	if payload.Limit > maxPotentialLimit {
		payload.Limit = maxPotentialLimit
	}

	if payload.Index < 0 {
		payload.Index = 0
	}

	return payload
}

func preparePotentialMedia(imgBase64 string, uploadedBy uint) (*potentialMediaPayload, error) {
	if imgBase64 == "" {
		return nil, nil
	}

	filePath, mimeType, err := file.SaveBase64(imgBase64, "uploads/potentials")
	if err != nil {
		return nil, err
	}

	media := &potentialMediaPayload{
		FilePath: filePath,
		MimeType: mimeType,
	}

	if uploadedBy != 0 {
		media.UploadedBy = &uploadedBy
	}

	return media, nil
}
