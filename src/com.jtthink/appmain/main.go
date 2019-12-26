package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)


func main()  {

	//文件读取操作
content:="hello abc"
ioutil.WriteFile("./files/1.txt",[]byte(content),666)
byt,_:=ioutil.ReadFile("./files/1.txt")

fmt.Println(string(byt))






}
