package src

import (

	"github.com/gin-gonic/gin"
	//"github.com/gomodule/redigo/redis"
	//"github.com/pquerna/ffjson/ffjson"
	//"log"
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
	//c.JSON(200,CreateTopic(101,"101帖子详情"))

	//直接db查询
	//db, _ := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	//db.LogMode(true)
	//tc:=TopicModel{}
	//db.Table("topics").First(&tc,c.Param("topic_id")) //用table设置真实表明，
	//c.JSON(200,tc)

	//使用DBhelper查询
	tid:=c.Param("topic_id")
	topics:=Topics{}
	//DBHelper.Find(&topics,tid)
	//c.JSON(200,topics)
	//key:="topic_"+tid
	//conn:=RedisDefaultPool.Get()
	//defer conn.Close()
	//top,err:=redis.Bytes(conn.Do("get",key))
	//if err !=nil{ //没取到
	//
	//	DBHelper.Find(&topics,tid)
	//	retDate,_:=ffjson.Marshal(topics)
	//	if topics.TopicID==0{
	//		conn.Do("setex",key,20,retDate)
	//	}else{
	//		conn.Do("setex",key,50,retDate)
	//	}
	//
	//	log.Print("从数据库读取")
	//	c.JSON(200,topics)
	//
	//
	//}else{
	//	log.Print("从redis读取")
	//	ffjson.Unmarshal(top,&topics)
	//	c.JSON(200,topics)
	//}


		DBHelper.Find(&topics,tid)
		c.Set("dbResult",topics)

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
func NewTopics(c *gin.Context)  {
	topics:=TopicsModel{}
	err:=c.BindJSON(&topics)
	if err!=nil {
		c.String(400,"参数错误:%s",err.Error())
	}else {
		c.JSON(200,topics)
	}
}
func DeleteTopic(c *gin.Context)  {
	c.String(200,"删除帖子")
}