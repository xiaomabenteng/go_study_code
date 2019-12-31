package core

import (
	"net/http"
)

type MyRouter struct {
	Mapping map[string]map[string]http.HandlerFunc
}

func DefaultRouter()  *MyRouter{
	return &MyRouter{make(map[string]map[string]http.HandlerFunc)}
}

func (this *MyRouter) Get(path string,f http.HandlerFunc)  {
	if this.Mapping["GET"]==nil{
		this.Mapping["GET"]=make(map[string]http.HandlerFunc)
	}
	this.Mapping["GET"][path]=f
}
func (this *MyRouter) Post(path string,f http.HandlerFunc)  {
	if this.Mapping["POST"]==nil{
		this.Mapping["POST"]=make(map[string]http.HandlerFunc)
	}
	this.Mapping["POST"][path]=f
}
func (this *MyRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {

	f:=this.Mapping[request.Method][request.URL.Path]
	f(writer,request)

}