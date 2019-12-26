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

	c1:=make(chan string,len(users))
	c2:=make(chan string,len(ages))

	go func() {
		for _,v:=range users{
			c1<-v
		}
		close(c1)
	}()
	go func() {
		for _,v:=range ages{
			c2<-v
		}
		close(c2)
	}()

	for v:=range c1{ //range channel 只有v 没有index。channel可以使用range遍历，但是有区别在于它会一直遍历，直到channel被关闭，(关闭的channel，取值不会阻塞，返回零值)
		fmt.Println(v)
	}
	for v:=range c2{
		fmt.Println(v)
	}








}
