package core

type MyController struct {
	Ctx *MyContext
}

func (this *MyController) Init(ctx *MyContext) {
	this.Ctx=ctx
}


type ControllerInterface interface {
	Init(ctx *MyContext)
	GET()
	POST()
}