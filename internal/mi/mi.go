package mi

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modmi                       = windows.NewLazySystemDLL("mi.dll")
	procMIApplicationInitialize = modmi.NewProc("MI_Application_InitializeV1")
)

// FeatureDecl is a base type for [PropertyDecl], [ParameterDecl], [MethodDecl]
type FeatureDecl struct {
	flags         uint32
	code          uint32
	name          *uint16
	qualifiers    uintptr // **Qualifier?
	numQualifiers uint32
}

// PropertyDecl represents a CIM property (or reference)
type PropertyDecl struct {
	FeatureDecl
	_type      Type // uint32
	className  *uint16
	subscript  uint32
	offset     uint32
	origin     *uint16
	propagator *uint16
	value      uintptr // void*
}

// ParameterDecl represents a CIM property (or reference)
type ParameterDecl struct {
	FeatureDecl
	_type     Type
	className *uint16
	subscript uint32
	offset    uint32
}

// MethodDecl represents a CIM method
type MethodDecl struct {
	FeatureDecl
	parameters    uintptr // **ParameterDecl?
	numParameters uint32
	size          uint32
	returnType    uint32 // Type?
	origin        *uint16
	propagator    *uint16
	schema        *SchemaDecl
	function      unsafe.Pointer // MethodDecl_Invoke
}

type SchemaDecl struct {
	qualifierDecls    **QualifierDecl
	numQualifierDecls uint32
	classDecls        **ClassDecl
	numClassDecls     uint32
}
