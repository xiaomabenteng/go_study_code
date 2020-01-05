package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"topic.chainphp.com/src"
)

//type Topic struct {
//	TopicID int
//	TopicTitle string
//}
func main()  {
	db, _ := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
db.LogMode(true)
	//db.SingularTable(true)//设置不让gorm自动给表明加复数

	//tc:=src.TopicClass{}
	//db.Table("topic_class").First(&tc,2) //用table设置真实表明，
	//fmt.Println(tc)
	//
	//tcs:=&[]src.TopicClass{}
	//db.Table("topic_class").Find(&tcs)
	//fmt.Println(tcs)

	tcs1:=[]src.TopicClass{}
	db.Table("topic_class").Where("class_name = ?", "技术类").Find(&tcs1)
	//db.Table("topic_class").Where(&src.TopicClass{ClassName:"技术类"}).Find(&tcs1)

	fmt.Println(tcs1)



	//使用原生sql查询
	//rows,_:=db.Raw("select topic_id,topic_title from topics").Rows()
	//defer db.Close()
	//for rows.Next(){
	//	var t_id int
	//	var t_title string
	//	rows.Scan(&t_id,&t_title) //注意sql出几个数据，必须scan几个数据
	//	fmt.Println(t_id,t_title)
	//
	//}



	//route:=gin.Default()
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("topicurl", src.TopicUrl)
	//	v.RegisterValidation("topicValidate", src.TopicValidate)
	//}
	//
	//
	//v1:=route.Group("/v1/topics") //路由分组
	//{//代码块。跟路由分组没关系，仅仅是为了代码看起来清晰。
	//	v1.GET("", src.GetTopicList)
	//	v1.GET("/:topic_id", src.GetTopicDetial)
	//
	//	//v1.Use(src.MustLogin())
	//	v1.POST("", src.NewTopic)
	//	v1.DELETE("/:topic_id", src.DeleteTopic)
	//}
	//v2:=route.Group("/v1/mtopics") //路由分组
	//{
	//	v2.Use(src.MustLogin())
	//	{
	//		v2.POST("", src.NewTopics)
	//	}
	//}
	//
	//route.Run()



}

