package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个容量为1的channel
	ch := make(chan int, 1)

	//向channel中写入数据: 0
	ch <- 0

	go func() {
		//从channel中读取数据
		<-ch
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i)
		}

		//执行结束,向channel中写入数据: 0
		ch <- 0
	}()

	go func() {
		//从channel中读取数据
		<-ch
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i * 10)
		}
		//执行结束,向channel中写入数据: 0
		ch <- 0
	}()
	time.Sleep(time.Second)
}
