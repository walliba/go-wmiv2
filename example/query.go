package main

import (
	"fmt"
	"os"

	"github.com/walliba/go-wmiv2"
)

func main() {
	client := wmiv2.GetClient()

	defer client.Close()

	result := client.Query("select name, processid from win32_process where processid = 9880 or processid = 10904")

	fmt.Printf("Process '%s' with PID %d\n", result[0]["Name"], result[0]["ProcessId"])
	fmt.Printf("Process '%s' with PID %d\n", result[1]["Name"], result[1]["ProcessId"])

	os.Exit(0)
}
