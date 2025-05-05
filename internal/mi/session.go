package mi

import (
	"syscall"
	"unsafe"
)

type MI_Session struct {
	reserved1 uint64
	reserved2 int64 // ptrdiff_t
	ft        *MI_SessionFT
}

type MI_SessionFT struct {
	Close               uintptr
	GetApplication      uintptr
	GetInstance         uintptr
	ModifyInstance      uintptr
	CreateInstance      uintptr
	DeleteInstance      uintptr
	Invoke              uintptr
	EnumerateInstances  uintptr
	QueryInstances      uintptr
	AssociatorInstances uintptr
	ReferenceInstaces   uintptr
	Subscribe           uintptr
	GetClass            uintptr
	EnumerateClasses    uintptr
	TestConnection      uintptr
}

func (session *MI_Session) Close() uint64 {

	r0, _, _ := syscall.SyscallN(session.ft.Close, uintptr(unsafe.Pointer(session)), 0, uintptr(0))

	return uint64(r0)
}

func (session *MI_Session) EnumerateInstances(namespace string, class string) *MI_Operation {
	// var namespace = "root\\cimv2"
	// var class = "Win32_Process"

	ns, _ := syscall.UTF16PtrFromString(namespace)
	c, _ := syscall.UTF16PtrFromString(class)

	var miOperation = MI_OPERATION_NULL

	_, _, _ = syscall.SyscallN(session.ft.EnumerateInstances,
		uintptr(unsafe.Pointer(session)),
		0,
		0,
		uintptr(unsafe.Pointer(ns)),
		uintptr(unsafe.Pointer(c)),
		0,
		0,
		uintptr(unsafe.Pointer(&miOperation)),
	)

	return &miOperation
}
