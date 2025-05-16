package util

import (
	"syscall"
	"unsafe"
)

// Converts a []*uint16 to []string
func UTF16PtrsToStrings(ptrs []*uint16) []string {
	result := make([]string, len(ptrs))
	for i, p := range ptrs {
		if p != nil {
			result[i] = UTF16PtrToString(p)
		}
	}
	return result
}

// Copied to avoid dependency bloat; slightly modified
// https://cs.opensource.google/go/x/sys/+/refs/tags/v0.32.0:windows/syscall_windows.go;l=117
func UTF16PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}

	if *p == 0 {
		return ""
	}

	end := unsafe.Pointer(p)
	n := 0
	for *(*uint16)(end) != 0 {
		end = unsafe.Add(end, unsafe.Sizeof(*p))
		n++
	}
	return syscall.UTF16ToString(unsafe.Slice(p, n))
}

// UTF16PtrZero zeroes a UTF-16 string, effectively setting all characters to 0
func UTF16PtrZero(p *uint16) {
	if p == nil {
		return
	}

	if *p == 0 {
		return
	}

	end := unsafe.Pointer(p)

	for *(*uint16)(end) != 0 {
		*(*uint16)(end) = 0
		end = unsafe.Add(end, unsafe.Sizeof(*p))
	}
}
