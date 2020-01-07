package main

import (
	"net/http"
)

type web1handle struct {
}
func (web1handle)  ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("8081网站首页"))
}
type web2handle struct {
}
func (web2handle)  ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("8082网站首页"))
}
func main()  {

	c:=make(chan bool)
	go func() {
		http.ListenAndServe(":8081",web1handle{})
	}()
	go func() {
		http.ListenAndServe(":8082",web2handle{})
	}()
	<-c

}
