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

func (x *Array[T]) Pointer() unsafe.Pointer {
	return unsafe.Pointer(x.data)
}

func (x *Array[T]) Size() uint32 {
	return x.size
}

func (x *Array[T]) MakeSlice() []T {
	return genericSlice[T](x)
}

func genericSlice[T any](arr MIArray) []T {
	size := arr.Size()

	// set slice backing array
	mem := unsafe.Slice((*T)(arr.Pointer()), size)

	// make a copy of the slice to avoid use-after-free
	r := make([]T, size)
	_ = copy(r, mem)

	return r
}
