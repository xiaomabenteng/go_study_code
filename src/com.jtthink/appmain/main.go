package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func sum(max int)  {
	res:=0
	for i:=0;i<=max;i++ {
		res=res+i
	}
	fmt.Println(res)
}
func main()  {

	go sum(100) //使用go关键字开启协程
	time.Sleep(time.Second) //不加这一句时，因为main函数执行完时候，协程还未执行完时不会打印出结果

}
