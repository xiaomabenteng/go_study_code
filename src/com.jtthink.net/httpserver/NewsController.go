package main

import "com.jtthink.net/httpserver/core"

type NewsControlelr struct {
	core.MyController
}

func (this *NewsControlelr) GET()  {
	this.Ctx.WriteJSON(map[string]string{"name":"zhangsan"})
}
func (this *NewsControlelr) POST()  {
	this.Ctx.WriteString("this newscontroller post")
}
