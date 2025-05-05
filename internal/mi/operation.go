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

func (operation *MI_Operation) Close() uint64 {
	r0, _, _ := syscall.SyscallN(operation.ft.Close,
		uintptr(unsafe.Pointer(operation)), // [in, out] MI_Operation *operation
	)

	return uint64(r0)
}

func (operation *MI_Operation) Cancel() {
	panic("not implemented")
}

func (operation *MI_Operation) GetSession() {
	panic("not implemented")
}

func (operation *MI_Operation) GetInstance(moreResults *bool) (*MI_Instance, uint64) {
	var instance = &MI_Instance{}

	r0, _, _ := syscall.SyscallN(operation.ft.GetInstance,
		uintptr(unsafe.Pointer(operation)),   // [in] 				MI_Operation		*operation
		uintptr(unsafe.Pointer(&instance)),   // 					const MI_Instance	**instance
		uintptr(unsafe.Pointer(moreResults)), // [out, optional] 	MI_Boolean 			*moreResults
		0,                                    // [out, optional] 	MI_Result			*result
		0,                                    // 					const MI_Char		**errorMessage
		0,                                    // 					const MI_Instance	**completionDetails
	)

	return instance, uint64(r0)
}

func (operation *MI_Operation) GetIndication() {
	panic("not implemented")
}

func (operation *MI_Operation) GetClass() {
	panic("not implemented")
}
