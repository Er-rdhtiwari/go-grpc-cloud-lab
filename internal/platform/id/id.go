package id

import "github.com/google/uuid"

// Generator is a tiny interface to make ID creation deterministic in tests.
type Generator interface {
	New() string
}

type UUIDGenerator struct{}

func (UUIDGenerator) New() string { return uuid.NewString() }
