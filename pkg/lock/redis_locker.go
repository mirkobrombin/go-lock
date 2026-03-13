package lock

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	// ErrLockNotAcquired indicates that a lock could not be acquired because it is already held.
	ErrLockNotAcquired = errors.New("lock: not acquired")
	// ErrNotImplemented indicates that the requested locker is only a placeholder.
	ErrNotImplemented = errors.New("lock: not implemented")
)

// RedisLocker is a small Redis-backed locker using SET NX with an expiry.
type RedisLocker struct {
	client *redis.Client
}

// NewRedisLocker creates a Redis-backed locker.
func NewRedisLocker(client *redis.Client) *RedisLocker {
	return &RedisLocker{client: client}
}

// Acquire acquires a lock or returns ErrLockNotAcquired when it is already held.
func (r *RedisLocker) Acquire(ctx context.Context, key string, ttl time.Duration) error {
	ok, err := r.client.SetNX(ctx, key, "1", ttl).Result()
	if err != nil {
		return err
	}
	if !ok {
		return ErrLockNotAcquired
	}
	return nil
}

// TryLock attempts to acquire a lock and reports whether it succeeded.
func (r *RedisLocker) TryLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	ok, err := r.client.SetNX(ctx, key, "1", ttl).Result()
	return ok, err
}

// Release releases a lock.
func (r *RedisLocker) Release(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()
	return err
}
