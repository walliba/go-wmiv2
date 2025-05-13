// wmiv2 is the public-facing package that encapsulates access to the underlying MI API.
package wmiv2

import (
	"fmt"
	"os"
	"sync"

	"github.com/walliba/go-wmiv2/internal/interop"
	"github.com/walliba/go-wmiv2/internal/mi"
	"github.com/walliba/go-wmiv2/internal/mi/util"
)

/*
Client should function as a singleton for the lifespan of the application

TODO: write code to support above statement
*/
type Client struct {
	IsInitialized bool
	app           *mi.Application
}

var instance *Client
var once sync.Once

func GetClient() *Client {
	once.Do(func() {
		fmt.Println("initializing client")
		app, err := mi.MI_Application_Initialize()

		if err != mi.RESULT_OK {
			panic("error initializing MI client")
		}

		if app == nil {
			panic("MI_Application instance is null")
		}

		instance = &Client{true, app}
	})
	fmt.Println("Getting client")
	return instance
}

func (c *Client) Close() {
	if err := c.app.Close(); err != mi.RESULT_OK {
		panic("Failed to close MI_Application handle")
	}
}

// func (c *Client) EnumerateAllInstances() {
// 	session, err := c.app.NewSession()

// 	if err != 0 {
// 		panic(fmt.Sprintf("Failed on session creation, HRESULT = %d", err))
// 	}

// 	defer func() {
// 		if err := session.Close(); err != mi.RESULT_OK {
// 			panic("Failed to close MI_Session handle")
// 		}
// 	}()

// 	// fmt.Printf("*MI_Session: %#x\n", &session)

// 	// var miOperation

// 	operation := session.EnumerateInstances("root\\cimv2", "Win32_Process")

// 	defer func() {
// 		if err := operation.Close(); err != mi.RESULT_OK {
// 			panic("Failed to close MI_Operation handle")
// 		}
// 	}()
// 	// if err != 0 {
// 	// 	panic(fmt.Sprintf("failed to enumerate session instances: %v", err))
// 	// }
// 	// fmt.Printf("*MI_Operation: %#x\n", &operation)

// 	// var moreResults bool = true

// 	var instanceCount uint32 = 0

// 	for moreResults := true; moreResults; {
// 		// var instance *mi.MI_Instance

// 		instance, err := operation.GetInstance(&moreResults)

// 		if err != 0 {
// 			fmt.Println("failed on operation->GetInstance")
// 			break
// 		}

// 		// fmt.Printf("*MI_Instance: %#x\n", &instance)

// 		if instance != nil {
// 			// MI_Value value;
// 			var value mi.Value
// 			// MI_Type type;
// 			var vType mi.Type
// 			// MI_Uint32 flags;
// 			var flags mi.Flag
// 			// MI_UInt32 index;
// 			// var index uint32

// 			err := instance.GetElement("Name", &value, &vType, &flags)

// 			if err != mi.RESULT_OK {
// 				return
// 			}

// 			fmt.Println(value.As(vType))

// 			// fmt.Println(flags.HasFlag(mi.FLAG_READONLY))
// 			// fmt.Println(flags.GetFlags())
// 			// fmt.Println(mi.ListSetFlags(flags))
// 			// err := instance.GetElementCount(&count)

// 			// if err != 0 {
// 			// 	fmt.Println("error getting count")
// 			// 	break
// 			// }

// 			// fmt.Println(count)

// 		}
// 		instanceCount++
// 		// value, err := instance.GetElement()

// 		// if err != 0 {
// 		// 	panic("failed on instance->GetElement")
// 		// }

// 		// fmt.Printf("*MI_Value: %#x\n", &value)

// 		// var s *string
// 		// s = (*string)(unsafe.Pointer(&value))

// 		// fmt.Println(s)
// 		if !moreResults {
// 			break
// 		}
// 	}

// 	fmt.Printf("Done: %d\n", instanceCount)
// }

func IQuery(query string) *[]map[string]any {
	session := interop.NewSession()

	return session.Query("root\\cimv2", query)
}

func (c *Client) Query(query string) []map[string]any {
	session, err := c.app.NewSession(nil, nil)

	if err != mi.RESULT_OK {
		panic(fmt.Sprintf("Failed on session creation, HRESULT = %d", err))
	}

	defer func() {
		fmt.Println("attempting to close session")
		if err := session.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Session handle")
		}
	}()

	operation := session.QueryInstances("root\\cimv2", query)

	defer func() {
		fmt.Println("attempting to close operation")
		if err := operation.Close(); err != mi.RESULT_OK {
			panic("Failed to close MI_Operation handle")
		}
	}()

	// TODO: use a concurrency-safe map slice or alternative
	result := make([]map[string]any, 0)
	instanceCount := 0
	for moreResults := true; moreResults; {

		instance, err := operation.GetInstance(&moreResults)

		if err != 0 {
			fmt.Println("failed on operation->GetInstance")
			continue
		}

		if instance != nil {
			instanceCount++
			instanceMap := make(map[string]any)
			result = append(result, instanceMap)

			var elementCount uint32

			err := instance.GetElementCount(&elementCount)

			if err != mi.RESULT_OK {
				fmt.Println("error getting element count")
			}

			var i uint32
			for i = 0; i < elementCount; i++ {
				// MI_Value value;
				value := new(mi.Value)
				// MI_Type type;
				vType := new(mi.Type)
				// MI_Uint32 flags;
				flags := new(mi.Flag)

				name, err := instance.GetElementAt(&i, value, vType, flags)

				if err != mi.RESULT_OK {
					fmt.Fprintf(os.Stderr, "error %d: getting element at index: %d\n", err, i)
					continue
				}

				if flags.HasFlag(mi.FLAG_NULL) {
					// Omitting this results in a smaller slice, and still allows for indexing into the result map (returns nil)
					// key := windows.UTF16PtrToString(name)
					// instanceMap[key] = nil
					continue
				}

				key := util.UTF16PtrToString(name)
				instanceMap[key] = value.As(*vType)
			}
		}
	}

	return result
}

func (c *Client) NewDestinationCredentials(username string, password string) {}
