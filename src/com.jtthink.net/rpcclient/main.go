package main

import (
	"fmt"
	"net/rpc"
)

func main()  {

	client,_:=rpc.Dial("tcp","127.0.0.1:8082")
	username:=""
	client.Call("UserService.GetUserName",1011,&username)
	fmt.Println(username)
}
