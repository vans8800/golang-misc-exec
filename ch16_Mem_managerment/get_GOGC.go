package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	fmt.Println("GOGC = ", debug.SetGCPercent(200))
}
