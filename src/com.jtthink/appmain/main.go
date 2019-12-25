package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func sum(min,max int,c chan int)  {
	res:=0
	for i:=min;i<=max;i++ {
		res=res+i
	}
	c<-res
}
func main()  {

	start:=time.Now()
	num:=100000000
	c:=make(chan int,2)


	////sum(0,num,c) 同步执行
	//go sum(0,num,c) 协程执行，两者基本相同
	//res:=<-c //数据接收者

	go sum(0,num/2,c) //将数据分成两份，分别放入两个协程运行。总体运行时间减少为一半
	go sum(num/2+1,num,c)

	res1,res2:=<-c,<-c

	end:=time.Now()
	fmt.Println(end.Sub(start),res1+res2)




}
