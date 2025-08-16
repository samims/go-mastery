// buffer.go
// Example: Using sync.Pool to manage reusable buffers.

package main

import (
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		fmt.Println("Creating new buffer")
		return make([]byte, 1024) // 1 KB buffer
	},
}

func main() {
	// Take buffer from pool or create new
	buf1 := bufPool.Get()
	bufByte1 := buf1.([]byte)

	fmt.Println("Got buffer 1 len:", len(bufByte1))

	// return buffer to pool
	bufPool.Put(buf1)

	// Reuse
	buf2 := bufPool.Get().([]byte)
	fmt.Println("Got buffer 2 (reused), len:", len(buf2))

}
