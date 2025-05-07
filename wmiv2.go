package wmiv2

import (
	"fmt"
	"os"
	"sync"

	"github.com/walliba/go-wmiv2/internal/mi"
	"golang.org/x/sys/windows"
)

type Client struct {
	IsInitialized bool
	app           *mi.Application
}

var instance *Client
var once sync.Once

func GetClient() *Client {
	once.Do(func() {
		app, err := mi.MI_Application_Initialize()

		if err != mi.RESULT_OK {
			panic("error initializing MI client")
		}

		if app == nil {
			panic("MI_Application instance is null")
		}

		instance = &Client{true, app}
	})

	return instance
}

func (c *Client) Close() {
	if err := c.app.Close(); err != mi.RESULT_OK {
		panic("Failed to close MI_Application handle")
	}
}

// .Description
// This will be (eventually) the abstraction layer to make using the API more go-like

func (c *Client) EnumerateAllInstances() {
	session, err := c.app.NewSession()

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

			fmt.Println(value.As(vType))

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

func (c *Client) Query(query string) {
	session, err := c.app.NewSession()

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

		if instance != nil {
			// MI_UInt32 index;
			// var index uint32
			var elementCount uint32

			err := instance.GetElementCount(&elementCount)

			if err != mi.RESULT_OK {
				fmt.Println("error getting element count")
			}

			var i uint32
			for i = 0; i < elementCount; i++ {

				// MI_Value value;
				value := &mi.Value{}
				// MI_Type type;
				var vType mi.Type
				// MI_Uint32 flags;
				var flags mi.Flag

				// fmt.Printf("&i: %p :: ", &i)
				name, err := instance.GetElementAt(&i, value, &vType, &flags)

				if flags.HasFlag(mi.FLAG_NULL) {
					continue
				}

				if err != mi.RESULT_OK {
					fmt.Printf("error %d: getting element at index: %d\n", err, i)
				}

				fmt.Fprintf(os.Stdout, "%s: %v\n", windows.UTF16PtrToString(name), value.As(vType))
			}
		}
	}

}
