package mi

import (
	"syscall"
	"unsafe"
)

// _MI_Application
type Application struct {
	reserved1 uint64
	reserved2 int64
	ft        *applicationFT
}

type applicationFT struct {
	close                          uintptr
	newSession                     uintptr
	newHostedProvider              uintptr
	newInstance                    uintptr
	newDestinationOptions          uintptr
	newOperationOptions            uintptr
	newSubscriptionDeliveryOptions uintptr
	newSerializer                  uintptr
	newDeserializer                uintptr
	newInstanceFromClass           uintptr
	newClass                       uintptr
}

func (app *Application) Close() Result {
	r0, _, _ := syscall.SyscallN(app.ft.close, uintptr(unsafe.Pointer(app)))

	return Result(r0)
}

func (app *Application) NewSession(destination *uint16, options *DestinationOptions) (*Session, Result) {
	// TODO: verify if this allows GC to free?
	session := &Session{}

	r0, _, _ := syscall.SyscallN(app.ft.newSession,
		uintptr(unsafe.Pointer(app)),         // self *Application
		uintptr(0),                           // protocol *uint16
		uintptr(unsafe.Pointer(destination)), // destination *uint16
		uintptr(unsafe.Pointer(options)),     // options *DestinationOptions
		uintptr(0),                           // callbacks *SessionCallbacks
		uintptr(0),                           // extendedError **Instance
		uintptr(unsafe.Pointer(session)),     // session *Session
	)

	return session, Result(r0)
}

func (app *Application) NewHostedProvider() {
	panic("not implemented")
}

func (app *Application) NewInstance() {
	panic("not implemented")
}

func (app *Application) NewDestinationOptions() (*DestinationOptions, Result) {
	options := &DestinationOptions{}

	r0, _, _ := syscall.SyscallN(app.ft.newDestinationOptions,
		uintptr(unsafe.Pointer(app)),     // self *Application
		uintptr(unsafe.Pointer(options)), // options *DestinationOptions
	)

	return options, Result(r0)
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

// Initializes the *mi.Application passed
func MI_Application_Initialize(application *Application) Result {

	r0, _, _ := procMIApplicationInitialize.Call(
		0,                                    // flags uint32
		0,                                    // applicationID *uint16
		0,                                    // extendedError **Instance
		uintptr(unsafe.Pointer(application)), // application *Application
	)

	return Result(r0)
}
