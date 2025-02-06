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
			//获得读锁(共享锁)
			lock.RLock()
			//确保最终释放读锁(共享锁)
			defer lock.RUnlock()
			for i :=1; i<= 5; i++ {
				time.Sleep(time.Millisecond)
				//模拟耗时操作
				fmt.Println(i *10)
			}
	}()

	time.Sleep(time.Second*2)
}
