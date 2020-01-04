package main

import (
	"github.com/gin-gonic/gin"
)

type Topic struct {
	TopicID int
	TopicTitle string
}
func main()  {

	route:=gin.Default()
	route.GET("/", func(context *gin.Context) {
		//context.JSON(http.StatusOK,Topic{TopicID:100,TopicTitle:"话题"})
		context.JSON(200,gin.H{"name":"zhangsan"})
	})
	route.Run()


}
