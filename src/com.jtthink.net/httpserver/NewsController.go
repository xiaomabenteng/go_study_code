package main

import "com.jtthink.net/httpserver/core"

type NewsControlelr struct {
	core.MyController
}

func (this *NewsControlelr) GET()  {
	//this.Ctx.WriteJSON(core.Map{"name":"zhangsan"})
	p:=this.GetParam("username","lisi")
	this.Ctx.WriteString(p)
}
func (this *NewsControlelr) POST()  {
	//this.Ctx.WriteString("this newscontroller post")

	//p:=this.PostParam("username","wangwu")
	//this.Ctx.WriteString(p)

	user:=UserModel{}
	this.JSONParam(&user)
	this.Ctx.WriteJSON(user)
}

