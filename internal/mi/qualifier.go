package mi

import "unsafe"

// QualifierDecl represents a CIM qualifier declaration
type QualifierDecl struct {
	name      *uint16
	_type     uint32
	scope     uint32
	flavor    uint32
	subscript uint32
	value     unsafe.Pointer // void*
}

// Qualifier represents a CIM qualifier
type Qualifier struct {
	name   *uint16
	_type  uint32
	flavor uint32
	value  unsafe.Pointer // void*
}
