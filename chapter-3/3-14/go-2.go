package main

import (
	f "fmt"
	"runtime"
)

func main() {
	go f.Println("Yeah!")
	go f.Println("Yeah!")
	f.Printf("NumCPU: %d\n", runtime.NumCPU())
	f.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	f.Printf("Version: %d\n", runtime.Version())
}
