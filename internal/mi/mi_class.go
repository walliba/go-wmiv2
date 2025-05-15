package mi

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
