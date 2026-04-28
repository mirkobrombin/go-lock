# Go Lock

> [!CAUTION]
> go-lock is now part of the [go-foundation](https://github.com/mirkobrombin/go-foundation) framework. The v1.0.0 release mirrors go-lock v0.2.0, but future versions may introduce breaking changes. Please migrate your project.

A focused **distributed locking** toolkit for Go with a minimal Redis implementation and clear extension points for stronger backends.

## Features

- **Redis Locker:** Provides a small `SETNX`-based lock implementation.
- **Shared Interface:** Matches the `go-foundation` lock contract.
- **Explicit Placeholders:** Ships extension points for Redlock and etcd-backed implementations.
- **Library-Friendly API:** Uses simple acquire, try-lock, and release operations.

## Installation

```bash
go get github.com/mirkobrombin/go-lock
```

## Quick Start

```go
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

    if err := locker.Acquire(context.Background(), "resource", 5*time.Second); err != nil {
        panic(err)
    }
}
```

## Documentation

- [Getting Started](docs/getting-started.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
