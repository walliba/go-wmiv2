package types

import (
	"unsafe"
)

type MIArray interface {
	Pointer() unsafe.Pointer
	Size() uint32
}

type Array[T any] struct {
	data *T
	size uint32
}

func (x Array[T]) Pointer() unsafe.Pointer {
	return unsafe.Pointer(x.data)
}

func (x Array[T]) Size() uint32 {
	return x.size
}

func (x Array[T]) MakeSlice() []T {
	return genericSlice[T](x)
}

func genericSlice[T any](arr MIArray) []T {
	var size = arr.Size()
	const max = 1 << 20

	// reinterpret array to slice
	mem := (*[max]T)(arr.Pointer())[:size:size]

	// make a copy of the slice (C array) since mi.dll will free the memory after iteration
	r := make([]T, size)
	_ = copy(r, mem)

	// return copy
	return r
}
