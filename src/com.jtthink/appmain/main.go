package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func sum(max int,c chan int)  {
	res:=0
	for i:=0;i<=max;i++ {
		res=res+i
	}
	c<-res //数据写入者
}
func main()  {

	c:=make(chan int) //声明channel
	go sum(100,c) // 通过channel实现主线程和协程交互
	res:=<-c //数据接收者
	fmt.Println(res)
	//上面的代码中
	//    对于发送者：如果没有接收者读取 channel （<- channel），则发送者 (channel <-) 会一直阻塞。
	//   同样对于接收者： 接收操作是阻塞的直到发送者发送数据
	//   根据上面的特性，就能实现当goroutine执行完成后  得到数据



}
