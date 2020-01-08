package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"topic.chainphp.com/src"
)

//type Topic struct {
//	TopicID int
//	TopicTitle string
//}
func main2()  {

}
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
func main5()  {
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
		//v1.GET("/:topic_id", src.CacheDecorator(src.GetTopicDetial,"topic_id","topic_%s",src.Topics{}))

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

	// route.Run() //gin框架封装的httpserver方法，其实就是下面的原生方法

	//使用go原生启动一个httpserver
	server:=&http.Server{
		Addr:":8080",
		Handler:route,
	}

	go(func() {
		err:=server.ListenAndServe()
		if err !=nil{
			log.Fatal("服务器启动失败",err)
		}
	})()
	go func() {
		src.InitDB()
	}()

	//c:=make(chan os.Signal)
	//signal.Notify(c)  //监听os信号
	//s:=<-c //从chan读取数据，阻塞主进程，让服务一直继续
	//fmt.Println(s) //打印出收到的os信号

	src.ServerNotify()
	//在这里可以做一些资源释放等善后操作。。。
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()
	err:=server.Shutdown(ctx)
	if err !=nil{
		log.Fatalf("服务器关闭")
	}
	log.Println("服务器优雅退出")

}








func main()  {


	var v1 int            // 整型
	var v2 string         // 字符串
	var v3 bool           // 布尔型
	var v4 [10]int        // 数组，数组元素类型为整型
	var v5 struct {       // 结构体，成员变量 f 的类型为64位浮点型
		f float64
	}
	var v6 *int           // 指针，指向整型
	var v7 map[string]int   // map（字典），key为字符串类型，value为整型
	var v16 []int
	var v8 func(a int) int  // 函数，参数类型为整型，返回值类型为整型
	var(                    //同时声明多个变量
		v12 int
		v13 string
	)
	var v14,v15 int

	fmt.Println(v1,v2,v3,v4,v5,v6,v7,v8,v12,v13,v14,v15,v16)
	//仅声明变量类型，声明之后会自动初始化对应类型的0值(int类型的0，string类型的"",bool类型的false等)
	//打印结果 0  false [0 0 0 0 0 0 0 0 0 0] {0} <nil> map[] <nil>

	var v9 int = 10   // 方式一，常规的初始化操作
	var v10 = 10       // 方式二，此时变量类型会被编译器自动推导出来
	v11 := 10          // 方式三，可以省略 var，编译器可以自动推导出v3的类型
	fmt.Println(v9,v10,v11)

	var i=1
	var j=2
	i, j = j, i    //多重赋值
	fmt.Println(i,j)

	name,_:=GetName()  //匿名变量 使用_申明改变量，该变量被丢弃
	fmt.Println(name)


}
func GetName() (username,err string){
	return "zhangsan","no"
}
