package wmiv2

import "golang.org/x/sys/windows"

var (
	modmi = windows.NewLazySystemDLL("mi.dll")
)
