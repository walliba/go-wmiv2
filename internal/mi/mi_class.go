package mi

import (
	"syscall"
	"unsafe"

	"github.com/walliba/go-wmiv2/internal/mi/util"
)

type Class struct {
	ft            *classFT
	classDecl     *ClassDecl
	namespaceName *uint16
	serverName    *uint16
	reserved      [4]int64
}

type classFT struct {
	getClassName         uintptr
	getNameSpace         uintptr
	getServerName        uintptr
	getElementCount      uintptr
	getElement           uintptr
	getElementAt         uintptr
	getClassQualifierSet uintptr
	getMethodCount       uintptr
	getMethodAt          uintptr
	getMethod            uintptr
	getParentClassName   uintptr
	getParentClass       uintptr
	delete               uintptr
	clone                uintptr
}

func (c *Class) GetClassName() (string, Result) {
	var w_className unsafe.Pointer

	r0, _, _ := syscall.SyscallN(c.ft.getClassName,
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(&w_className)),
	)

	className := util.UTF16PtrToString((*uint16)(w_className))

	return className, Result(r0)
}

func (c *Class) GetNameSpace() (string, Result) {
	var w_namespace *uint16

	r0, _, _ := syscall.SyscallN(c.ft.getNameSpace,
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(&w_namespace)),
	)

	namespace := util.UTF16PtrToString(w_namespace)

	return namespace, Result(r0)
}

func (c *Class) GetServerName() (string, Result) {
	var w_serverName *uint16

	r0, _, _ := syscall.SyscallN(c.ft.getServerName,
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(&w_serverName)),
	)

	serverName := util.UTF16PtrToString(w_serverName)

	return serverName, Result(r0)
}

func (c *Class) GetElementCount() (*uint32, Result) {
	count := new(uint32)

	r0, _, _ := syscall.SyscallN(c.ft.getElementCount,
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(count)),
	)

	return count, Result(r0)
}

func (c *Class) GetElement() {
	panic("not implemented")
}

/*
GetElementAt signature

	[in] 		const MI_Class* self,
					MI_Uint32 index,
	[out, null] 	const MI_Char** name,
	[out, opt]      MI_Value* value,
	[out, opt]      MI_Boolean* valueExists,
	[out, opt]      MI_Type* type,
	[out, null]  	MI_Char **referenceClass,
	[out, opt]  	MI_QualifierSet *qualifierSet,
	[out, opt]      MI_Uint32* flags
*/
func (c *Class) GetElementAt(index uint32, value *Value, valueType *Type, flags *Flag) (string, Result) {

	var w_name *uint16

	r0, _, _ := syscall.SyscallN(c.ft.getElementAt,
		uintptr(unsafe.Pointer(c)),
		uintptr(index),
		uintptr(unsafe.Pointer(&w_name)),
		uintptr(unsafe.Pointer(value)),
		0,
		uintptr(unsafe.Pointer(valueType)),
		0,
		0,
		uintptr(unsafe.Pointer(flags)),
	)

	name := util.UTF16PtrToString(w_name)

	return name, Result(r0)
}
