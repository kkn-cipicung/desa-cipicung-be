package mapdata

import (
	"context"
	"fmt"
	"strings"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddMapPayload) error
	Detail(ctx context.Context) (*MapOutput, error)
	FindActive(ctx context.Context) (*MapOutput, error)
	Update(ctx context.Context, payload EditMapPayload) error
	Activate(ctx context.Context, payload MapPayload) error
	Delete(ctx context.Context, payload MapPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddMapPayload) error {
	normalizeMapPayload(&payload)
	if err := validateMapPayload(payload); err != nil {
		return err
	}
	return s.repository.Create(ctx, payload)
}

func (s *service) Detail(ctx context.Context) (*MapOutput, error) {
	item, err := s.repository.Detail(ctx)
	if err != nil {
		return nil, err
	}
	output := mapMapOutput(*item)
	return &output, nil
}

func (s *service) FindActive(ctx context.Context) (*MapOutput, error) {
	item, err := s.repository.FindActive(ctx)
	if err != nil {
		return nil, err
	}
	output := mapMapOutput(*item)
	return &output, nil
}

func (s *service) Update(ctx context.Context, payload EditMapPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: map id must be greater than 0", utils.ErrInvalidPayload)
	}
	normalizeMapPayload(&payload.AddMapPayload)
	if err := validateMapPayload(payload.AddMapPayload); err != nil {
		return err
	}
	return s.repository.Update(ctx, payload)
}

func (s *service) Activate(ctx context.Context, payload MapPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: map id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Activate(ctx, payload)
}

func (s *service) Delete(ctx context.Context, payload MapPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: map id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func normalizeMapPayload(payload *AddMapPayload) {
	payload.Elevation = strings.TrimSpace(payload.Elevation)
	payload.Coordinate = strings.TrimSpace(payload.Coordinate)
}

func validateMapPayload(payload AddMapPayload) error {
	if payload.Elevation == "" {
		return fmt.Errorf("%w: elevation is required", utils.ErrInvalidPayload)
	}
	if payload.Coordinate == "" {
		return fmt.Errorf("%w: coordinate is required", utils.ErrInvalidPayload)
	}
	if payload.HamletOne == nil {
		return fmt.Errorf("%w: hamlet_one is required", utils.ErrInvalidPayload)
	}
	if *payload.HamletOne < 0 {
		return fmt.Errorf("%w: hamlet_one cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.HamletTwo == nil {
		return fmt.Errorf("%w: hamlet_two is required", utils.ErrInvalidPayload)
	}
	if *payload.HamletTwo < 0 {
		return fmt.Errorf("%w: hamlet_two cannot be negative", utils.ErrInvalidPayload)
	}
	return nil
}

func mapMapOutput(item MapResponse) MapOutput {
	return MapOutput{
		Elevation:  item.Elevation,
		Coordinate: item.Coordinate,
		HamletOne:  item.HamletOne,
		HamletTwo:  item.HamletTwo,
		Population: item.HamletOne + item.HamletTwo,
	}
}
