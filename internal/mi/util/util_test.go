package util_test

import (
	"syscall"
	"testing"
	"unsafe"

	"github.com/walliba/go-wmiv2/internal/mi/util"
)

func Test_UTF16PtrZero(t *testing.T) {
	const str = "this is my sensitive data"

	w_string, _ := syscall.UTF16PtrFromString(str)

	end := unsafe.Pointer(w_string)

	util.UTF16PtrZero(w_string)

	for range len(str) {
		if *(*uint16)(end) != 0 {
			t.Fatal("UTF16 string is not fully zeroed")
		}
	}
}
