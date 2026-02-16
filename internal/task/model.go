package task

import "time"

// Task is our domain model.
// Keep it minimal; transport mapping (protobuf) will come later.
type Task struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
}
