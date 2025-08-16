package main

type Data struct {
	value int
}

type HeapStruct struct {
	A [1024]int
}

//go:noinline
func NewHeap() *HeapStruct {
	return &HeapStruct{} // escape to heap, pointer return
}
