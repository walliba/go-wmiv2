package wmiv2

import (
	"fmt"

	"golang.org/x/sys/windows"

	"github.com/walliba/go-wmiv2/internal/mi"
)

// .Description
// This will be (eventually) the abstraction layer to make using the API more go-like

func EnumerateAllInstances() {
	app, err := mi.MI_Application_Initialize()

	if err != 0 {
		panic("failed to init")
	}

	if app == nil {
		panic("MI_Application is not initialized")
	}

	// fmt.Printf("*MI_Application: %#x\n", &app)

	defer func() {
		if err := app.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Application handle")
		}
	}()

	session, err := app.NewSession()

	if err != 0 {
		panic(fmt.Sprintf("Failed on session creation, HRESULT = %d", err))
	}

	defer func() {
		if err := session.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Session handle")
		}
	}()

	// fmt.Printf("*MI_Session: %#x\n", &session)

	// var miOperation

	operation := session.EnumerateInstances("root\\cimv2", "Win32_Process")

	defer func() {
		if err := operation.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Operation handle")
		}
	}()
	// if err != 0 {
	// 	panic(fmt.Sprintf("failed to enumerate session instances: %v", err))
	// }
	// fmt.Printf("*MI_Operation: %#x\n", &operation)

	// var moreResults bool = true

	var instanceCount uint32 = 0

	for moreResults := true; moreResults; {
		// var instance *mi.MI_Instance

		instance, err := operation.GetInstance(&moreResults)

		if err != 0 {
			fmt.Println("failed on operation->GetInstance")
			break
		}

		// fmt.Printf("*MI_Instance: %#x\n", &instance)

		if instance != nil {
			// MI_Value value;
			var value mi.Value
			// MI_Type type;
			var vType mi.Type
			// MI_Uint32 flags;
			var flags mi.Flag
			// MI_UInt32 index;
			// var index uint32

			err := instance.GetElement("Name", &value, &vType, &flags)

			if err != mi.RESULT_OK {
				return
			}

			fmt.Println(value.As(&vType))

			// fmt.Println(flags.HasFlag(mi.FLAG_READONLY))
			// fmt.Println(flags.GetFlags())
			// fmt.Println(mi.ListSetFlags(flags))
			// err := instance.GetElementCount(&count)

			// if err != 0 {
			// 	fmt.Println("error getting count")
			// 	break
			// }

			// fmt.Println(count)

		}
		instanceCount++
		// value, err := instance.GetElement()

		// if err != 0 {
		// 	panic("failed on instance->GetElement")
		// }

		// fmt.Printf("*MI_Value: %#x\n", &value)

		// var s *string
		// s = (*string)(unsafe.Pointer(&value))

		// fmt.Println(s)
		if !moreResults {
			break
		}
	}

	fmt.Printf("Done: %d\n", instanceCount)
}

func Query(query string) {
	app, err := mi.MI_Application_Initialize()

	if err != 0 {
		panic("failed to init")
	}

	if app == nil {
		panic("MI_Application is not initialized")
	}

	defer func() {
		if err := app.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Application handle")
		}
	}()

	session, err := app.NewSession()

	if err != 0 {
		panic(fmt.Sprintf("Failed on session creation, HRESULT = %d", err))
	}

	defer func() {
		if err := session.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Session handle")
		}
	}()

	operation := session.QueryInstances("root\\cimv2", query)

	defer func() {
		if err := operation.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Operation handle")
		}
	}()

	for moreResults := true; moreResults; {

		instance, err := operation.GetInstance(&moreResults)

		if err != 0 {
			fmt.Println("failed on operation->GetInstance")
			continue
		}

		// namespace, err := instance.GetNameSpace()

		// if err != mi.RESULT_OK {
		// 	fmt.Println("failed on instance->GetNameSpace")
		// 	break
		// }

		// fmt.Println(windows.UTF16PtrToString(namespace))

		if instance != nil {
			// MI_UInt32 index;
			// var index uint32
			var eCount uint32

			err := instance.GetElementCount(&eCount)

			if err != mi.RESULT_OK {
				fmt.Println("error getting element count")
			}

			for i := uint32(0); i < eCount; i++ {
				// MI_Value value;
				var value mi.Value
				// MI_Type type;
				var vType mi.Type
				// MI_Uint32 flags;
				var flags mi.Flag

				// BUG: flip flopping i index
				name, err := instance.GetElementAt(i, &value, &vType, &flags)

				if flags&mi.FLAG_NULL != 0 {
					continue
				}

				if err != mi.RESULT_OK {
					fmt.Printf("error %d: getting element at index: %d\n", err, i)
				}

				fmt.Println(windows.UTF16PtrToString(name))

				// fmt.Printf("%v", value.As(&vType))

			}
		}
	}

}
