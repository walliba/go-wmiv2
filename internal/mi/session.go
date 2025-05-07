package mi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Session struct {
	reserved1 uint64
	reserved2 int64 // ptrdiff_t
	ft        *SessionFT
}

type SessionFT struct {
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

func (s *Session) Close() Result {
	r0, _, _ := syscall.SyscallN(s.ft.Close, uintptr(unsafe.Pointer(s)), 0, uintptr(0))
	return Result(r0)
}

func (s *Session) QueryInstances(namespace string, query string) *Operation {

	ns, _ := windows.UTF16PtrFromString(namespace)
	d, _ := syscall.UTF16PtrFromString("WQL")
	q, _ := syscall.UTF16PtrFromString(query)

	var miOperation = MI_OPERATION_NULL

	_, _, _ = syscall.SyscallN(s.ft.QueryInstances,
		uintptr(unsafe.Pointer(s)),            // Session
		0,                                     // Flags
		uintptr(0),                            // Options
		uintptr(unsafe.Pointer(ns)),           // CIM Namespace
		uintptr(unsafe.Pointer(d)),            // Query dialect
		uintptr(unsafe.Pointer(q)),            // Query string
		0,                                     // Callbacks
		uintptr(unsafe.Pointer(&miOperation)), // Operation
	)

	return &miOperation

}

func (s *Session) EnumerateInstances(namespace string, class string) *Operation {
	// var namespace = "root\\cimv2"
	// var class = "Win32_Process"

	ns, _ := syscall.UTF16PtrFromString(namespace)
	c, _ := syscall.UTF16PtrFromString(class)

	var miOperation = MI_OPERATION_NULL

	_, _, _ = syscall.SyscallN(s.ft.EnumerateInstances,
		uintptr(unsafe.Pointer(s)),
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
