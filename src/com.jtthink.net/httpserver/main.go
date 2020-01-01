package main

import (
	"com.jtthink.net/httpserver/core"
	"fmt"
	"net/http"
)

func main()  {

	router:=core.DefaultRouter()
	//router.Add("/", &NewsControlelr{})

	err:=http.ListenAndServe(":8099",router)
	if err!=nil{
		fmt.Println(err)
	}
}
