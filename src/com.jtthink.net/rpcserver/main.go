package main

import (
	"net"
	"net/rpc"
)

type UserService struct {

}

func (*UserService) GetUserName (userid int,username *string)  error{

	if userid==101 {
		*username="zhangsan"
	}else{
		*username="guest"
	}
	return nil
}
func main()  {


	lis,_:=net.Listen("tcp",":8082")
	rpc.Register(new(UserService))

	for  {
		client,_:=lis.Accept()
		rpc.ServeConn(client)
	}



}
