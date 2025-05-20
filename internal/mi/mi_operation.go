package mi

import (
	"syscall"
	"unsafe"
)

type Operation struct {
	reserved1 uint64
	reserved2 ptrdiff_t
	ft        *operationFT
}

type operationFT struct {
	close         uintptr
	cancel        uintptr
	getSession    uintptr
	getInstance   uintptr
	getIndication uintptr
	getClass      uintptr
}

var MI_OPERATION_NULL = Operation{0, 0, nil}

func (o *Operation) Close() Result {
	r0, _, _ := syscall.SyscallN(o.ft.close,
		uintptr(unsafe.Pointer(o)), // [in, out] MI_Operation *operation
	)

	return Result(r0)
}

func (o *Operation) Cancel() Result {
	r0, _, _ := syscall.SyscallN(o.ft.cancel,
		uintptr(unsafe.Pointer(o)),
		0,
	)

	return Result(r0)
}

func (o *Operation) GetSession() {
	panic("not implemented")
}

// Calling MI_Operation_Close before retrieving the last result where moreResults is set to MI_FALSE will cause the MI_Operation_Close function to stop responding.
func (o *Operation) GetInstance(moreResults *bool, args ...any) (*Instance, Result) {
	instance := &Instance{}

	r0, _, _ := syscall.SyscallN(o.ft.getInstance,
		uintptr(unsafe.Pointer(o)),           // [in] 				MI_Operation		*operation
		uintptr(unsafe.Pointer(&instance)),   // 					const MI_Instance	**instance
		uintptr(unsafe.Pointer(moreResults)), // [out, optional] 	MI_Boolean 			*moreResults
		0,                                    // [out, optional] 	MI_Result			*result
		0,                                    // 					const MI_Char		**errorMessage
		0,                                    // 					const MI_Instance	**completionDetails
	)

	return instance, Result(r0)
}

func (o *Operation) GetIndication() {
	panic("not implemented")
}

/*
GetClass signature

	[in]		operation *Operation
	[out]		classResult **Class
	[out, optional] moreResults *bool
	[out, optional]	result *Result
	[out, optional]	errorMessage **uint16
	[out, optional] completionDetails **Instance
*/
func (o *Operation) GetClass(moreResults *bool, result *Result, errorMessage *uint16, completionDetails **Instance) (*Class, Result) {
	classResult := &Class{}

	r0, _, _ := syscall.SyscallN(o.ft.getClass,
		uintptr(unsafe.Pointer(o)),
		uintptr(unsafe.Pointer(&classResult)),
		uintptr(unsafe.Pointer(moreResults)),
		uintptr(unsafe.Pointer(result)),
		uintptr(unsafe.Pointer(errorMessage)),
		uintptr(unsafe.Pointer(completionDetails)),
	)

	return classResult, Result(r0)
}
