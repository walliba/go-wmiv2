package mi

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type ptrdiff_t = int64

var (
	modmi                       = windows.NewLazySystemDLL("mi.dll")
	procMIApplicationInitialize = modmi.NewProc("MI_Application_InitializeV1")
)

// FeatureDecl is a base type for [PropertyDecl], [ParameterDecl], [MethodDecl]
type FeatureDecl struct {
	flags         uint32
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32
}

// [ObjectDecl] is a base type for [ParameterDecl] and [PropertyDecl],
// which allows functions to be written that work on the common fields of these two types.
type ObjectDecl struct {
	/* Fields from [FeatureDecl] */

	flags         uint32
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32

	/* Fields of this object */

	properties    **PropertyDecl
	numProperties uint32
	size          uint32
}

// PropertyDecl represents a CIM property (or reference)
type PropertyDecl struct {
	/* Fields from [FeatureDecl] */

	flags         uint32
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32

	/* Fields from [ParameterDecl] */

	paramType Type
	// className is the name of the reference class
	className *uint16
	// subscript is the "array subscript"
	subscript uint32
	// offset is the offset of this field within the structure
	offset uint32

	/* Fields of this object */

	// ancestor class that first defined a property with this name
	origin *uint16

	// ancestor class that last defined a property with this name
	propagator *uint16

	// value is an [unsafe.Pointer] to the underlying value of this property
	value unsafe.Pointer
}

// ParameterDecl represents a CIM property (or reference)
type ParameterDecl struct {
	/* Fields from [FeatureDecl] */

	flags         uint32
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32

	/* Fields of this object */

	// fieldType is the type of this field
	paramType Type
	// className is the name of the reference class
	className *uint16
	// subscript is the "array subscript"
	subscript uint32
	// offset is the offset of this field within the structure
	offset uint32
}

// MethodDecl represents a CIM method
type MethodDecl struct {
	/* Fields from [FeatureDecl] */

	flags         uint32
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32

	/* Fields from [ObjectDecl] */

	properties    **PropertyDecl
	numProperties uint32
	size          uint32

	/* Fields of this object */

	returnType Type // Type?
	origin     *uint16
	propagator *uint16
	schema     *SchemaDecl
	function   unsafe.Pointer // MethodDecl_Invoke
}

type ClassDecl struct {
	/* Fields from [FeatureDecl] */

	flags         uint32
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32

	/* Fields from [ObjectDecl] */

	properties    **PropertyDecl
	numProperties uint32
	size          uint32

	/* Fields of this object */

	superClass     *uint16      // *uint16
	superClassDecl *ClassDecl   // *ClassDecl
	methods        **MethodDecl // *MethodDecl
	numMethods     uint32
	schema         *SchemaDecl // *SchemaDecl
	ft             uintptr     // *ProviderFT
	owningClass    *Class      // *Class
}

type SchemaDecl struct {
	qualifierDecls    **QualifierDecl
	numQualifierDecls uint32
	classDecls        **ClassDecl
	numClassDecls     uint32
}

// QualifierDecl represents a CIM qualifier declaration
type QualifierDecl struct {
	// name is the name of this qualifier (UTF-16LE string)
	name *uint16

	// type is the type of this qualifier
	_type uint32

	// scope is the scope of this qualifier
	scope uint32

	// flavor is the flavor of this qualifier
	flavor uint32

	// subscript is the array subscript (for arrays only)
	// this is from mi.h; i don't know what this means
	subscript uint32

	// value is a pointer to an arbitrary type
	value unsafe.Pointer // void*
}

// Qualifier represents a CIM qualifier
type Qualifier struct {
	// name is the name of this qualifier (UTF-16LE string)
	name *uint16
	// type is the type of this qualifier
	_type Type
	// flavor is the flavor of this qualifier
	flavor uint32
	// value is a pointer to an arbitrary type
	value unsafe.Pointer // void*
}
