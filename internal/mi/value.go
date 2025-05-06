package mi

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type Value struct {
	raw [8]byte
}

func (v *Value) As(t *Type) any {

	// fmt.Printf("parsing value as type: %d\n", *t)
	switch *t {
	case MI_BOOLEAN:
		return *(*bool)(unsafe.Pointer(&v.raw))
	case MI_UINT8:
		return *(*uint8)(unsafe.Pointer(&v.raw))
	case MI_SINT8:
		return *(*int8)(unsafe.Pointer(&v.raw))
	case MI_UINT16:
		return *(*uint16)(unsafe.Pointer(&v.raw))
	case MI_SINT16:
		return *(*int16)(unsafe.Pointer(&v.raw))
	case MI_UINT32:
		return *(*uint32)(unsafe.Pointer(&v.raw))
	case MI_SINT32:
		return *(*int32)(unsafe.Pointer(&v.raw))
	case MI_STRING:
		return windows.UTF16PtrToString((*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&v.raw[0]))))
	}

	return nil
}
