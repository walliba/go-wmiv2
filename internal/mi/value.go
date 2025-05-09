package mi

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/walliba/go-wmiv2/internal/mi/types"
	"github.com/walliba/go-wmiv2/internal/mi/util"
	"golang.org/x/sys/windows"
)

type Type = int

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
	raw unsafe.Pointer
}

type decoderFn func(ptr unsafe.Pointer) any

var decoderTable = [...]decoderFn{
	MI_BOOLEAN:    func(p unsafe.Pointer) any { return *(*bool)(p) },
	MI_UINT8:      func(p unsafe.Pointer) any { return *(*uint8)(p) },
	MI_SINT8:      func(p unsafe.Pointer) any { return *(*int8)(p) },
	MI_UINT16:     func(p unsafe.Pointer) any { return *(*uint16)(p) },
	MI_SINT16:     func(p unsafe.Pointer) any { return *(*int16)(p) },
	MI_UINT32:     func(p unsafe.Pointer) any { return *(*uint32)(p) },
	MI_SINT32:     func(p unsafe.Pointer) any { return *(*int32)(p) },
	MI_UINT64:     func(p unsafe.Pointer) any { return *(*uint64)(p) },
	MI_SINT64:     func(p unsafe.Pointer) any { return *(*int64)(p) },
	MI_REAL32:     func(p unsafe.Pointer) any { return *(*float32)(p) },
	MI_REAL64:     func(p unsafe.Pointer) any { return *(*float64)(p) },
	MI_CHAR16:     func(p unsafe.Pointer) any { return *(*rune)(p) }, // I don't know if this actually works; I can't find a class that uses MI_CHAR16
	MI_DATETIME:   func(p unsafe.Pointer) any { return *(*types.DateTime)(p) },
	MI_STRING:     func(p unsafe.Pointer) any { return windows.UTF16PtrToString(*(**uint16)(p)) },
	MI_REFERENCE:  func(p unsafe.Pointer) any { return "<not_implemented>" },
	MI_INSTANCE:   func(p unsafe.Pointer) any { return (*(**Instance)(p)).String() }, // Embedded instances get their memory prematurely freed, so just passing the string until I find a solution that isn't cloning
	MI_BOOLEANA:   func(p unsafe.Pointer) any { return (*(*types.Array[bool])(p)).MakeSlice() },
	MI_UINT8A:     func(p unsafe.Pointer) any { return (*(*types.Array[uint8])(p)).MakeSlice() },
	MI_SINT8A:     func(p unsafe.Pointer) any { return (*(*types.Array[int8])(p)).MakeSlice() },
	MI_UINT16A:    func(p unsafe.Pointer) any { return (*(*types.Array[uint16])(p)).MakeSlice() },
	MI_SINT16A:    func(p unsafe.Pointer) any { return (*(*types.Array[int16])(p)).MakeSlice() },
	MI_UINT32A:    func(p unsafe.Pointer) any { return (*(*types.Array[uint32])(p)).MakeSlice() },
	MI_SINT32A:    func(p unsafe.Pointer) any { return (*(*types.Array[int32])(p)).MakeSlice() },
	MI_UINT64A:    func(p unsafe.Pointer) any { return (*(*types.Array[uint64])(p)).MakeSlice() },
	MI_SINT64A:    func(p unsafe.Pointer) any { return (*(*types.Array[int64])(p)).MakeSlice() },
	MI_REAL32A:    func(p unsafe.Pointer) any { return (*(*types.Array[float32])(p)).MakeSlice() },
	MI_REAL64A:    func(p unsafe.Pointer) any { return (*(*types.Array[float64])(p)).MakeSlice() },
	MI_CHAR16A:    func(p unsafe.Pointer) any { return (*(*types.Array[rune])(p)).MakeSlice() }, // I don't know if this actually works; I can't find a class that uses MI_CHAR16[]
	MI_DATETIMEA:  func(p unsafe.Pointer) any { return (*(*types.Array[types.DateTime])(p)).MakeSlice() },
	MI_STRINGA:    func(p unsafe.Pointer) any { return util.UTF16PtrsToStrings((*(*types.Array[*uint16])(p)).MakeSlice()) },
	MI_REFERENCEA: func(p unsafe.Pointer) any { return "<not_implemented>" },
	MI_INSTANCEA:  func(p unsafe.Pointer) any { return "<not_implemented>" },
}

// As reinterpret casts, or allocates (make), the memory at v.raw with the Go equivalent denoted by t
func (v *Value) As(t Type) any {

	ptr := v.GetPointer()

	// Can this be nil..?
	if ptr == nil {
		panic("<value pointer is null>")
	}

	if t > len(decoderTable)-1 {
		fmt.Fprintf(os.Stderr, "<unsupported type: %d>\n", t)
		return nil
	}

	return decoderTable[t](ptr)
}

// GetPointer retrieves a pointer to the address of the underlying raw data
func (v Value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(&v.raw)
}
