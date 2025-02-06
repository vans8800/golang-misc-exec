package main

import(
	"fmt"
	"time"
	"sync"
)

func main() {
	//定义一个独占锁
	lock := sync.Mutex{}

	go func() {
		//锁定
		lock.Lock()

		//保证锁释放
		defer lock.Unlock()

		for i:= 1;i <= 5; i++ {
			//模拟耗时操作
			time.Sleep(time.Millisecond)
			//打印i的值
			fmt.Println(i)
		}
	}()


	go func(){
		//锁定
		lock.Lock()

		//保证锁释放
		defer lock.Unlock()

		for i:=1;i <= 5;i++ {
			time.Sleep(time.Millisecond)//拟耗时操作
			//打印i*10的值
			fmt.Println(i *10)
		}
	}()


	time.Sleep(time.Second*1)

}
