package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.GOMAXPROCS(1)

	go func() {
		fmt.Println("Child go corountinue")
	}()

	runtime.Gosched()
}
