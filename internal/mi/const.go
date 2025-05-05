package mi

import (
	"syscall"
	"unsafe"
)

type MI_RESULT uint64

const (
	MI_RESULT_OK                                  MI_RESULT = iota // 0
	MI_RESULT_FAILED                                               // 1
	MI_RESULT_ACCESS_DENIED                                        // 2
	MI_RESULT_INVALID_NAMESPACE                                    // 3
	MI_RESULT_INVALID_PARAMETER                                    // 4
	MI_RESULT_INVALID_CLASS                                        // 5
	MI_RESULT_NOT_FOUND                                            // 6
	MI_RESULT_NOT_SUPPORTED                                        // 7
	MI_RESULT_CLASS_HAS_CHILDREN                                   // 8
	MI_RESULT_CLASS_HAS_INSTANCES                                  // 9
	MI_RESULT_INVALID_SUPERCLASS                                   // 10
	MI_RESULT_ALREADY_EXISTS                                       // 11
	MI_RESULT_NO_SUCH_PROPERTY                                     // 12
	MI_RESULT_TYPE_MISMATCH                                        // 13
	MI_RESULT_QUERY_LANGUAGE_NOT_SUPPORTED                         // 14
	MI_RESULT_INVALID_QUERY                                        // 15
	MI_RESULT_METHOD_NOT_AVAILABLE                                 // 16
	MI_RESULT_METHOD_NOT_FOUND                                     // 17
	MI_RESULT_NAMESPACE_NOT_EMPTY                                  // 20
	MI_RESULT_INVALID_ENUMERATION_CONTEXT                          // 21
	MI_RESULT_INVALID_OPERATION_TIMEOUT                            // 22
	MI_RESULT_PULL_HAS_BEEN_ABANDONED                              // 23
	MI_RESULT_PULL_CANNOT_BE_ABANDONED                             // 24
	MI_RESULT_FILTERED_ENUMERATION_NOT_SUPPORTED                   // 25
	MI_RESULT_CONTINUATION_ON_ERROR_NOT_SUPPORTED                  // 26
	MI_RESULT_SERVER_LIMITS_EXCEEDED                               // 27
	MI_RESULT_SERVER_IS_SHUTTING_DOWN                              // 28
)

type MI_Value struct {
	raw [8]byte
}

func (v *MI_Value) ToString() string {
	ptr := *(*uintptr)(unsafe.Pointer(&v.raw[0]))
	return UTF16PtrToString((*uint16)(unsafe.Pointer(ptr)))
}

func UTF16PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	end := unsafe.Pointer(p)
	n := 0
	for *(*uint16)(end) != 0 {
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
		n++
	}
	return syscall.UTF16ToString(unsafe.Slice(p, n))
}

// type MI_ENUMERATION

// type MI_ClassDecl struct {
// 	flags uint32
// 	code  uint32
// 	name  *uint16
// }
