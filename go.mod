module github.com/mirkobrombin/go-lock

go 1.24.0

toolchain go1.24.12

require (
	github.com/alicebob/miniredis/v2 v2.37.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/mirkobrombin/go-foundation v0.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/yuin/gopher-lua v1.1.1 // indirect
)

replace github.com/mirkobrombin/go-foundation => ../go-foundation
