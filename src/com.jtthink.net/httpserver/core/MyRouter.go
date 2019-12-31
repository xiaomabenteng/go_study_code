package core

import (
	"net/http"
)

type MyHandleFunc func(Ctx *MyContext)

type MyRouter struct {
	Mapping map[string]map[string]MyHandleFunc
	Ctx *MyContext
}

func DefaultRouter()  *MyRouter{
	return &MyRouter{make(map[string]map[string]MyHandleFunc),&MyContext{}}
}

func (this *MyRouter) Get(path string,f MyHandleFunc)  {
	if this.Mapping["GET"]==nil{
		this.Mapping["GET"]=make(map[string]MyHandleFunc)
	}
	this.Mapping["GET"][path]=f
}
func (this *MyRouter) Post(path string,f MyHandleFunc)  {
	if this.Mapping["POST"]==nil{
		this.Mapping["POST"]=make(map[string]MyHandleFunc)
	}
	this.Mapping["POST"][path]=f
}
func (this *MyRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {

	this.Ctx.request=request
	this.Ctx.ResponseWriter=writer

	if f,OK:=this.Mapping[request.Method][request.URL.Path];OK{
		f(this.Ctx)
	}


}