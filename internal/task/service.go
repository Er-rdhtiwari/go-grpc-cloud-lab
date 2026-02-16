package task

import (
	"context"
	"strings"

	"go-grpc-cloud-lab/internal/platform/clock"
	"go-grpc-cloud-lab/internal/platform/id"
)

// Service is the business-facing contract.
// Transport (gRPC handlers) will call this, not storage directly.
type Service interface {
	CreateTask(ctx context.Context, title string) (Task, error)
	GetTask(ctx context.Context, id string) (Task, error)
	ListTasks(ctx context.Context, limit int) ([]Task, error)
}

// service is the concrete implementation.
// pointer receiver: standard for services (mutability-ready + consistent method set).
type service struct {
	repo  Repository
	ids   id.Generator
	clock clock.Clock
}

type Option func(*service)

// NewService: explicit wiring, no magic DI.
// Options pattern: keeps constructor stable as deps grow.
func NewService(repo Repository, ids id.Generator, clk clock.Clock, opts ...Option) Service {
	s := &service{
		repo:  repo,
		ids:   ids,
		clock: clk,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *service) CreateTask(ctx context.Context, title string) (Task, error) {
	// Always accept context: allows cancellation/timeouts from gRPC layer.
	title = strings.TrimSpace(title)
	if title == "" || len(title) > 200 {
		return Task{}, ErrInvalidTitle
	}

	t := Task{
		ID:        s.ids.New(),
		Title:     title,
		Done:      false,
		CreatedAt: s.clock.Now(),
	}

	if err := s.repo.Create(ctx, t); err != nil {
		return Task{}, err
	}
	return t, nil
}

func (s *service) GetTask(ctx context.Context, id string) (Task, error) {
	if strings.TrimSpace(id) == "" {
		return Task{}, ErrNotFound // keep simple for now; Day 3 will distinguish InvalidArgument vs NotFound
	}
	return s.repo.Get(ctx, id)
}

func (s *service) ListTasks(ctx context.Context, limit int) ([]Task, error) {
	// Defensive default: avoids accidental “return everything”.
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	return s.repo.List(ctx, limit)
}
