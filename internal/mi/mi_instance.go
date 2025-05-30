package mi

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/walliba/go-wmiv2/internal/mi/util"
)

type Instance struct {
	ft         *instanceFT
	classDecl  uintptr
	serverName uintptr
	nameSpace  uintptr
	reserved   [4]ptrdiff_t
}

type InstanceExFT struct {
	parent    instanceFT
	normalize uintptr
}

type instanceFT struct {
	clone           uintptr
	destruct        uintptr
	delete          uintptr
	isA             uintptr
	getClassName    uintptr
	setNameSpace    uintptr
	getNameSpace    uintptr
	getElementCount uintptr
	addElement      uintptr
	setElement      uintptr
	setElementAt    uintptr
	getElement      uintptr
	getElementAt    uintptr
	clearElement    uintptr
	clearElementAt  uintptr
	getServerName   uintptr
	setServerName   uintptr
	getClass        uintptr
}

func (i *Instance) String() string {
	className, err := i.GetClassName()

	if err != RESULT_OK {
		return "<class_instance>"
	}

	return className
}

func (i *Instance) isValid() bool {
	return i != nil && i.ft != nil
}

// Clone creates a copy of the given instance on the heap. Upon
// a successful return, newInstance points to a newly created instance. The
// new instance should eventually call [Delete].
func (i *Instance) Clone() (*Instance, Result) {
	if !i.isValid() {
		return nil, RESULT_INVALID_PARAMETER
	}

	var newInstance *Instance

	r0, _, _ := syscall.SyscallN(i.ft.clone,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&newInstance)),
	)

	// TODO: needs testing for stability
	runtime.AddCleanup(&newInstance, func(i *Instance) {
		if i.isValid() {
			i.Delete()
		}
	}, newInstance)

	return newInstance, Result(r0)
}

// Destruct deletes an instance that was created on the stack or as a member of a structure.
func (i *Instance) Destruct() Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	r0, _, _ := syscall.SyscallN(i.ft.destruct,
		uintptr(unsafe.Pointer(i)),
	)

	return Result(r0)
}

// Delete releases an instance that was created on the heap. Instances
// created with [Clone] should eventually be passed to Delete
func (i *Instance) Delete() Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	r0, _, _ := syscall.SyscallN(i.ft.delete,
		uintptr(unsafe.Pointer(i)),
	)

	return Result(r0)
}

func (i *Instance) IsA() {
	panic("not implemented")
}

func (i *Instance) GetClassName() (string, Result) {
	if !i.isValid() {
		return "", RESULT_INVALID_PARAMETER
	}

	var w_className *uint16

	r0, _, _ := syscall.SyscallN(i.ft.getClassName,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&w_className)),
	)

	className := util.UTF16PtrToString(w_className)

	return className, Result(r0)
}

func (i *Instance) SetNameSpace() {
	panic("not implemented")
}

func (i *Instance) GetNameSpace() (string, Result) {
	if !i.isValid() {
		return "", RESULT_INVALID_PARAMETER
	}

	var w_namespace *uint16

	r0, _, _ := syscall.SyscallN(i.ft.getNameSpace,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&w_namespace)),
	)

	namespace := util.UTF16PtrToString(w_namespace)

	return namespace, Result(r0)
}

func (i *Instance) GetElementCount(count *uint32) Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	r0, _, _ := syscall.SyscallN(i.ft.getElementCount,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(count)),
	)

	return Result(r0)
}

func (i *Instance) AddElement() {
	panic("not implemented")
}

func (i *Instance) SetElement() {
	panic("not implemented")
}

func (i *Instance) SetElementAt() {
	panic("not implemented")
}

// GetElement retrieves the value of the property with the given name.
func (i *Instance) GetElement(name string, v *Value, t *Type, f *Flag) Result {
	if !i.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	// Discard error because the syscall will return a RESULT_INVALID_PARAMETER anyway
	n, _ := syscall.UTF16PtrFromString(name)

	r0, _, _ := syscall.SyscallN(i.ft.getElement,
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(n)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(t)),
		uintptr(unsafe.Pointer(f)),
		0,
	)

	return Result(r0)
}

// GetElementAt retrieves the value of the property at the given index.
func (i *Instance) GetElementAt(index uint32, value *Value, valueType *Type, flags *Flag) (string, Result) {
	if !i.isValid() {
		return "", RESULT_INVALID_PARAMETER
	}

	// name := new(uint16)
	var w_name *uint16

	r0, _, _ := syscall.SyscallN(i.ft.getElementAt,
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&w_name)),
		uintptr(unsafe.Pointer(value)),
		uintptr(unsafe.Pointer(valueType)),
		uintptr(unsafe.Pointer(flags)),
	)

	name := util.UTF16PtrToString(w_name)

	return name, Result(r0)
}

func (i *Instance) ClearElement() {
	panic("not implemented")
}

func (i *Instance) ClearElementAt() {
	panic("not implemented")
}

// GetServerName retrieves the server name from the instance. The resultant name
// memory is owned by the instance and will be destroyed when
// the instance is deleted.
func (i *Instance) GetServerName() {
	panic("not implemented")
}

func (i *Instance) SetServerName() {
	panic("not implemented")
}

func (i *Instance) GetClass() {
	panic("not implemented")
}
