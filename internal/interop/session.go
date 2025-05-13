package interop

import (
	"fmt"
	"os"

	"github.com/walliba/go-wmiv2/internal/mi"
	"github.com/walliba/go-wmiv2/internal/mi/util"
)

type Session interface {
	Query(namespace string, query string) *[]map[string]any
}

type miSession struct {
	raw *mi.Session
}

func (s *miSession) Query(namespace string, query string) *[]map[string]any {
	operation := s.raw.QueryInstances(namespace, query)

	defer operation.Close()

	result := make([]map[string]any, 0)

	for moreResults := true; moreResults; {
		instance, err := operation.GetInstance(&moreResults)

		if err != mi.RESULT_OK {
			fmt.Println("failed on operation->GetInstance")
			continue
		}

		if instance != nil {
			instanceMap := make(map[string]any)

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

			result = append(result, instanceMap)
		}
	}

	return &result
}

func NewSession() Session {
	initApp()
	session, err := app.NewSession(nil, nil)

	if err != mi.RESULT_OK {
		fmt.Println("error: creating session")
		return nil
	}

	return &miSession{raw: session}
}
