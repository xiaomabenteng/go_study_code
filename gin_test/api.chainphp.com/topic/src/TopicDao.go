package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MustLogin()  gin.HandlerFunc{

	return func(c *gin.Context) {
		if c.Query("token")=="" {
			c.String(http.StatusUnauthorized,"缺少token参数")
			c.Abort()
		}else {
			c.Next()
		}
	}
}
func GetTopicList(c *gin.Context)  {

	query:=TopicQuery{}
	err:=c.BindQuery(&query)
	if err!=nil{
		c.String(400,"参数错误%s",err.Error())
	}else{
		c.JSON(200,query)
	}

}
func GetTopicDetial(c *gin.Context)  {
	//c.String(200,"ID%s的帖子的详细",c.Param("topic_id"))
	c.JSON(200,CreateTopic(101,"101帖子详情"))
}
func NewTopic(c *gin.Context)  {
	topic:=TopicModel{}
	err:=c.BindJSON(&topic)
	if err!=nil {
		c.String(400,"参数错误:%s",err.Error())
	}else {
		c.JSON(200,topic)
	}
}
func DeleteTopic(c *gin.Context)  {
	c.String(200,"删除帖子")
}