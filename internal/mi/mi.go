package mi

type QualifierDecl struct {
	name      *uint16
	_type     Type
	scope     uint32
	flavor    uint32
	subscript uint32
	value     uintptr // void*
}

type Qualifier struct {
	name   *uint16
	_type  Type // uint32
	flavor uint32
	value  uintptr // void*
}

type PropertyDecl struct {
	flags         Flag // uint32
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32
	_type         Type // uint32
	className     *uint16
	subscript     uint32
	offset        uint32
	origin        *uint16
	propagator    *uint16
	value         uintptr // void*
}

type ParameterDecl struct {
	flags         Flag
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32
	_type         Type
	className     *uint16
	subscript     uint32
	offset        uint32
}

type SchemaDecl struct {
	qualifierDecls    **QualifierDecl
	numQualifierDecls uint32
	classDecls        **ClassDecl
	numClassDecls     uint32
}

type MethodDecl struct {
	flags         Flag
	code          uint32
	name          *uint16
	qualifiers    **Qualifier
	numQualifiers uint32
	parameters    **ParameterDecl
	numParameters uint32
	size          uint32
	returnType    uint32 // Type?
	origin        *uint16
	propagator    *uint16
	schema        *SchemaDecl
	function      uintptr // MethodDecl_Invoke
}
