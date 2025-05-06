package mi

import (
	"syscall"
	"unsafe"
)

type MI_Operation struct {
	reserved1 uint64
	reserved2 int64 // ptrdiff_t
	ft        *MI_OperationFT
}

type MI_OperationFT struct {
	Close         uintptr
	Cancel        uintptr
	GetSession    uintptr
	GetInstance   uintptr
	GetIndication uintptr
	GetClass      uintptr
}

var MI_OPERATION_NULL = MI_Operation{0, 0, nil}

func (o *MI_Operation) Close() Result {
	r0, _, _ := syscall.SyscallN(o.ft.Close,
		uintptr(unsafe.Pointer(o)), // [in, out] MI_Operation *operation
	)

	return Result(r0)
}

func (o *MI_Operation) Cancel() Result {
	r0, _, _ := syscall.SyscallN(o.ft.Cancel,
		uintptr(unsafe.Pointer(o)),
		0,
	)

	return Result(r0)
}

func (o *MI_Operation) GetSession() {
	panic("not implemented")
}

// Calling MI_Operation_Close before retrieving the last result where moreResults is set to MI_FALSE will cause the MI_Operation_Close function to stop responding.
func (o *MI_Operation) GetInstance(moreResults *bool, args ...any) (*MI_Instance, Result) {
	var instance = &MI_Instance{}

	r0, _, _ := syscall.SyscallN(o.ft.GetInstance,
		uintptr(unsafe.Pointer(o)),           // [in] 				MI_Operation		*operation
		uintptr(unsafe.Pointer(&instance)),   // 					const MI_Instance	**instance
		uintptr(unsafe.Pointer(moreResults)), // [out, optional] 	MI_Boolean 			*moreResults
		0,                                    // [out, optional] 	MI_Result			*result
		0,                                    // 					const MI_Char		**errorMessage
		0,                                    // 					const MI_Instance	**completionDetails
	)

	return instance, Result(r0)
}

func (o *MI_Operation) GetIndication() {
	panic("not implemented")
}

func (o *MI_Operation) GetClass() {
	panic("not implemented")
}
