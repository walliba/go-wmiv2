package mi

import (
	"syscall"
	"unsafe"
)

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

// func (app *MI_Application) Initialize() MI_RESULT {

// 	r0, _, _ := procMIApplicationInitialize.Call(
// 		uintptr(0), 0, 0, uintptr(unsafe.Pointer(&app)),
// 	)

// 	return MI_RESULT(r0)
// }

// func MI_Application_NewSession()

// func (mi *MI) Initialize() MI_RESULT {
// 	app, hresult := MI_Application_Initialize()
// }

// DEBUG: REMOVE BEFORE RELEASE
func (app *MI_Application) GetFt() *MI_ApplicationFT {
	return app.ft
}

// TODO: Convert to app.Initialize()
func MI_Application_Initialize() (*MI_Application, uint64) {
	flags := uint32(0)

	application := &MI_Application{}

	r0, _, _ := procMIApplicationInitialize.Call(
		uintptr(flags), 0, 0, uintptr(unsafe.Pointer(application)),
	)

	return application, uint64(r0)
}

func (app *MI_Application) Close() uint64 {
	if app.ft == nil || app.ft.Close == 0 {
		return 1
	}

	r0, _, _ := syscall.SyscallN(app.ft.Close, uintptr(unsafe.Pointer(app)))

	return uint64(r0)
}

func (app *MI_Application) NewSession() (*MI_Session, uint64) {
	if app.ft == nil || app.ft.Close == 0 {
		return nil, 1
	}

	session := &MI_Session{}

	r0, _, _ := syscall.SyscallN(app.ft.NewSession, uintptr(unsafe.Pointer(app)), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(session)))

	return session, uint64(r0)
}
