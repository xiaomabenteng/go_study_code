package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v9"
	"os"
	"os/signal"
	"time"
	"topic.chainphp.com/src"
)

//type Topic struct {
//	TopicID int
//	TopicTitle string
//}
func main4()  {

	conn:=src.RedisDefaultPool.Get()
	str,_:=redis.String(conn.Do("get","name"))
	fmt.Println(str)


}
func main3()  {
	count:=0
	go func() {
		for{
			fmt.Println("执行",count)
			count++
			time.Sleep(time.Second*1)
		}
	}()



	c:=make(chan os.Signal)

	//go func() { //启动一个协程5秒后发送 退出信号，整个程序停止运行
	//	for i:=0;i<=5;i++{
	//		if i==5 {
	//			c<-os.Interrupt
	//		}
	//		time.Sleep(time.Second*1)
	//	}
	//}()
	go func() { //使用上下文包的形式。 启动一个协程5秒后发送 退出信号，整个程序停止运行
		ctx,_:=context.WithTimeout(context.Background(),time.Second*5)
		select {
			case <-ctx.Done():
				c<-os.Interrupt
		}
	}()

	signal.Notify(c)  //监听os信号
	s:=<-c
	fmt.Println(s)

}
func main()  {
	//db, _ := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	//db.LogMode(true)
	//db.SingularTable(true)//设置不让gorm自动给表明加复数

	//tc:=src.TopicClass{}
	//db.Table("topic_class").First(&tc,2) //用table设置真实表明，
	//fmt.Println(tc)
	//
	//tcs:=&[]src.TopicClass{}
	//db.Table("topic_class").Find(&tcs)
	//fmt.Println(tcs)

	//tcs1:=[]src.TopicClass{}
	//db.Table("topic_class").Where("class_name = ?", "技术类").Find(&tcs1)
	////db.Table("topic_class").Where(&src.TopicClass{ClassName:"技术类"}).Find(&tcs1)
	//fmt.Println(tcs1)

	////数据插入
	//topics:=src.TopicModel{
	//	TopicTitle:"TopicTitle",
	//	TopicShortTitle:"TopicShortTitle",
	//	UserIP:"127.0.0.1",
	//	TopicScore:0,TopicUrl:"testurl",
	//	TopicDate:time.Now()}
	//fmt.Println(db.Table("topics").Create(&topics).RowsAffected)
	//fmt.Println(topics.TopicID)



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



	route:=gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", src.TopicUrl)
		v.RegisterValidation("topicValidate", src.TopicValidate)
	}


	v1:=route.Group("/v1/topics") //路由分组
	{//代码块。跟路由分组没关系，仅仅是为了代码看起来清晰。
		v1.GET("", src.GetTopicList)
		v1.GET("/:topic_id", src.GetTopicDetial)

		//v1.Use(src.MustLogin())
		v1.POST("", src.NewTopic)
		v1.DELETE("/:topic_id", src.DeleteTopic)
	}
	v2:=route.Group("/v1/mtopics") //路由分组
	{
		v2.Use(src.MustLogin())
		{
			v2.POST("", src.NewTopics)
		}
	}

	route.Run()



}

