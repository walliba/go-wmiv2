package mi

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/walliba/go-wmiv2/internal/mi/types"
	"github.com/walliba/go-wmiv2/internal/mi/util"
	"golang.org/x/sys/windows"
)

type Type uint32

const (
	MI_BOOLEAN    Type = iota
	MI_UINT8                // 1
	MI_SINT8                // 2
	MI_UINT16               // 3
	MI_SINT16               // 4
	MI_UINT32               // 5
	MI_SINT32               // 6
	MI_UINT64               // 7
	MI_SINT64               // 8
	MI_REAL32               // 9
	MI_REAL64               // 10
	MI_CHAR16               // 11
	MI_DATETIME             // 12
	MI_STRING               // 13
	MI_REFERENCE            // 14
	MI_INSTANCE             // 15
	MI_BOOLEANA             // 16
	MI_UINT8A               // 17
	MI_SINT8A               // 18
	MI_UINT16A              // 19
	MI_SINT16A              // 20
	MI_UINT32A              // 21
	MI_SINT32A              // 22
	MI_UINT64A              // 23
	MI_SINT64A              // 24
	MI_REAL32A              // 25
	MI_REAL64A              // 26
	MI_CHAR16A              // 27
	MI_DATETIMEA            // 28
	MI_STRINGA              // 29
	MI_REFERENCEA           // 30
	MI_INSTANCEA            // 31
	MI_ARRAY      Type = 16 // 16
)

type Value struct {
	raw unsafe.Pointer
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
	case MI_UINT16A:
		// 19
		return (*(*types.Array[uint16])(ptr)).MakeSlice()
	case MI_STRINGA:
		// 29
		return util.UTF16PtrsToStrings((*(*types.Array[*uint16])(ptr)).MakeSlice())
	default:
		fmt.Fprintf(os.Stderr, "<unsupported type: %d>\n", t)
	}

	return nil
}

func (v *Value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(&v.raw)
}

// func MIArrayToSlice[T ](v T) []T {

// 	result := make([]T, 0, sa.GetSize())

// }
