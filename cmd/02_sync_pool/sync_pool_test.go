package main

import (
	"sync"
	"testing"
)

func TestBufferPoolBasic(t *testing.T) {
	var bufPool = sync.Pool{
		New: func() any {
			return make([]byte, 1024) // always 1KB buffer
		},
	}

	// First Get -> should create new
	buf1 := bufPool.Get().([]byte)
	if len(buf1) != 1024 {
		t.Errorf("expected buffer length 1024, got %d", len(buf1))
	}

	// put back
	bufPool.Put(buf1)

	// Next Get → should reuse
	buf2 := bufPool.Get().([]byte)
	if len(buf2) != 1024 {
		t.Errorf("expected buffer length 1024, got %d", len(buf2))
	}

}

func BenchmarkBufferPoolWithCorrectSyncPool(b *testing.B) {
	var bufPool = sync.Pool{New: func() any {
		return make([]byte, 1024)
	}}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := bufPool.Get().([]byte)
			bufPool.Put(buf)
			buf2 := bufPool.Get().([]byte)
			bufPool.Put(buf2)

		}
	})
}

// BenchmarkBufferPoolWithWrongSyncPool test without putting back the pool obj
// to the pool, which causes repeated creation rather than reuse
func BenchmarkBufferPoolWithWrongSyncPool(b *testing.B) {
	var bufPool = sync.Pool{New: func() any {
		return make([]byte, 1024)
	}}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// gets
			_ = bufPool.Get().([]byte)
			// again gets without putting
			_ = bufPool.Get().([]byte)
			_ = bufPool.Get().([]byte)

		}
	})
}
