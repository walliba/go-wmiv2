package mi

import (
	"syscall"
	"unsafe"
)

type Session struct {
	reserved1 uint64
	reserved2 int64 // ptrdiff_t
	ft        *sessionFT
}

type sessionFT struct {
	close               uintptr
	getApplication      uintptr
	getInstance         uintptr
	modifyInstance      uintptr
	createInstance      uintptr
	deleteInstance      uintptr
	invoke              uintptr
	enumerateInstances  uintptr
	queryInstances      uintptr
	associatorInstances uintptr
	referenceInstances  uintptr
	subscribe           uintptr
	getClass            uintptr
	enumerateClasses    uintptr
	testConnection      uintptr
}

func (s *Session) Close() Result {
	r0, _, _ := syscall.SyscallN(s.ft.close, uintptr(unsafe.Pointer(s)), 0, uintptr(0))
	return Result(r0)
}

func (s *Session) GetApplication() { panic("not implemented") }

func (s *Session) GetInstance() { panic("not implemented") }

func (s *Session) ModifyInstance() { panic("not implemented") }

func (s *Session) CreateInstance() { panic("not implemented") }

func (s *Session) DeleteInstance() { panic("not implemented") }

func (s *Session) Invoke() { panic("not implemented") }

func (s *Session) EnumerateInstances(namespace string, class string) *Operation {
	// var namespace = "root\\cimv2"
	// var class = "Win32_Process"

	ns, _ := syscall.UTF16PtrFromString(namespace)
	c, _ := syscall.UTF16PtrFromString(class)

	operation := new(Operation)

	_, _, _ = syscall.SyscallN(s.ft.enumerateInstances,
		uintptr(unsafe.Pointer(s)),
		0,
		0,
		uintptr(unsafe.Pointer(ns)),
		uintptr(unsafe.Pointer(c)),
		0,
		0,
		uintptr(unsafe.Pointer(operation)),
	)

	return operation
}

func (s *Session) QueryInstances(namespace string, query string) *Operation {
	ns, _ := syscall.UTF16PtrFromString(namespace)
	d, _ := syscall.UTF16PtrFromString("WQL")
	q, _ := syscall.UTF16PtrFromString(query)

	operation := new(Operation)

	_, _, _ = syscall.SyscallN(s.ft.queryInstances,
		uintptr(unsafe.Pointer(s)),         // Session
		0,                                  // Flags
		uintptr(0),                         // Options
		uintptr(unsafe.Pointer(ns)),        // CIM Namespace
		uintptr(unsafe.Pointer(d)),         // Query dialect
		uintptr(unsafe.Pointer(q)),         // Query string
		0,                                  // Callbacks
		uintptr(unsafe.Pointer(operation)), // Operation
	)

	return operation

}

func (s *Session) AssociatorInstances() { panic("not implemented") }

func (s *Session) ReferenceInstances() { panic("not implemented") }

func (s *Session) Subscribe() { panic("not implemented") }

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

	_, _, _ = syscall.SyscallN(s.ft.getClass,
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

// Doc: https://learn.microsoft.com/en-us/windows/win32/api/mi/nf-mi-mi_session_enumerateclasses
func (s *Session) EnumerateClasses(namespace string, className string, classNamesOnly bool) *Operation {

	w_namespace, _ := syscall.UTF16PtrFromString(namespace)
	// Returns classes that derive from className. (e.g, CIM_Process returns Win32_Process)
	w_className, _ := syscall.UTF16PtrFromString(className)
	operation := new(Operation)

	_, _, _ = syscall.SyscallN(s.ft.enumerateClasses,
		uintptr(unsafe.Pointer(s)),
		0, // uintptr(flags),
		0, // uintptr(unsafe.Pointer(options)), // *OperationOptions
		uintptr(unsafe.Pointer(w_namespace)),
		uintptr(unsafe.Pointer(w_className)),
		uintptr(*(*uint8)(unsafe.Pointer(&classNamesOnly))), // cast bool to uint8 (unsigned char)
		0, // uintptr(unsafe.Pointer(callbacks)),
		uintptr(unsafe.Pointer(operation)),
	)

	return operation
}

func (s *Session) TestConnection() { panic("not implemented") }
