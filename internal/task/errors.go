package task

import "errors"

// NOTE: We'll refine error strategy (wrapping/typed errors + grpc mapping) on Day 3.
// For now, sentinel errors let us use errors.Is(...) safely.
var (
	ErrInvalidTitle = errors.New("invalid title")
	ErrNotFound     = errors.New("task not found")
)
