package mi

import (
	"syscall"
	"unsafe"
)

// _MI_Application
type MI_Application struct {
	reserved1 uint64
	reserved2 int64
	ft        *MI_ApplicationFT
}

type MI_ApplicationFT struct {
	Close                          uintptr
	NewSession                     uintptr
	NewHostedProvider              uintptr
	NewInstance                    uintptr
	NewDestinationOptions          uintptr
	NewOperationOptions            uintptr
	NewSubscriptionDeliveryOptions uintptr
	NewSerializer                  uintptr
	NewDeserializer                uintptr
	NewInstanceFromClass           uintptr
	NewClass                       uintptr
}

func (app *MI_Application) Close() Result {
	r0, _, _ := syscall.SyscallN(app.ft.Close, uintptr(unsafe.Pointer(app)))

	return Result(r0)
}

func (app *MI_Application) NewSession() (*MI_Session, Result) {
	session := &MI_Session{}

	r0, _, _ := syscall.SyscallN(app.ft.NewSession, uintptr(unsafe.Pointer(app)), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(session)))

	return session, Result(r0)
}

func (app *MI_Application) NewHostedProvider() {
	panic("not implemented")
}

func (app *MI_Application) NewInstance() {
	panic("not implemented")
}

func (app *MI_Application) NewDestinationOptions() {
	panic("not implemented")
}

// TODO: Convert to app.Initialize()
func MI_Application_Initialize() (*MI_Application, Result) {
	flags := uint32(0)

	application := &MI_Application{}

	r0, _, _ := procMIApplicationInitialize.Call(
		uintptr(flags), 0, 0, uintptr(unsafe.Pointer(application)),
	)

	return application, Result(r0)
}
