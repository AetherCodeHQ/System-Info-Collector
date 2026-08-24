package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("System Info Collector")
	fmt.Println("====================")
	fmt.Printf("OS:         %s\n", runtime.GOOS)
	fmt.Printf("Arch:       %s\n", runtime.GOARCH)
	fmt.Printf("CPUs:       %d\n", runtime.NumCPU())
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Time:       %s\n", time.Now().Format(time.RFC3339))
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Heap alloc: %d MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("Heap use:   %d MB\n", m.HeapInuse/1024/1024)
	fmt.Printf("Sys memory: %d MB\n", m.Sys/1024/1024)
	fmt.Printf("GC cycles:  %d\n", m.NumGC)
	exe, _ := os.Executable()
	fmt.Printf("Executable: %s\n", exe)
	envCount := len(os.Environ())
	fmt.Printf("Env vars:   %d\n", envCount)
}