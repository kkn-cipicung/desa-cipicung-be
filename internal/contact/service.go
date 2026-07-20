package contact

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddContactPayload) error
	Detail(ctx context.Context) (*ContactOutput, error)
	Update(ctx context.Context, payload EditContactPayload) error
	Delete(ctx context.Context, payload ContactPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddContactPayload) error {
	normalizeContactPayload(&payload)
	if err := validateContactPayload(payload); err != nil {
		return err
	}
	return s.repository.Create(ctx, payload)
}

func (s *service) Detail(ctx context.Context) (*ContactOutput, error) {
	item, err := s.repository.Detail(ctx)
	if err != nil {
		return nil, err
	}

	return &ContactOutput{
		Office: ContactOffice{
			Name:       fmt.Sprintf("Kantor Kepala Desa %s", item.Name),
			Address:    item.Address,
			District:   item.District,
			Regency:    item.Regency,
			Province:   item.Province,
			PostalCode: item.PostalCode,
		},
		Contact: ContactInfo{
			Email:   item.Email,
			Phone:   item.Phone,
			Website: item.Website,
		},
		SocialMedia: []ContactSocialMedia{},
		ServiceHour: []ContactServiceHour{
			{Day: "Senin-Kamis", Time: "08.00-15.00"},
			{Day: "Jumat", Time: "08.00-11.30"},
			{Day: "Sabtu-Minggu", Time: "Tutup"},
		},
	}, nil
}

func (s *service) Update(ctx context.Context, payload EditContactPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: contact id must be greater than 0", utils.ErrInvalidPayload)
	}
	normalizeContactPayload(&payload.AddContactPayload)
	if err := validateContactPayload(payload.AddContactPayload); err != nil {
		return err
	}
	return s.repository.Update(ctx, payload)
}

func (s *service) Delete(ctx context.Context, payload ContactPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: contact id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func normalizeContactPayload(payload *AddContactPayload) {
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Province = strings.TrimSpace(payload.Province)
	payload.Regency = strings.TrimSpace(payload.Regency)
	payload.District = strings.TrimSpace(payload.District)
	payload.PostalCode = strings.TrimSpace(payload.PostalCode)
	payload.Address = strings.TrimSpace(payload.Address)
	payload.Phone = strings.TrimSpace(payload.Phone)
	payload.Email = strings.TrimSpace(payload.Email)
	payload.Website = strings.TrimSpace(payload.Website)
}

func validateContactPayload(payload AddContactPayload) error {
	if payload.Name == "" {
		return fmt.Errorf("%w: name is required", utils.ErrInvalidPayload)
	}
	if payload.Province == "" {
		return fmt.Errorf("%w: province is required", utils.ErrInvalidPayload)
	}
	if payload.Regency == "" {
		return fmt.Errorf("%w: regency is required", utils.ErrInvalidPayload)
	}
	if payload.District == "" {
		return fmt.Errorf("%w: district is required", utils.ErrInvalidPayload)
	}
	if payload.Address == "" {
		return fmt.Errorf("%w: address is required", utils.ErrInvalidPayload)
	}
	return nil
}
