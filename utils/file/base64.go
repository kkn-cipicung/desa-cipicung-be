package file

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveBase64(b64 string, uploadDir string) (string, string, error) {
	if b64 == "" {
		return "", "", errors.New("empty base64 string")
	}

	parts := strings.SplitN(b64, ",", 2)
	var rawBase64 string
	mimeType := "image/jpeg"

	if len(parts) == 2 {
		rawBase64 = parts[1]
		meta := parts[0]
		if strings.HasPrefix(meta, "data:") && strings.HasSuffix(meta, ";base64") {
			mimeType = meta[5 : len(meta)-7]
		}
	} else {
		rawBase64 = parts[0]
	}

	data, err := base64.StdEncoding.DecodeString(rawBase64)
	if err != nil {
		return "", "", fmt.Errorf("failed to decode base64: %w", err)
	}

	ext := ".jpg"
	if strings.Contains(mimeType, "png") {
		ext = ".png"
	} else if strings.Contains(mimeType, "gif") {
		ext = ".gif"
	} else if strings.Contains(mimeType, "webp") {
		ext = ".webp"
	}

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", "", fmt.Errorf("failed to create directory: %w", err)
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", "", fmt.Errorf("failed to write file: %w", err)
	}

	return filePath, mimeType, nil
}
