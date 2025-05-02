package mi

type MI_Operation struct {
	reserved1 uint64
	reserved2 int64
	ft        *MI_OperationFT
}

type MI_OperationFT struct {
	Close         uintptr
	Cancel        uintptr
	GetSession    uintptr
	GetInstance   uintptr
	GetIndication uintptr
	GetClass      uintptr
}
