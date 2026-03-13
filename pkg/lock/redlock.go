package lock

import (
	"context"
	"time"
)

// Redlock is a placeholder for a multi-node Redis lock implementation.
type Redlock struct{}

// NewRedlock creates a placeholder Redlock implementation.
func NewRedlock() *Redlock {
	return &Redlock{}
}

// Acquire reports that the placeholder implementation is not available yet.
func (r *Redlock) Acquire(ctx context.Context, key string, ttl time.Duration) error {
	return ErrNotImplemented
}

// TryLock reports that the placeholder implementation is not available yet.
func (r *Redlock) TryLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return false, ErrNotImplemented
}

// Release reports that the placeholder implementation is not available yet.
func (r *Redlock) Release(ctx context.Context, key string) error {
	return ErrNotImplemented
}
