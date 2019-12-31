package main

import (
	"fmt"
	"net/http"
)
type MyHandler struct {

}

func (*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	fmt.Println(request.RequestURI)
	if request.RequestURI=="/woshi"{
		writer.Write([]byte("zhangsan"))
	}else{
		writer.Write([]byte("hello myhandler"))
	}

}

func main()  {

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello"))
	//})
	//http.HandleFunc("/abc", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("abc"))
	//})
	////err:=http.ListenAndServe(":8099",nil) //启动一个http服务，监听指定端口
	////if err!=nil {
	////	fmt.Println(err)
	////}
	//server:=http.Server{Addr:":8099",Handler:nil} //第二种方式启动http服务，其实和第一种是一样的
	//err:=server.ListenAndServe()
	//if err!=nil{
	//	fmt.Println(err)
	//}



	//handler:=new(MyHandler)
	//server:=http.Server{Addr:":8099",Handler:handler} //传递自定义handler
	//err:=server.ListenAndServe()
	//if err!=nil{
	//	fmt.Println(err)
	//}

	mymux:=http.NewServeMux()                // 自定义路由管理
	mymux.Handle("/hello",new(MyHandler)) //第一种方式,使用自定义handler

	mymux.HandleFunc("/abc", func(writer http.ResponseWriter, request *http.Request) {  // 第二种方式
		writer.Write([]byte("abc"))
	})
	server:=http.Server{Addr:":8099",Handler:mymux}
	err:=server.ListenAndServe()
	if err!=nil{
		fmt.Println(err)
	}
}
