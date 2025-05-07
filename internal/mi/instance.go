package mi

import (
	"syscall"
	"unsafe"
)

type MI_Instance struct {
	ft         *MI_InstanceFT
	classDecl  uintptr
	serverName uintptr
	nameSpace  uintptr
	reserved   [4]int64
}

type MI_InstanceExFT struct {
	parent    MI_InstanceFT
	Normalize uintptr
}

type MI_InstanceFT struct {
	Clone           uintptr
	Destruct        uintptr
	Delete          uintptr
	IsA             uintptr
	GetClassName    uintptr
	SetNameSpace    uintptr
	GetNameSpace    uintptr
	GetElementCount uintptr
	AddElement      uintptr
	SetElement      uintptr
	SetElementAt    uintptr
	GetElement      uintptr
	GetElementAt    uintptr
	ClearElement    uintptr
	ClearElementAt  uintptr
	GetServerName   uintptr
	SetServerName   uintptr
	GetClass        uintptr
}

func (i *MI_Instance) isValid() bool {
	return i != nil && i.ft != nil
}

func (i *MI_Instance) Clone() {
	panic("not implemented")
}

func (i *MI_Instance) Destruct() Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	r0, _, _ := syscall.SyscallN(i.ft.Destruct,
		uintptr(unsafe.Pointer(i)),
	)

	return Result(r0)
}

func (i *MI_Instance) Delete() Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	r0, _, _ := syscall.SyscallN(i.ft.Delete,
		uintptr(unsafe.Pointer(i)),
	)

	return Result(r0)
}

func (i *MI_Instance) IsA() {
	panic("not implemented")
}

func (i *MI_Instance) GetClassName() {
	panic("not implemented")
}

func (i *MI_Instance) SetNameSpace() {
	panic("not implemented")
}

func (i *MI_Instance) GetNameSpace() (*uint16, Result) {
	if !i.isValid() {
		return nil, RESULT_INVALID_PARAMETER
	}

	var namespace *uint16

	r0, _, _ := syscall.SyscallN(i.ft.GetNameSpace,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&namespace)),
	)

	return namespace, Result(r0)
}

func (i *MI_Instance) GetElementCount(count *uint32) Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	r0, _, _ := syscall.SyscallN(i.ft.GetElementCount,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(count)),
	)

	return Result(r0)
}

func (i *MI_Instance) AddElement() {
	panic("not implemented")
}

func (i *MI_Instance) SetElement() {
	panic("not implemented")
}

func (i *MI_Instance) SetElementAt() {
	panic("not implemented")
}

// Gets the value of the property with the given name.
func (i *MI_Instance) GetElement(name string, v *Value, t *Type, f *Flag) Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	// Discard error because the syscall will return a RESULT_INVALID_PARAMETER anyway
	n, _ := syscall.UTF16PtrFromString(name)

	r0, _, _ := syscall.SyscallN(i.ft.GetElement,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(n)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(t)),
		uintptr(unsafe.Pointer(f)),
		0,
	)

	return Result(r0)
}

// Get the value of the property at the given index.
func (i *MI_Instance) GetElementAt(idx *uint32, v *Value, t *Type, f *Flag) (*uint16, Result) {
	if !i.isValid() {
		return nil, RESULT_INVALID_PARAMETER
	}

	var name *uint16

	r0, _, _ := syscall.SyscallN(i.ft.GetElementAt,
		uintptr(unsafe.Pointer(i)),
		uintptr(*idx),
		uintptr(unsafe.Pointer(&name)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(t)),
		uintptr(unsafe.Pointer(f)),
	)

	return name, Result(r0)
}

func (i *MI_Instance) ClearElement() {
	panic("not implemented")
}

func (i *MI_Instance) ClearElementAt() {
	panic("not implemented")
}

func (i *MI_Instance) GetServerName() {
	panic("not implemented")
}

func (i *MI_Instance) SetServerName() {
	panic("not implemented")
}

func (i *MI_Instance) GetClass() {
	panic("not implemented")
}
