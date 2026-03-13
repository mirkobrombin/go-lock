package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mirkobrombin/go-lock/pkg/lock"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	locker := lock.NewRedisLocker(client)

	if err := locker.Acquire(context.Background(), "example-resource", 5*time.Second); err != nil {
		panic(err)
	}

	if err := locker.Release(context.Background(), "example-resource"); err != nil {
		panic(err)
	}
}
