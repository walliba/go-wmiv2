package main

import (
	"fmt"

	"github.com/walliba/go-wmiv2/wmiv2"
)

func main() {
	fmt.Println("Hello, world!")
	app, err := wmiv2.MI_Application_Initialize()

	if err != 0 {
		panic("failed to init")
	}

	// defer app.Close()

	if app == nil {
		panic("MI_Application is not initialized")
	}

	fmt.Printf("Close*: %#x\n", app.GetFt().Close)

	app.Close()

	fmt.Printf("Close*: %#x\n", app.GetFt().Close)

	// r0, _, _ := syscall.SyscallN(app.Ft.NewSession,
	// 	uintptr(unsafe.Pointer(app))
	// )

	// app.ft.Close()
}
