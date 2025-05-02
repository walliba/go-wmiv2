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

func (session *MI_Session) Close() MI_RESULT {

	r0, _, _ := syscall.SyscallN(session.ft.Close, uintptr(unsafe.Pointer(session)), 0, uintptr(0))

	return MI_RESULT(r0)
}
