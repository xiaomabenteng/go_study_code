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

	c:=make(chan int,2) //带缓冲的channel。声明channel,第二个参数是容量。带有容量=2的 channel (好比队列)，只有当队列塞满时发送者会阻塞，队列空时接受者会阻塞
	sum(100,c) //
	res:=<-c //数据接收者
	fmt.Println(res)




}
