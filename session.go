package wmiv2

import (
	"fmt"

	"github.com/walliba/go-wmiv2/internal/mi"
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

		fmt.Println(class)
	}
}

func (s *miSession) Query(query string) []Instance {
	return nil
	operation := s.raw.QueryInstances("root\\cimv2", query)

	// This will hang if moreResults = true
	defer operation.Close()

	return nil
}
