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

	//变量的作用域.函数体外声明代表全局变量可以在整个包，甚至外部包使用(变量以大写字母开头)。函数体内声明为局部变量，只在函数体内有效。
	//var v1 int            // 整型
	//var v2 string         // 字符串
	//var v3 bool           // 布尔型
	//var v4 [10]int        // 数组，数组元素类型为整型
	//var v5 struct {       // 结构体，成员变量 f 的类型为64位浮点型
	//	f float64
	//}
	//var v6 *int           // 指针，指向整型
	//var v7 map[string]int   // map（字典），key为字符串类型，value为整型
	//var v16 []int
	//var v8 func(a int) int  // 函数，参数类型为整型，返回值类型为整型
	//var(                    //同时声明多个变量
	//	v12 int
	//	v13 string
	//)
	//var v14,v15 int
	//
	//fmt.Println(v1,v2,v3,v4,v5,v6,v7,v8,v12,v13,v14,v15,v16)
	////仅声明变量类型，声明之后会自动初始化对应类型的0值(int类型的0，string类型的"",bool类型的false等)
	////打印结果 0  false [0 0 0 0 0 0 0 0 0 0] {0} <nil> map[] <nil>
	//
	//var v9 int = 10   // 方式一，常规的初始化操作
	//var v10 = 10       // 方式二，此时变量类型会被编译器自动推导出来
	//v11 := 10          // 方式三，可以省略 var，编译器可以自动推导出v3的类型
	//fmt.Println(v9,v10,v11)
	//
	//var i=1
	//var j=2
	//i, j = j, i    //多重赋值
	//fmt.Println(i,j)
	//
	//name,_:=GetName()  //匿名变量 使用_申明改变量，该变量被丢弃
	//fmt.Println(name)
	//
	//
	////常量定义,由于常量的赋值是一个编译期行为，所以右值不能出现任何需要运行期才能得出结果的表达式
	//const Pi float64 = 3.14159265358979323846
	//const zero = 0.0 // 无类型浮点常量
	//const (          // 通过一个 const 关键字定义多个常量，和 var 类似
	//	size int64 = 1024
	//	eof = -1  // 无类型整型常量
	//)
	//const u, v float32 = 0, 3  // u = 0.0, v = 3.0，常量的多重赋值
	//const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量


	//预定义常量 true、false 和 iota(可以被认为是一个可被编译器修改的常量，在每一个 const 关键字出现时被重置为 0，然后在下一个 const 出现之前，每出现一次 iota，其所代表的数字会自动增 1)
	//作用于同变量
	const (    // iota 被重置为 0
		c0 = iota   // c0 = 0
		c1 = iota   // c1 = 1
		c2 = iota   // c2 = 2
	)
	const (
		u = iota * 2;  // u = 0
		v = iota * 2;  // v = 2
		w = iota * 2;  // w = 4
	)
	const x = iota;  // x = 0
	const y = iota;  // y = 0

	const (             //枚举
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println(Sunday,Monday) //输出 0 1


/*
//基本数据类型
	布尔类型：bool
	整型：int8、byte、int16、int、uint、uintptr 等
	浮点类型：float32、float64
	复数类型：complex64、complex128
	字符串：string
	字符类型：rune
	错误类型：error
//复合类型
   指针（pointer）
   数组（array） 类似php数组
   切片（slice）
   字典（map）  类似php关联数组
   通道（chan）
   结构体（struct） 类似php中class
   接口（interface）
 */
/*
整型
	类型	长度（单位：字节）	说明	         值范围	默认值
	int8	1	         带符号8位整型	        -128~127	0
	uint8	1	         无符号8位整型，与 byte 类型等价	0~255	0
	int16	2	         带符号16位整型	     -32768~32767	0
	uint16	2	         无符号16位整型	       0~65535	0
	int32	4	        带符号32位整型，与 rune 类型等价	-2147483648~2147483647	0
	uint32	4	        无符号32位整型	     0~4294967295	0
	int64	8	        带符号64位整型	     -9223372036854775808~9223372036854775807	0
	uint64	8	       无符号64位整型	    0~18446744073709551615	0
	int	32位或64位	   与具体平台相关	    与具体平台相关	0    《php中的int类型》
	uint	32位或64位	与具体平台相关	    与具体平台相关	0
	uintptr	与对应指针相同	无符号整型，足以存储指针值的未解释位	32位平台下为4字节，64位平台下为8字节	0
 */
	var int_value_1 int8
	int_value_2 := 8   // int_value_2 将会被自动推导为 int 类型
	//int_value_1 = int_value_2  // 编译错误
	int_value_1 = int8(int_value_2) // 强制类型转换,编译通过。不同类型之间不能赋值，也不能做运算，不能做比较


}
func GetName() (username,err string){
	return "zhangsan","no"
}
