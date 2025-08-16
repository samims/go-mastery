package _2_sync_pool

import "testing"

func BenchmarkSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pd := CreatePooled() // From 3_sync_pool
		ReleasePooled(pd)
	}
}
