package mi

import (
	"syscall"
	"unsafe"
)

// _MI_Application
type Application struct {
	reserved1 uint64
	reserved2 int64
	ft        *ApplicationFT
}

type ApplicationFT struct {
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

func (app *Application) Close() Result {
	r0, _, _ := syscall.SyscallN(app.ft.Close, uintptr(unsafe.Pointer(app)))

	return Result(r0)
}

func (app *Application) NewSession() (*Session, Result) {
	session := &Session{}

	r0, _, _ := syscall.SyscallN(app.ft.NewSession,
		uintptr(unsafe.Pointer(app)),     // self *Application
		uintptr(0),                       // protocol *uint16
		uintptr(0),                       // destination *uint16
		uintptr(0),                       // options *DestinationOptions
		uintptr(0),                       // callbacks *SessionCallbacks
		uintptr(0),                       // extendedError **Instance
		uintptr(unsafe.Pointer(session)), // session *Session
	)

	return session, Result(r0)
}

func (app *Application) NewHostedProvider() {
	panic("not implemented")
}

func (app *Application) NewInstance() {
	panic("not implemented")
}

func (app *Application) NewDestinationOptions() {
	panic("not implemented")
}

func (app *Application) NewOperationOptions() {
	panic("not implemented")
}

func (app *Application) NewSubscriptionDeliveryOptions() {
	panic("not implemented")
}

func (app *Application) NewSerializer() {
	panic("not implemented")
}

func (app *Application) NewDeserializer() {
	panic("not implemented")
}

func (app *Application) NewInstanceFromClass() {
	panic("not implemented")
}

func (app *Application) NewClass() {
	panic("not implemented")
}

// TODO: Convert to app.Initialize()
func MI_Application_Initialize() (*Application, Result) {
	flags := uint32(0)

	application := &Application{}

	r0, _, _ := procMIApplicationInitialize.Call(
		uintptr(flags), 0, 0, uintptr(unsafe.Pointer(application)),
	)

	return application, Result(r0)
}
