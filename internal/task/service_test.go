package task

import (
	"context"
	"errors"
	"testing"
	"time"
)

// ---- Manual stubs (no framework) ----

// stubRepo lets tests control repo behavior.
type stubRepo struct {
	createFn func(context.Context, Task) error
	getFn    func(context.Context, string) (Task, error)
	listFn   func(context.Context, int) ([]Task, error)
}

func (r stubRepo) Create(ctx context.Context, t Task) error {
	if r.createFn == nil {
		return nil
	}
	return r.createFn(ctx, t)
}
func (r stubRepo) Get(ctx context.Context, id string) (Task, error) {
	if r.getFn == nil {
		return Task{}, nil
	}
	return r.getFn(ctx, id)
}
func (r stubRepo) List(ctx context.Context, limit int) ([]Task, error) {
	if r.listFn == nil {
		return nil, nil
	}
	return r.listFn(ctx, limit)
}

// fake ID generator + fake clock for determinism
type fakeIDs struct{ v string }
func (f fakeIDs) New() string { return f.v }

type fakeClock struct{ t time.Time }
func (f fakeClock) Now() time.Time { return f.t }

// ---- Tests ----

func TestCreateTask_TableDriven(t *testing.T) {
	fixedTime := time.Date(2026, 2, 15, 10, 0, 0, 0, time.UTC)

	cases := []struct {
		name      string
		title     string
		wantErr   error
		wantTitle string
	}{
		{"empty title", "", ErrInvalidTitle, ""},
		{"spaces title", "   ", ErrInvalidTitle, ""},
		{"valid title trims", "  learn go  ", nil, "learn go"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var created Task

			repo := stubRepo{
				createFn: func(_ context.Context, got Task) error {
					created = got
					return nil
				},
			}

			svc := NewService(repo, fakeIDs{v: "t-123"}, fakeClock{t: fixedTime})

			got, err := svc.CreateTask(context.Background(), tc.title)

			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("expected error %v, got %v", tc.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.ID != "t-123" || got.CreatedAt != fixedTime {
				t.Fatalf("unexpected task meta: got ID=%s time=%v", got.ID, got.CreatedAt)
			}
			if got.Title != tc.wantTitle {
				t.Fatalf("expected title %q, got %q", tc.wantTitle, got.Title)
			}
			// ensure repo saw the same data
			if created != got {
				t.Fatalf("repo got %+v, want %+v", created, got)
			}
		})
	}
}

func TestGetTask_NotFoundPropagates(t *testing.T) {
	repo := stubRepo{
		getFn: func(_ context.Context, _ string) (Task, error) {
			return Task{}, ErrNotFound
		},
	}
	svc := NewService(repo, fakeIDs{v: "x"}, fakeClock{t: time.Now().UTC()})

	_, err := svc.GetTask(context.Background(), "missing-id")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
