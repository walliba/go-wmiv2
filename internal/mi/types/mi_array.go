package types

import (
	"unsafe"
)

type MIArray[T any] interface {
	Pointer() *T
	Size() uint32
}

type Array[T any] struct {
	data *T
	size uint32
}

func (x Array[T]) Pointer() *T {
	return x.data
}

func (x Array[T]) Size() uint32 {
	return x.size
}

func (x Array[T]) MakeSlice() []T {
	return genericSlice(x)
}

func genericSlice[T any](arr MIArray[T]) []T {
	return unsafe.Slice(arr.Pointer(), arr.Size())
}
