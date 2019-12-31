package core

import "net/http"

type MyContext struct { //上下文对象
	request *http.Request
	http.ResponseWriter
}

func (this *MyContext) WriteString(str string)  {
	this.Write([]byte(str))
}
