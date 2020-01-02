package main

import (
	prod "com.jtthink.net/pbfiles"
	"context"
	"google.golang.org/grpc"
	"net"
)
//protobuf 服务端
//type ProdService struct {}
//
//func (this *ProdService) GetStock(req prod.ProdRequest,res *prod.ProdResponse)  error{
//
//	if req.ProdId==500{
//		res.ProdStock=20
//	}else {
//		res.ProdStock=30
//	}
//	return nil
//}

type ProdServiceImpl struct {}

func (*ProdServiceImpl)  GetProdStock(ctx context.Context, req *prod.ProdRequest) (*prod.ProdResponse, error){
  	if req.ProdId==500{
  		return &prod.ProdResponse{ProdStock:500},nil
	}else {
		return &prod.ProdResponse{ProdStock:0},nil
	}
}

func main()  {


	lis,_:=net.Listen("tcp",":8082")
	//rpc.Register(new(ProdService)) //protobuf 服务端
	//for  {
	//	client,_:=lis.Accept()
	//	rpc.ServeConn(client)
	//}

	//grpc 服务端
	myserver:=grpc.NewServer()
	prod.RegisterProdServiceServer(myserver,new(ProdServiceImpl))
	myserver.Serve(lis)



}
