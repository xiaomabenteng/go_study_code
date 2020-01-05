package main

import (
	"github.com/gin-gonic/gin"
	"topic.chainphp.com/src"
)

type Topic struct {
	TopicID int
	TopicTitle string
}
func main()  {


	////go 代码块
	//a:=1
	//{
	//	b:=2 //b声明在代码块里面，作用域仅代码块内可用
	//	fmt.Println(a)
	//	fmt.Println(b)
	//}
	//fmt.Println(a)
	//fmt.Println(b)


	route:=gin.Default()
	//route.GET("/v1/topics", func(c *gin.Context) {
	//	if c.Query("username")  == ""{
	//		c.String(200,"topic列表")
	//	}else{
	//		c.String(200,"用户%s的topic列表",c.Query("username"))
	//	}
	//})
	//route.GET("/v1/topics/:topic_id", func(c *gin.Context) {
	//	c.String(200,"ID%s的帖子的详细",c.Param("topic_id"))
	//})

	v1:=route.Group("/v1/topics") //路由分组
	{//代码块。跟路由分组没关系，仅仅是为了代码看起来清晰。
		v1.GET("", src.GetTopicList)
		v1.GET("/:topic_id", src.GetTopicDetial)

		//v1.Use(src.MustLogin())
		v1.POST("", src.NewTopic)
		v1.DELETE("/:topic_id", src.DeleteTopic)
	}


	route.Run()



}
