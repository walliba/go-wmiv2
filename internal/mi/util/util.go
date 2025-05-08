package util

import "golang.org/x/sys/windows"

// Converts a []*uint16 to []string
func UTF16PtrsToStrings(ptrs []*uint16) []string {
	result := make([]string, len(ptrs))
	for i, p := range ptrs {
		if p != nil {
			result[i] = windows.UTF16PtrToString(p)
		}
	}
	return result
}
