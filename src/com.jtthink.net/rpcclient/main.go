package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main()  {

	client,_:=jsonrpc.Dial("tcp","127.0.0.1:8082")
	username:=""
	client.Call("UserService.GetUserName",101,&username)
	fmt.Println(username)
}
