package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)


func main()  {

	users:=strings.Split("zhangsan,lisi,wangwu",",")
	ages:=strings.Split("19,21,22",",")

	c1,c2:=make(chan bool),make(chan bool)
	ret:=make([]string,0)

	go func() {
		for _,v:=range users{
			<-c1            //读取channel是空，阻塞
			ret=append(ret, v)
			c2<-true        //c2设置值，让下面不阻塞
		}
	}()
	go func() {
		for _,v:=range ages{
			<-c2            //c2有值，不阻塞
			ret=append(ret, v)
			c1<-true        //c1设置值，让c1不阻塞
		}
	}()

	c1<-true  //主进程运行到这里，设置c1,让c1协程不阻塞。
	fmt.Println(ret)
	//这样运行让c1和c2循环阻塞，可能交替打印出c1和c2内容。但是c1和c2里面运行时间长的话，主函数运行打印程序时，协程还未运行结束。此时仅能打印出部分数据








}
