package core

import (
	"github.com/pquerna/ffjson/ffjson"
	"net/http"
)

type MyContext struct { //上下文对象
	request *http.Request
	http.ResponseWriter
}

func (this *MyContext) WriteString(str string)  {
	this.Write([]byte(str))
}
func (this *MyContext) WriteJSON(m interface{})  {
	this.Header().Add("Content-type","application/json")
	ret,err:=ffjson.Marshal(m)
	if err !=nil{
		panic(err)
	}
	this.Write(ret)
}