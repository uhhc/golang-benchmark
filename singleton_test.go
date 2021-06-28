package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

/*
go test -bench=. -benchmem -run=none
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkGetInstanceV1-12       677327029                1.771 ns/op           0 B/op          0 allocs/op
BenchmarkGetInstanceV2-12       93841928                13.19 ns/op            0 B/op          0 allocs/op
 */

type singleton struct {
}

var (
	instance    *singleton
	mu          sync.RWMutex
	initialized uint32
)

func GetInstanceV1() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

func GetInstanceV2() *singleton {
	mu.RLock()
	r := instance
	mu.RUnlock()

	if r != nil {
		return r
	}

	mu.Lock()
	defer mu.Unlock()

	instance = &singleton{}
	return instance
}

func BenchmarkGetInstanceV1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetInstanceV1()
	}
}

func BenchmarkGetInstanceV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetInstanceV2()
	}
}
