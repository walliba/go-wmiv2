package mi

import "golang.org/x/sys/windows"

var (
	modmi                       = windows.NewLazySystemDLL("mi.dll")
	procMIApplicationInitialize = modmi.NewProc("MI_Application_InitializeV1")
)
