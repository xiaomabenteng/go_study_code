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
func GetTopicDetial(c *gin.Context)  {
	c.String(200,"ID%s的帖子的详细",c.Param("topic_id"))
}
func NewTopic(c *gin.Context)  {
	c.String(200,"新增帖子")
}
func DeleteTopic(c *gin.Context)  {
	c.String(200,"删除帖子")
}