package main

import (
	"fmt"
	"os"
)

// system_info_collector - Collect system info
func system_info_collector(path string) {
	fmt.Println("========================================")
	fmt.Println("  System-Info-Collector")
	fmt.Println("  Collect system info")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	system_info_collector(path)
}
