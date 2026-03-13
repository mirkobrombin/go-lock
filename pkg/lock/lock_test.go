package lock_test

import (
	"context"
	"errors"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	fglock "github.com/mirkobrombin/go-foundation/pkg/lock"
	"github.com/mirkobrombin/go-lock/pkg/lock"
)

func TestLockerImplementationsSatisfyInterface(t *testing.T) {
	var _ fglock.Locker = &lock.RedisLocker{}
	var _ fglock.Locker = &lock.Redlock{}
	var _ fglock.Locker = &lock.EtcdLocker{}
}

func TestRedisLockerAcquireRelease(t *testing.T) {
	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	locker := lock.NewRedisLocker(client)

	if err := locker.Acquire(context.Background(), "resource", time.Minute); err != nil {
		t.Fatalf("Acquire() error = %v", err)
	}

	if err := locker.Acquire(context.Background(), "resource", time.Minute); !errors.Is(err, lock.ErrLockNotAcquired) {
		t.Fatalf("Acquire() error = %v, want ErrLockNotAcquired", err)
	}

	if err := locker.Release(context.Background(), "resource"); err != nil {
		t.Fatalf("Release() error = %v", err)
	}

	if err := locker.Acquire(context.Background(), "resource", time.Minute); err != nil {
		t.Fatalf("Acquire() after release error = %v", err)
	}
}

func TestRedisLockerTryLock(t *testing.T) {
	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	locker := lock.NewRedisLocker(client)

	ok, err := locker.TryLock(context.Background(), "resource", time.Minute)
	if err != nil {
		t.Fatalf("TryLock() error = %v", err)
	}
	if !ok {
		t.Fatalf("TryLock() = false, want true")
	}

	ok, err = locker.TryLock(context.Background(), "resource", time.Minute)
	if err != nil {
		t.Fatalf("TryLock() second error = %v", err)
	}
	if ok {
		t.Fatalf("TryLock() second = true, want false")
	}
}

func TestPlaceholderLockersReturnNotImplemented(t *testing.T) {
	if err := lock.NewRedlock().Acquire(context.Background(), "resource", time.Second); !errors.Is(err, lock.ErrNotImplemented) {
		t.Fatalf("Redlock Acquire() error = %v, want ErrNotImplemented", err)
	}
	if err := lock.NewEtcdLocker().Release(context.Background(), "resource"); !errors.Is(err, lock.ErrNotImplemented) {
		t.Fatalf("EtcdLocker Release() error = %v, want ErrNotImplemented", err)
	}
}
