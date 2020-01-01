package main

import "com.jtthink.net/httpserver/core"

type NewsControlelr struct {
	core.MyController
}

func (this *NewsControlelr) GET()  {
	this.Ctx.WriteString("this newscontroller get")
}
func (this *NewsControlelr) POST()  {
	this.Ctx.WriteString("this newscontroller post")
}
