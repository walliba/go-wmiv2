package mi

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/walliba/go-wmiv2/internal/mi/types"
	"golang.org/x/sys/windows"
)

type Value struct {
	raw *int64
}

func (v *Value) As(t Type) any {

	ptr := v.GetPointer()

	if ptr == nil {
		panic("<value pointer is null>")
	}

	switch t {
	case MI_BOOLEAN:
		// 0
		return *(*bool)(ptr)
	case MI_UINT8:
		// 1
		return *(*uint8)(ptr)
	case MI_SINT8:
		// 2
		return *(*int8)(ptr)
	case MI_UINT16:
		// 3
		return *(*uint16)(ptr)
	case MI_SINT16:
		// 4
		return *(*int16)(ptr)
	case MI_UINT32:
		// 5
		return *(*uint32)(ptr)
	case MI_SINT32:
		// 6
		return *(*int32)(ptr)
	case MI_UINT64:
		// 7
		return *(*uint64)(ptr)
	case MI_DATETIME:
		// 12
		return *(*types.DateTime)(ptr)
	case MI_STRING:
		// 13
		return windows.UTF16PtrToString(*(**uint16)((ptr)))
	default:
		fmt.Fprintf(os.Stderr, "<unsupported type: %d>\n", t)
	}

	return nil
}

func (v *Value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(&v.raw)
}
