package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	 currentTime := time.Now().Format("2006-01-02 15:04:05")
	 for i, arg :=  range os.Args {
		fmt.Printf("参数%d = %s\n", i, arg)
	 }
	 fmt.Println(currentTime)
}

