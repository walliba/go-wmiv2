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

	// var miOperation

	operation := session.EnumerateInstances("root\\cimv2", "Win32_Process")

	// if err != 0 {
	// 	panic(fmt.Sprintf("failed to enumerate session instances: %v", err))
	// }
	fmt.Printf("*MI_Operation: %#x\n", &operation)

	var moreResults bool

	for ok := true; ok; ok = moreResults {
		// var instance *mi.MI_Instance

		instance, err := operation.GetInstance(&moreResults)

		if err != 0 {
			fmt.Println("failed on operation->GetInstance")
			break
		}

		// fmt.Printf("*MI_Instance: %#x\n", &instance)

		if instance != nil {
			// MI_Value value;
			// MI_Type type;
			// MI_Uint32 flags;
			// var count uint32

			var value mi.MI_Value

			err := instance.GetElement("Name", &value)

			if err != 0 {
				return
			}

			fmt.Println(value.ToString())
			// err := instance.GetElementCount(&count)

			// if err != 0 {
			// 	fmt.Println("error getting count")
			// 	break
			// }

			// fmt.Println(count)

		}

		// value, err := instance.GetElement()

		// if err != 0 {
		// 	panic("failed on instance->GetElement")
		// }

		// fmt.Printf("*MI_Value: %#x\n", &value)

		// var s *string
		// s = (*string)(unsafe.Pointer(&value))

		// fmt.Println(s)
	}
}
