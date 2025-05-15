package mi

import (
	"syscall"
	"unsafe"
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
	ReferenceInstances  uintptr
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
	ns, _ := syscall.UTF16PtrFromString(namespace)
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

/*
GetClass signature

	[in]		session *Session
					flags Flag
	[in, optional]	options *OperationOptions
*/
func (s *Session) GetClass(namespaceName string, className string) *Operation {

	namespace, _ := syscall.UTF16PtrFromString(namespaceName)
	class, _ := syscall.UTF16PtrFromString(className)

	operation := new(Operation)

	_, _, _ = syscall.SyscallN(s.ft.GetClass,
		uintptr(unsafe.Pointer(s)),
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(namespace)),
		uintptr(unsafe.Pointer(class)),
		uintptr(0),
		uintptr(unsafe.Pointer(operation)),
	)

	return operation
}

func (s *Session) EnumerateClasses(namespace string, classNamesOnly bool) *Operation {

	w_namespace, _ := syscall.UTF16PtrFromString(namespace)
	// supplying className seems to panic, not sure how this is used
	// w_className, _ := syscall.UTF16PtrFromString(className)

	operation := new(Operation)

	_, _, _ = syscall.SyscallN(s.ft.EnumerateClasses,
		uintptr(unsafe.Pointer(s)),
		0,
		0,
		uintptr(unsafe.Pointer(w_namespace)),
		0, // uintptr(unsafe.Pointer(w_className)),
		uintptr(unsafe.Pointer(&classNamesOnly)),
		0,
		uintptr(unsafe.Pointer(operation)),
	)

	return operation
}
