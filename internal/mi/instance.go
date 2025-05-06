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

func (instance *MI_Instance) Clone() {
	panic("not implemented")
}

func (instance *MI_Instance) Destruct() {
	panic("not implemented")
}

func (instance *MI_Instance) Delete() {
	panic("not implemented")
}

func (instance *MI_Instance) IsA() {
	panic("not implemented")
}

func (instance *MI_Instance) GetClassName() {
	panic("not implemented")
}

func (instance *MI_Instance) SetNameSpace() {
	panic("not implemented")
}

func (instance *MI_Instance) GetNameSpace() {
	panic("not implemented")
}

func (instance *MI_Instance) GetElementCount(count *uint32) uint64 {

	r0, _, _ := syscall.SyscallN(instance.ft.GetElementCount,
		uintptr(unsafe.Pointer(instance)),
		uintptr(unsafe.Pointer(count)),
	)

	return uint64(r0)
}

func (instance *MI_Instance) AddElement() {
	panic("not implemented")
}

func (instance *MI_Instance) SetElement() {
	panic("not implemented")
}

func (instance *MI_Instance) SetElementAt() {
	panic("not implemented")
}

func (instance *MI_Instance) GetElement(name string, value *MI_Value) uint64 {

	// var value uintptr

	n, _ := syscall.UTF16PtrFromString(name)
	// n, err := syscall.UTF16FromString(name)

	// if err != nil {
	// 	panic("error on utf16 conversion")
	// }

	r0, _, _ := syscall.SyscallN(instance.ft.GetElement,
		uintptr(unsafe.Pointer(instance)),
		uintptr(unsafe.Pointer(n)),
		uintptr(unsafe.Pointer(value)),
		0,
		0,
		0,
	)

	return uint64(r0)
}

func (instance *MI_Instance) GetElementAt() {
	panic("not implemented")
}

func (instance *MI_Instance) ClearElement() {
	panic("not implemented")
}

func (instance *MI_Instance) ClearElementAt() {
	panic("not implemented")
}

func (instance *MI_Instance) GetServerName() {
	panic("not implemented")
}

func (instance *MI_Instance) SetServerName() {
	panic("not implemented")
}

func (instance *MI_Instance) GetClass() {
	panic("not implemented")
}
