package mi

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Value struct {
	raw *int64
}

func (v *Value) As(t Type) any {
	switch t {
	case MI_BOOLEAN:
		return *(*bool)(v.GetPointer())
	case MI_UINT8:
		return *(*uint8)(v.GetPointer())
	case MI_SINT8:
		return *(*int8)(v.GetPointer())
	case MI_UINT16:
		return *(*uint16)(v.GetPointer())
	case MI_SINT16:
		return *(*int16)(v.GetPointer())
	case MI_UINT32:
		return *(*uint32)(v.GetPointer())
	case MI_SINT32:
		return *(*int32)(v.GetPointer())
	case MI_STRING:
		ptr := v.GetPointer()
		if ptr != nil {
			return windows.UTF16PtrToString((*uint16)(*(*unsafe.Pointer)(ptr)))
		}
	default:
		fmt.Fprintf(os.Stderr, "<unsupported type: %d>", t)
	}

	return nil
}

func (v *Value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(&v.raw)
}
