package category

import (
	"context"
	"testing"

	"cipicung.id/be/pkg/models"
)

type categoryRepositorySpy struct {
	listPayload ListCategoryPayload
}

func (r *categoryRepositorySpy) Create(context.Context, AddCategoryPayload, string) error { return nil }
func (r *categoryRepositorySpy) List(_ context.Context, payload ListCategoryPayload) ([]models.Category, error) {
	r.listPayload = payload
	return []models.Category{}, nil
}
func (r *categoryRepositorySpy) FindByID(context.Context, CategoryPayload) (*models.Category, error) {
	return nil, nil
}
func (r *categoryRepositorySpy) Update(context.Context, EditCategoryPayload, string) error {
	return nil
}
func (r *categoryRepositorySpy) Delete(context.Context, CategoryPayload) error { return nil }

func TestListNormalizesCategoryType(t *testing.T) {
	repository := &categoryRepositorySpy{}
	service := NewService(repository)

	if _, err := service.List(context.Background(), ListCategoryPayload{Type: "  Business  "}); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if repository.listPayload.Type != "business" {
		t.Fatalf("normalized type = %q, want business", repository.listPayload.Type)
	}
}
