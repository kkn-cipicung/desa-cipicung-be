package utils

import (
	"context"
	"strings"
	"time"

	"cipicung.id/be/utils/file"
	"github.com/jmoiron/sqlx"
)

func GenerateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

type MediaPayload struct {
	FilePath   string
	MimeType   string
	UploadedBy *uint
}

func PrepareMedia(imgBase64 string, uploadDir string, uploadedBy uint) (*MediaPayload, error) {
	if imgBase64 == "" {
		return nil, nil
	}

	filePath, mimeType, err := file.SaveBase64(imgBase64, uploadDir)
	if err != nil {
		return nil, err
	}

	media := &MediaPayload{
		FilePath: filePath,
		MimeType: mimeType,
	}

	if uploadedBy != 0 {
		media.UploadedBy = &uploadedBy
	}

	return media, nil
}

func InsertMedia(ctx context.Context, tx *sqlx.Tx, media *MediaPayload) (*uint, error) {
	query := `
		INSERT INTO media (file_path, mime_type, uploaded_by)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id uint
	if err := tx.QueryRowContext(ctx, query, media.FilePath, media.MimeType, media.UploadedBy).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func ParseDate(date string) (string, error) {
	parsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	return parsed.Format("2006-01-02"), nil
}
