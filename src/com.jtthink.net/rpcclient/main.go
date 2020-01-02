package main

import (
	prod "com.jtthink.net/pbfiles"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main()  {

	//protobuf 客户端
	//client,_:=rpc.Dial("tcp","127.0.0.1:8082")
	//
	//ret:=prod.ProdResponse{}
	//client.Call("ProdService.GetStock",prod.ProdRequest{ProdId:500},&ret)
	//fmt.Println(ret.ProdStock)

	//grpc 客户端调用
	conn,err:=grpc.Dial("127.0.0.1:8082",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client:=prod.NewProdServiceClient(conn)
	ret,err:=client.GetProdStock(context.Background(),&prod.ProdRequest{ProdId:600})
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(ret.ProdStock)

}
