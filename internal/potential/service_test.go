package potential

import (
	"context"
	"testing"

	"cipicung.id/be/utils"
)

type potentialRepositorySpy struct {
	createPayload AddPotentialPayload
	updatePayload EditPotentialPayload
}

func (r *potentialRepositorySpy) Create(_ context.Context, payload AddPotentialPayload, _ *utils.MediaPayload) error {
	r.createPayload = payload
	return nil
}
func (r *potentialRepositorySpy) List(context.Context, ListPotentialPayload) ([]PotentialResponse, error) {
	return []PotentialResponse{}, nil
}
func (r *potentialRepositorySpy) FindByID(context.Context, PotentialPayload) (*PotentialResponse, error) {
	return nil, nil
}
func (r *potentialRepositorySpy) Update(_ context.Context, payload EditPotentialPayload, _ *utils.MediaPayload) error {
	r.updatePayload = payload
	return nil
}
func (r *potentialRepositorySpy) Delete(context.Context, PotentialPayload) error { return nil }

func TestCreateAllowsOptionalLocationID(t *testing.T) {
	repository := &potentialRepositorySpy{}
	service := NewService(repository)
	locationID := uint(0)

	err := service.Create(context.Background(), AddPotentialPayload{
		UploadedBy:  1,
		CategoryID:  2,
		Title:       "Kerajinan",
		Slug:        "kerajinan",
		Description: "Deskripsi",
		LocationID:  &locationID,
		Location:    &PotentialLocationInput{ID: &locationID},
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if repository.createPayload.LocationID != nil {
		t.Fatalf("LocationID = %v, want nil", repository.createPayload.LocationID)
	}
	if repository.createPayload.Location != nil {
		t.Fatalf("Location = %v, want nil", repository.createPayload.Location)
	}
}

func TestUpdateAcceptsNestedLocationID(t *testing.T) {
	repository := &potentialRepositorySpy{}
	service := NewService(repository)
	locationID := uint(3)

	err := service.Update(context.Background(), EditPotentialPayload{
		ID:          1,
		UploadedBy:  1,
		CategoryID:  2,
		Title:       "Kerajinan",
		Slug:        "kerajinan",
		Description: "Deskripsi",
		Location:    &PotentialLocationInput{ID: &locationID},
	})
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	if repository.updatePayload.Location == nil || repository.updatePayload.Location.ID == nil || *repository.updatePayload.Location.ID != locationID {
		t.Fatalf("Location.ID was not preserved")
	}
}
