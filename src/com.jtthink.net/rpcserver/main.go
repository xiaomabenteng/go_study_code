package main

import (
	prod "com.jtthink.net/pbfiles"
	"net"
	"net/rpc"
)

type ProdService struct {}

func (this *ProdService) GetStock(req prod.ProdRequest,res *prod.ProdResponse)  error{

	if req.ProdId==500{
		res.ProdStock=20
	}else {
		res.ProdStock=30
	}
	return nil
}
func main()  {


	lis,_:=net.Listen("tcp",":8082")
	rpc.Register(new(ProdService))

	for  {
		client,_:=lis.Accept()
		rpc.ServeConn(client)
	}



}
