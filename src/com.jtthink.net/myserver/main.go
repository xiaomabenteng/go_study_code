package main

import (
	"fmt"
	"net"
)

func main()  {

	lis,err:=net.Listen("tcp","127.0.0.1:8099")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer lis.Close()
	fmt.Println("创建监听成功，等待客户端连接")
	for{
		client,err:=lis.Accept()
		if err!=nil{
			fmt.Println(err)
			return
		}
		func (c net.Conn){
			defer c.Close()
			buf:=make([]byte,4096)
			n,err:=c.Read(buf)
			if err!=nil{
				fmt.Println(err)
				return
			}
			fmt.Printf(string(buf[:n]))
			c.Write([]byte("abc"))
		}(client)
	}





}