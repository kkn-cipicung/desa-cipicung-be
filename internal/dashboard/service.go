package dashboard

import "context"

const (
	defaultDashboardLimit = 10
	maxDashboardLimit     = 100
)

type Service interface {
	Create(ctx context.Context, payload AddDashboardPayload) error
	List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error)
	FindByID(ctx context.Context, payload DashboardPayload) (*DashboardResponse, error)
	Update(ctx context.Context, payload EditDashboardPayload) error
	Delete(ctx context.Context, payload DashboardPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddDashboardPayload) error {
	return s.repository.Create(ctx, payload)
}

func (s *service) List(ctx context.Context, payload ListDashboardPayload) ([]DashboardResponse, error) {
	payload = normalizeListDashboardPayload(payload)
	return s.repository.List(ctx, payload)
}

func (s *service) FindByID(ctx context.Context, payload DashboardPayload) (*DashboardResponse, error) {
	return s.repository.FindByID(ctx, payload)
}

func (s *service) Update(ctx context.Context, payload EditDashboardPayload) error {
	return s.repository.Update(ctx, payload)
}

func (s *service) Delete(ctx context.Context, payload DashboardPayload) error {
	return s.repository.Delete(ctx, payload)
}

func normalizeListDashboardPayload(payload ListDashboardPayload) ListDashboardPayload {
	if payload.Limit <= 0 {
		payload.Limit = defaultDashboardLimit
	}

	if payload.Limit > maxDashboardLimit {
		payload.Limit = maxDashboardLimit
	}

	if payload.Index < 0 {
		payload.Index = 0
	}

	return payload
}
