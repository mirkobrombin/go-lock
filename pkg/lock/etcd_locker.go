package lock

import (
	"context"
	"time"
)

// EtcdLocker is a placeholder for an etcd lease-backed locker implementation.
type EtcdLocker struct{}

// NewEtcdLocker creates a placeholder etcd-backed locker.
func NewEtcdLocker() *EtcdLocker {
	return &EtcdLocker{}
}

// Acquire reports that the placeholder implementation is not available yet.
func (e *EtcdLocker) Acquire(ctx context.Context, key string, ttl time.Duration) error {
	return ErrNotImplemented
}

// TryLock reports that the placeholder implementation is not available yet.
func (e *EtcdLocker) TryLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return false, ErrNotImplemented
}

// Release reports that the placeholder implementation is not available yet.
func (e *EtcdLocker) Release(ctx context.Context, key string) error {
	return ErrNotImplemented
}
