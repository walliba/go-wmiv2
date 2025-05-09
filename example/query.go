package main

import (
	"fmt"
	"os"

	"github.com/walliba/go-wmiv2"
)

func main() {
	client := wmiv2.GetClient()

	defer client.Close()

	result := client.Query("select name, processid from win32_process")

	for i := range result {
		fmt.Printf("Process '%s' with PID %d\n", result[i]["Name"], result[i]["ProcessId"])
	}

	os.Exit(0)
}
