package mi

type Class struct {
	ft            *ClassFT
	classDecl     uintptr
	namespaceName *uint16
	serverName    *uint16
	reserved      [4]int64
}

type ClassFT struct {
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

type ClassDecl struct {
	flags          Flag // uint32
	code           uint32
	name           uintptr // *uint16
	qualifiers     uintptr // *Qualifier
	numQualifiers  uint32
	properties     uintptr // *PropertyDecl
	numProperties  uint32
	size           uint32
	superClass     uintptr // *uint16
	superClassDecl uintptr // *ClassDecl
	methods        uintptr // *MethodDecl
	numMethods     uint32
	schema         uintptr // *SchemaDecl
	ft             uintptr // ProviderFT
	owningClass    uintptr // *Class
}
