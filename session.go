package wmiv2

import (
	"fmt"
	"os"

	"github.com/walliba/go-wmiv2/internal/mi"
	"github.com/walliba/go-wmiv2/internal/mi/util"
)

type miSession struct {
	raw *mi.Session
}

func (s *miSession) Close() error {
	result := s.raw.Close()

	if result != mi.RESULT_OK {
		return fmt.Errorf("error: failed to close: %d", result)
	}

	return nil
}

func (s *miSession) GetClass(namespace string, className string) {
	operation := s.raw.GetClass(namespace, className)

	defer operation.Close()

	for moreResults := true; moreResults; {
		class, result := operation.GetClass(&moreResults, nil, nil, nil)

		if result != mi.RESULT_OK {
			fmt.Println("error: operation.GetClass")
		}

		className, result := class.GetClassName()

		if result != mi.RESULT_OK {
			fmt.Println("error: GetClassName")
		}

		fmt.Printf("retrieved class name: %s\n", className)

		namespace, result := class.GetNameSpace()

		if result != mi.RESULT_OK {
			fmt.Println("error: GetNamespace")
		}

		fmt.Printf("retrieved namespace: %s\n", namespace)

		serverName, _ := class.GetServerName()

		fmt.Printf("retrieved server name: %s\n", serverName)

		elementCount, _ := class.GetElementCount()

		fmt.Printf("retrieved element count: %d\n", *elementCount)

		fmt.Println("Getting element keys:")
		for i := range *elementCount {
			valueType := new(mi.Type)
			flags := new(mi.Flag)
			name, _ := class.GetElementAt(i, nil, valueType, flags)
			fmt.Printf("\t%s: %v :: %s\n", name, *valueType, flags.GetFlags())
		}

	}
}

func (s *miSession) Query(namespace string, query string) *[]map[string]any {
	operation := s.raw.QueryInstances(namespace, query)

	// This will hang if moreResults = true
	defer operation.Close()

	// using an arbitrary initial size
	// TODO: look into inferring this size from MI
	result := make([]map[string]any, 8)

	for moreResults := true; moreResults; {

		instance, err := operation.GetInstance(&moreResults)
		if err != mi.RESULT_OK {
			fmt.Println("failed on operation->GetInstance")
			continue
		}

		if instance != nil {
			// instanceCount++
			var elementCount uint32
			err := instance.GetElementCount(&elementCount)

			instanceMap := make(map[string]any, elementCount)
			result = append(result, instanceMap)

			if err != mi.RESULT_OK {
				fmt.Println("error getting element count")
			}

			for i := range elementCount {
				// MI_Value value;
				value := new(mi.Value)
				// MI_Type type;
				vType := new(mi.Type)
				// MI_Uint32 flags;
				flags := new(mi.Flag)

				name, err := instance.GetElementAt(i, value, vType, flags)

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

	return &result
}

func (s *miSession) GetClasses(namespace string, classNamesOnly bool) {
	operation := s.raw.EnumerateClasses(namespace, classNamesOnly)

	// This will hang if moreResults = true
	defer operation.Close()

	// using an arbitrary initial size
	// TODO: look into inferring this size from MI
	// result := make([]string, 8)

	for moreResults := true; moreResults; {
		class, result := operation.GetClass(&moreResults, nil, nil, nil)

		if result != mi.RESULT_OK {
			fmt.Println("error: operation.GetClass")
		}

		className, result := class.GetClassName()

		if result != mi.RESULT_OK {
			fmt.Println("error: GetClassName")
		}

		fmt.Printf("retrieved class name: %s\n", className)
	}
}
