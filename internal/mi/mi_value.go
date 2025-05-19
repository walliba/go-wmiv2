package mi

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/walliba/go-wmiv2/internal/mi/types"
	"github.com/walliba/go-wmiv2/internal/mi/util"
)

// Type is a type alias for uint32
//
// According to x64 MSVC 14.38 (per mi.dll PE headers), the enum should be an `unsigned __int64` (uint64), but this should be fine as the enum only contains 32
type Type = uint32

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

// Value resembles an MI_Value union allowing for arbitrary type casting
type Value struct {
	// The address of the arbitrary type
	raw [36]byte
}

// As reinterpret casts, or allocates (make), the memory at v.raw with the Go equivalent denoted by t
func (v *Value) As(t Type) any {
	ptr := v.GetPointer()

	if ptr == nil {
		panic("<value pointer is null>")
	}

	switch t {
	case MI_BOOLEAN:
		return *(*bool)(ptr)
	case MI_UINT8:
		return *(*uint8)(ptr)
	case MI_SINT8:
		return *(*int8)(ptr)
	case MI_UINT16:
		return *(*uint16)(ptr)
	case MI_SINT16:
		return *(*int16)(ptr)
	case MI_UINT32:
		return *(*uint32)(ptr)
	case MI_SINT32:
		return *(*int32)(ptr)
	case MI_UINT64:
		return *(*uint64)(ptr)
	case MI_SINT64:
		return *(*int64)(ptr)
	case MI_REAL32:
		return *(*float32)(ptr)
	case MI_REAL64:
		return *(*float64)(ptr)
	case MI_CHAR16:
		return *(*rune)(ptr) // rune is alias for int32
	case MI_DATETIME:
		return *(*types.DateTime)(ptr)
	case MI_STRING:
		return util.UTF16PtrToString(*(**uint16)(ptr))
	case MI_REFERENCE:
		return "<not_implemented>"
	case MI_INSTANCE:
		return (*(**Instance)(ptr)).String() // workaround for premature GC issue

	// Arrays
	case MI_BOOLEANA:
		return (*(*types.Array[bool])(ptr)).MakeSlice()
	case MI_UINT8A:
		return (*(*types.Array[uint8])(ptr)).MakeSlice()
	case MI_SINT8A:
		return (*(*types.Array[int8])(ptr)).MakeSlice()
	case MI_UINT16A:
		return (*(*types.Array[uint16])(ptr)).MakeSlice()
	case MI_SINT16A:
		return (*(*types.Array[int16])(ptr)).MakeSlice()
	case MI_UINT32A:
		return (*(*types.Array[uint32])(ptr)).MakeSlice()
	case MI_SINT32A:
		return (*(*types.Array[int32])(ptr)).MakeSlice()
	case MI_UINT64A:
		return (*(*types.Array[uint64])(ptr)).MakeSlice()
	case MI_SINT64A:
		return (*(*types.Array[int64])(ptr)).MakeSlice()
	case MI_REAL32A:
		return (*(*types.Array[float32])(ptr)).MakeSlice()
	case MI_REAL64A:
		return (*(*types.Array[float64])(ptr)).MakeSlice()
	case MI_CHAR16A:
		return (*(*types.Array[rune])(ptr)).MakeSlice()
	case MI_DATETIMEA:
		return (*(*types.Array[types.DateTime])(ptr)).MakeSlice()
	case MI_STRINGA:
		return util.UTF16PtrsToStrings((*(*types.Array[*uint16])(ptr)).MakeSlice())
	case MI_REFERENCEA, MI_INSTANCEA:
		return "<not_implemented>"

	default:
		fmt.Fprintf(os.Stderr, "<unsupported type: %d>\n", t)
		return nil
	}
}

// GetPointer retrieves a pointer to the address of the underlying raw data
func (v *Value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(&v.raw)
}
