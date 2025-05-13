/*
interop is the interface boundary between the unsafe MI binding and the (eventually) safe public API.

The idea is a 3-layer acyclic architecture: mi -> interop -> wmiv2
*/
package interop

import (
	"sync"

	"github.com/walliba/go-wmiv2/internal/mi"
)

var (
	once sync.Once
	app  *mi.Application
)

func initApp() {
	once.Do(func() {
		inst, err := mi.MI_Application_Initialize()

		if err != mi.RESULT_OK {
			panic("FAILED TO INIT MI_APPLICATION")
		}

		app = inst
	})
}
