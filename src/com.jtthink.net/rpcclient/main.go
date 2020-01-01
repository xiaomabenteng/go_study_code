package main

import (
	prod "com.jtthink.net/pbfiles"
	"fmt"
	"net/rpc"
)

func main()  {

	client,_:=rpc.Dial("tcp","127.0.0.1:8082")

	ret:=prod.ProdResponse{}
	client.Call("ProdService.GetStock",prod.ProdRequest{ProdId:600},&ret)
	fmt.Println(ret.ProdStock)
}
