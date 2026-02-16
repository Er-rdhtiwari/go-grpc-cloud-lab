package task

import "context"

// Repository is a small interface (ISP).
// WHY: service depends on behavior, not concrete storage implementation.
type Repository interface {
	Create(ctx context.Context, t Task) error
	Get(ctx context.Context, id string) (Task, error)
	List(ctx context.Context, limit int) ([]Task, error)
}
