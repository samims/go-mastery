package main

type StackStruct struct {
	A [1024]int
}

// NewStack returns by value so usually stack allocated
//
//go:noinline
func NewStack() StackStruct {
	return StackStruct{}
}
