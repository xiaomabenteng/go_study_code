package main

import (
	"com.jtthink.net/httpserver/core"
	"fmt"
	"net/http"
)

func main()  {

	router:=core.DefaultRouter()
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("get 这是根目录"))
	})
	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("post 这是根目录"))
	})
	err:=http.ListenAndServe(":8099",router)
	if err!=nil{
		fmt.Println(err)
	}
}
