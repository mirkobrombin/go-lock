# Getting Started

`go-lock` is intentionally small. It offers a ready-to-use Redis lock implementation and reserves richer distributed strategies for future backends.

## Included implementations

- `lock.RedisLocker` uses `SETNX` with a TTL.
- `lock.Redlock` documents the contract for a future multi-node Redis implementation.
- `lock.EtcdLocker` documents the contract for an etcd lease-backed implementation.

## Operational guidance

The Redis implementation is suitable for lightweight coordination. If you need stronger guarantees across multiple Redis nodes, build a dedicated Redlock implementation or integrate a battle-tested library behind the same interface.
