package main

import "time"
import "fmt"
import "runtime"

func handle() {
	fmt.Println("start go routine")
	fmt.Println("end go routine")
}

func main() {
	go handle()
	time.Sleep(time.Second)

	fmt.Println("CPU核心数: ",  runtime.NumCPU())
	fmt.Println("GOMAXPROCS 默认值: ",  runtime.GOMAXPROCS(0))

}
