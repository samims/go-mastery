package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkNaiveHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		naiveHandler(w, req)
		io.Copy(io.Discard, w.Result().Body)
		w.Result().Body.Close()
	}
}

func BenchmarkPooledHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		pooledHandler(w, req)
		io.Copy(io.Discard, w.Result().Body)
		w.Result().Body.Close()
	}
}
