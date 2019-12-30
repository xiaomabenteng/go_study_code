package main

import (
	"fmt"
	"net"
)

func main()  {

	lis,err:=net.Listen("tcp","127.0.0.1:8099")

	defer lis.Close()
	fmt.Println("创建监听成功，等待客户端连接")
	client,err:=lis.Accept()
	if err!=nil{
		fmt.Println(err)
		return
	}
	buf:=make([]byte,512)
	n,err:=client.Read(buf)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("读取到%d,内容是%s",n,string(buf))


}