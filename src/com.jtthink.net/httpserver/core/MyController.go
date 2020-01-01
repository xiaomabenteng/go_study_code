package core

import (
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
)

type MyController struct {
	Ctx *MyContext
}

func (this *MyController) Init(ctx *MyContext) {
	this.Ctx=ctx
}
//获取get参数
func (this *MyController) GetParam(key string,defaultValue ...string) string {
	ret:=this.Ctx.request.URL.Query().Get(key)
	if ret=="" && len(defaultValue)>0 {
		return defaultValue[0]
	}
	return ret
}
//获取post,单参数
func (this *MyController) PostParam(key string,defaultValue ...string) string {
	ret:=this.Ctx.request.PostFormValue(key)
	if ret=="" && len(defaultValue)>0 {
		return defaultValue[0]
	}
	return ret
}
//获取json参数
func (this *MyController) JSONParam(obj interface{})  {
	 body,err:=ioutil.ReadAll(this.Ctx.request.Body)
	if err!=nil {
		panic(err)
	}
	err=ffjson.Unmarshal(body,obj)
	if err != nil {
		panic(err)
	}


}
type ControllerInterface interface {
	Init(ctx *MyContext)
	GET()
	POST()
}