package news

import (
	"context"

	"cipicung.id/be/utils/file"
)

const (
	defaultNewsLimit = 10
	maxNewsLimit     = 100
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
	media, err := prepareNewsMedia(payload.ImgID, payload.UploadedBy)
	if err != nil {
		return err
	}

	return s.repository.Create(ctx, payload, media)
}

func (s *service) List(ctx context.Context, payload ListNewsPayload) ([]NewsResponse, error) {
	payload = normalizeListNewsPayload(payload)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload NewsByIdPayload) (*NewsResponse, error) {
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditNewsPayload) error {
	media, err := prepareNewsMedia(payload.ImgID, 0)
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, payload, media)
}

func (s *service) Delete(ctx context.Context, payload NewsPayload) error {
	return s.repository.Delete(ctx, payload)
}

func normalizeListNewsPayload(payload ListNewsPayload) ListNewsPayload {
	if payload.Limit <= 0 {
		payload.Limit = defaultNewsLimit
	}

	if payload.Limit > maxNewsLimit {
		payload.Limit = maxNewsLimit
	}

	if payload.Index < 0 {
		payload.Index = 0
	}

	return payload
}

func prepareNewsMedia(imgBase64 string, uploadedBy uint) (*newsMediaPayload, error) {
	if imgBase64 == "" {
		return nil, nil
	}

	filePath, mimeType, err := file.SaveBase64(imgBase64, "uploads/news")
	if err != nil {
		return nil, err
	}

	media := &newsMediaPayload{
		FilePath: filePath,
		MimeType: mimeType,
	}

	if uploadedBy != 0 {
		media.UploadedBy = &uploadedBy
	}

	return media, nil
}

func (s *service) FindByDate(ctx context.Context, payload NewsByDatePayload) ([]NewsResponse, error) {
	return s.repository.FindByDate(ctx, payload)
}
