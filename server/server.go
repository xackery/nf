package server

import (
	"sync"

	aasv1 "github.com/xackery/nf/proto/aas/v1"
)

// Backend implements the protobuf interface
type Backend struct {
	mu  *sync.RWMutex
	aas []*aasv1.Aa
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}
