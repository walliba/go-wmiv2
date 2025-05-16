package main

import (
	"fmt"
	"os"

	"github.com/walliba/go-wmiv2"
)

func main() {
	// Get MI_Application instance
	app := wmiv2.GetApplication()

	// defer cleanup of the MI_Application instance
	defer app.Close()

	// query by namespace and WQL string
	// NOTE: result is temporarily a *[]map[string]any. Working on better solution
	result := app.Query("root\\cimv2", "select name, processid from win32_process")

	// print results
	for i := range *result {
		fmt.Printf("Process '%s' with PID %d\n", (*result)[i]["Name"], (*result)[i]["ProcessId"])
	}

	os.Exit(0)
}
