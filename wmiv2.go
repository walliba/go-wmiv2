package wmiv2

import (
	"fmt"

	"github.com/walliba/go-wmiv2/internal/mi"
)

// .Description
// This will be (eventually) the abstraction layer to make using the API more go-like

func EnumerateAllInstances() {
	app, err := mi.MI_Application_Initialize()

	if err != 0 {
		panic("failed to init")
	}

	// defer app.Close()

	if app == nil {
		panic("MI_Application is not initialized")
	}

	fmt.Printf("*MI_Application: %#x\n", &app)

	defer func() {
		app.Close()
		fmt.Println("defer close app")
	}()

	session, err := app.NewSession()

	if err != 0 {
		panic(fmt.Sprintf("Failed on session creation, HRESULT = %d", err))
	}

	defer func() {
		session.Close()
		fmt.Println("defer close session")
	}()

	fmt.Printf("*MI_Session: %#x\n", &session)
}
