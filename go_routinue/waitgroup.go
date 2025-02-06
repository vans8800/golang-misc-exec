package main

import(
	"fmt"
	"sync"
	"time"
)

func main(){
	//创建等待组对象
	wg := sync.WaitGroup{}
	//为等待组计数器加2
	wg.Add(2)
	go func(){
		//子协程中打印1~5数据段
		for i :=1;i <= 5;i++ {
			time.Sleep(time.Millisecond)
			//模拟耗时操作
			fmt .Println(i)
		}
		//调用等待组的Done()方法，令其计数器的值减1
		wg.Done()
	}()

	go func(){
		//子协程中打印10~50数据段
		for i :=1;i <= 5;i++ {
			time.Sleep(time.Millisecond)
			//模拟耗时操作
			fmt.Println(i*10)
		}
		//调用等待组的Done()方法，令其计数器的值减1
		wg.Done()
	}()
	//阻塞主协程的执行，直至计数器为0
	wg.Wait()
}
