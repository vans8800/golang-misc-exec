package main

import(
	"fmt"
	"sync"
	"time"
)

func main(){
	//创建读写锁
	lock := sync.RWMutex{}
	go func(){
		//获得读锁(共享锁)
		lock.RLock()
		//确保最终释放读锁(共享锁)
		defer lock.RUnlock()

		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			//模拟耗时操作
			fmt .Println(i)
		}
	}()

	go func(){
			//针对写锁的操作将与读锁操作产生互斥效应，因此，两个协程仍将串行执行。

			//获得写锁(共享锁)
			lock.Lock()
			//确保最终释写锁(独占 锁)
			defer lock.Unlock()
			for i :=1; i<= 5; i++ {
				time.Sleep(time.Millisecond)
				//模拟耗时操作
				fmt.Println(i *10)
			}
	}()

	time.Sleep(time.Second*2)
}
