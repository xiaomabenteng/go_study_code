package main

import (
	"com.jtthink.net/httpserver/core"
	"fmt"
	"net/http"
)

func main()  {

	router:=core.DefaultRouter()
	router.Get("/", func(ctx *core.MyContext) {
		ctx.WriteString("ctx get")
	})
	router.Post("/", func(ctx *core.MyContext) {
		ctx.WriteString("ctx post")
	})
	err:=http.ListenAndServe(":8099",router)
	if err!=nil{
		fmt.Println(err)
	}
}
