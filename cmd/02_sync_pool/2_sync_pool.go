package _2_sync_pool

import "sync"

type PooledData struct {
	value int
}

var dataPool = sync.Pool{New: func() interface{} { return new(PooledData) }}

// CreatePooled returns pool obj
func CreatePooled() *PooledData {
	pd := dataPool.Get().(*PooledData)
	pd.value = 45
	return pd
}

// ReleasePooled returns object to pool
func ReleasePooled(pd *PooledData) {
	dataPool.Put(pd)
}
