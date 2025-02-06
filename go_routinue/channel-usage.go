package main

import (
	"fmt"
	"time"
)

// 使用channel实现两个goroutine的交替执行, 实现协程同步--互斥
// func main() {
// 	//创建一个容量为1的channel
// 	ch := make(chan int, 1)

// 	//向channel中写入数据: 0
// 	ch <- 0

// 	go func() {
// 		//从channel中读取数据
// 		<-ch
// 		for i := 1; i <= 5; i++ {
// 			time.Sleep(time.Millisecond)
// 			fmt.Println(i)
// 		}

// 		//执行结束,向channel中写入数据: 0
// 		ch <- 0
// 	}()

// 	go func() {
// 		//从channel中读取数据
// 		<-ch
// 		for i := 1; i <= 5; i++ {
// 			time.Sleep(time.Millisecond)
// 			fmt.Println(i * 10)
// 		}
// 		//执行结束,向channel中写入数据: 0
// 		ch <- 0
// 	}()
// 	time.Sleep(time.Second)
// }

//采用channel实现等待组的功能
func main() {
	count := 2
	//创建一个无缓冲的chanel
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond) //模拟耗时操作
			fmt.Println(i)
		}
		//子协程执行结束,向channel中写入数据: 0
		ch <- 0
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond) //模拟耗时操作
			fmt.Println(i * 10)
		}
		//子协程执行结束,向channel中写入数据: 0
		ch <- 0
	}()

	for i := 0; i < count; i++ {
		//主协程中,从channel中读取数据
		<-ch
	}

	fmt.Println("main goroutine exit")
}
