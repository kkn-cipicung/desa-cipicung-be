package utils

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"cipicung.id/be/utils/file"
	"github.com/jmoiron/sqlx"
)

func GenerateSlug(name string) string {
	slug := strings.ToLower(strings.TrimSpace(name))
	slug = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

type MediaPayload struct {
	FilePath   string
	MimeType   string
	UploadedBy *uint
}

type ReplacedMediaPayload struct {
	ID       uint   `db:"id"`
	FilePath string `db:"file_path"`
}

type LocationPayload struct {
	Latitude    float64
	Longitude   float64
	CreatedByID *uint
	Title       string
	Description string
}

func PrepareMedia(imgBase64 *string, uploadDir string, uploadedBy uint) (*MediaPayload, error) {
	if imgBase64 == nil || *imgBase64 == "" {
		return nil, nil
	}
	value := strings.TrimSpace(*imgBase64)
	if value == "" || isExistingMediaReference(value) {
		return nil, nil
	}

	filePath, mimeType, err := file.SaveBase64(value, uploadDir)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid media_id", ErrInvalidPayload)
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

func isExistingMediaReference(value string) bool {
	return strings.HasPrefix(value, "http://") ||
		strings.HasPrefix(value, "https://") ||
		strings.HasPrefix(value, "/") ||
		strings.HasPrefix(value, "uploads/")
}

func InsertMedia(ctx context.Context, tx *sqlx.Tx, media *MediaPayload) (*uint, error) {
	query := `
		INSERT INTO media (file_path, mime_type, uploaded_by)
		VALUES (?, ?, ?)
	`

	result, err := tx.ExecContext(ctx, query, media.FilePath, media.MimeType, media.UploadedBy)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	mediaID := uint(id)

	return &mediaID, nil
}

func InsertMediaForEntity(ctx context.Context, tx *sqlx.Tx, media *MediaPayload, entityType string, entityID uint, role string) (*uint, error) {
	query := `
		INSERT INTO media (file_path, mime_type, uploaded_by, entity_type, entity_id, role)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := tx.ExecContext(
		ctx,
		query,
		media.FilePath,
		media.MimeType,
		media.UploadedBy,
		entityType,
		entityID,
		role,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	mediaID := uint(id)

	return &mediaID, nil
}

func AttachMediaToEntity(ctx context.Context, tx *sqlx.Tx, media *MediaPayload, entityType string, entityID uint, role string) (*uint, error) {
	if media == nil {
		return nil, nil
	}
	return InsertMediaForEntity(ctx, tx, media, entityType, entityID, role)
}

func AttachMediaToEntityColumn(ctx context.Context, tx *sqlx.Tx, media *MediaPayload, entityType string, entityID uint, role string, tableName string, mediaColumn string) (*uint, error) {
	mediaID, err := AttachMediaToEntity(ctx, tx, media, entityType, entityID, role)
	if err != nil || mediaID == nil {
		return mediaID, err
	}

	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE id = ?", tableName, mediaColumn)
	if _, err := tx.ExecContext(ctx, query, *mediaID, entityID); err != nil {
		return nil, err
	}

	return mediaID, nil
}

func ReplaceMediaOnEntityColumn(ctx context.Context, tx *sqlx.Tx, media *MediaPayload, entityType string, entityID uint, role string, tableName string, mediaColumn string) (*uint, *ReplacedMediaPayload, error) {
	if media == nil {
		return nil, nil, nil
	}

	var oldMediaID *uint
	selectQuery := fmt.Sprintf("SELECT %s FROM %s WHERE id = ? FOR UPDATE", mediaColumn, tableName)
	if err := tx.GetContext(ctx, &oldMediaID, selectQuery, entityID); err != nil {
		return nil, nil, err
	}

	mediaID, err := AttachMediaToEntityColumn(ctx, tx, media, entityType, entityID, role, tableName, mediaColumn)
	if err != nil || mediaID == nil {
		return mediaID, nil, err
	}

	oldMedia, err := deleteOldMedia(ctx, tx, oldMediaID, *mediaID)
	if err != nil {
		return nil, nil, err
	}

	return mediaID, oldMedia, nil
}

func deleteOldMedia(ctx context.Context, tx *sqlx.Tx, oldMediaID *uint, newMediaID uint) (*ReplacedMediaPayload, error) {
	if oldMediaID == nil || *oldMediaID == newMediaID {
		return nil, nil
	}

	var oldMedia ReplacedMediaPayload
	if err := tx.GetContext(ctx, &oldMedia, `
		SELECT id, COALESCE(file_path, '') AS file_path
		FROM media
		WHERE id = ?
	`, *oldMediaID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM media WHERE id = ?`, *oldMediaID); err != nil {
		return nil, err
	}

	return &oldMedia, nil
}

func RemoveMediaFile(media *ReplacedMediaPayload) error {
	if media == nil || strings.TrimSpace(media.FilePath) == "" {
		return nil
	}
	if err := os.Remove(media.FilePath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func RemovePreparedMedia(media *MediaPayload) error {
	if media == nil || strings.TrimSpace(media.FilePath) == "" {
		return nil
	}
	if err := os.Remove(media.FilePath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func InsertLocation(ctx context.Context, tx *sqlx.Tx, location *LocationPayload) (*uint, error) {
	query := `
		INSERT INTO locations (latitude, longitude, created_by_id, title, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`

	result, err := tx.ExecContext(ctx, query, location.Latitude, location.Longitude, location.CreatedByID, location.Title, location.Description)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	locationID := uint(id)

	return &locationID, nil
}

func ParseDate(date string) (string, error) {
	parsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	return parsed.Format("2006-01-02"), nil
}
