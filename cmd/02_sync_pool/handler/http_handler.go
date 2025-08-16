package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

type Response struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

// global sync.Pool for buffer reuse
var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// naive handler -- allocates new buffer every request
func naiveHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Hello", Count: 55}
	buf := new(bytes.Buffer) // fresh allocation every time
	_ = json.NewEncoder(buf).Encode(resp)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(buf.Bytes())
}

// pooledHandler is the optimized one, reuses buffer pool
func pooledHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "hello",
		Count:   42,
	}

	// get or create a buffer Pool
	buf := bufPool.Get().(*bytes.Buffer)
	// reset before using
	buf.Reset()

	// use the buffer
	_ = json.NewEncoder(buf).Encode(resp)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(buf.Bytes())

	// put back the buffer for reuse
	bufPool.Put(buf)
}
