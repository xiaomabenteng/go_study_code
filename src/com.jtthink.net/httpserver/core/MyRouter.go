package core

import (
	"net/http"
)

type MyHandleFunc func(Ctx *MyContext)

type MyRouter struct {
	Mapping map[string]ControllerInterface

}

func DefaultRouter()  *MyRouter{
	return &MyRouter{make(map[string]ControllerInterface)}
}

//加入 path和controller的对应关系
func (this *MyRouter) Add(path string,c ControllerInterface)  {
	this.Mapping[path]=c
}

func (this *MyRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {

	if f,OK:=this.Mapping[request.URL.Path];OK{
		f.Init(&MyContext{request,writer})
		if request.Method=="GET"{
			f.GET()
		}else if request.Method=="POST" {
			f.POST()
		}
	}


}