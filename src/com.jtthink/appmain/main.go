package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)


func main()  {

	var WG sync.WaitGroup
	var mutex sync.Mutex
	WG.Add(2)
	list:=make([]int,0)  //两个协程并行运行，会出现协程安全性问题，在一个协程还未执行完时，被切换到其他协程，导致数据出错。使用锁机制，解决
	go func() {
		defer WG.Done()
		mutex.Lock()      //设置锁
		for i:=1;i<=1000000;i++{
			list=append(list,1)
		}
		fmt.Println(len(list))
		mutex.Unlock() //协程执行完毕释放锁
	}()
	go func() {
		defer WG.Done()
		mutex.Lock()
		for i:=1;i<=1000000;i++{
			list=append(list,1)
		}
		fmt.Println(len(list))
		mutex.Unlock()
	}()

	WG.Wait()
	fmt.Println(len(list))


}
