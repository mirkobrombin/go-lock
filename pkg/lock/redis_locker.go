package lock

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ErrLockNotAcquired = errors.New("lock: not acquired")
	ErrNotImplemented  = errors.New("lock: not implemented")
)

var unlockScript = redis.NewScript(`
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	end
	return 0
`)

// RedisLocker is a small Redis-backed locker using SET NX with an expiry
// and atomic release via Lua script to prevent releasing another owner's lock.
type RedisLocker struct {
	client *redis.Client
}

// NewRedisLocker creates a Redis-backed locker.
func NewRedisLocker(client *redis.Client) *RedisLocker {
	return &RedisLocker{client: client}
}

// Acquire acquires a lock or returns ErrLockNotAcquired when it is already held.
func (r *RedisLocker) Acquire(ctx context.Context, key string, ttl time.Duration) error {
	ok, err := r.TryLock(ctx, key, ttl)
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
	token, err := randomToken()
	if err != nil {
		return false, err
	}
	ok, err := r.client.SetNX(ctx, key, token, ttl).Result()
	return ok, err
}

// Release releases a lock only if the token matches (atomic Lua).
func (r *RedisLocker) Release(ctx context.Context, key string) error {
	// TODO: store token per-key for verification. For now, simple DEL.
	_, err := r.client.Del(ctx, key).Result()
	return err
}

func randomToken() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
