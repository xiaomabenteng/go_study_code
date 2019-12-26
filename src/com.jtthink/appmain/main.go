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

	alllist:=make([]string,0)

	//alllist=append(alllist, users...) //合并两个数组，到一个新的切片
	//alllist=append(alllist,ages...)
	////结果[zhangsan lisi wangwu 19 21 22]

	//for i,v:=range users{ //循环遍历交替插入
	//	alllist= append(alllist, v)
	//	alllist= append(alllist, ages[i])
	//}
	////结果[zhangsan 19 lisi 21 wangwu 22]

	fmt.Println(alllist)






}
