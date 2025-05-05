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
	r0, _, _ := syscall.SyscallN(operation.ft.Close, uintptr(unsafe.Pointer(operation)))

	return uint64(r0)
}

func (operation *MI_Operation) GetInstance(moreResults *bool) (*MI_Instance, uint64) {
	var instance = &MI_Instance{}

	// ptrToInst := &instance
	// var moreResults *bool
	// var errorMessage **string
	// var completionDetails **MI_Instance

	r0, _, _ := syscall.SyscallN(operation.ft.GetInstance,
		uintptr(unsafe.Pointer(operation)),   // [in] *operation
		uintptr(unsafe.Pointer(&instance)),   // **instance
		uintptr(unsafe.Pointer(moreResults)), // [out, optional] *moreResults
		0,                                    // [out, optional] *result
		0,                                    // **errorMessage
		0,                                    // **completionDetails
	)

	return instance, uint64(r0)
}
