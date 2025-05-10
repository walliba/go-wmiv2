package mi

import "unsafe"

type SessionCallbacks struct {
	callbackContext unsafe.Pointer // void*
	writeMessage    uintptr
	writeError      uintptr
}

func (sc *SessionCallbacks) WriteMessage() { panic("not implemented") }

func (sc *SessionCallbacks) WriteError() { panic("not implemented") }
