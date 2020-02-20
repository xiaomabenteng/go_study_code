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
	"reflect"
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
	//name,_:=GetName()  //匿名变量 使用_申明该变量，该变量被丢弃
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
/*
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
	//fmt.Println(Sunday,Monday) //输出 0 1
*/

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
/*
	var int_value_1 int8
	int_value_2 := 8   // int_value_2 将会被自动推导为 int 类型
	int_value_1 = int_value_2  // 编译错误
	int_value_1 = int8(int_value_2) // 强制类型转换,编译通过。不同类型之间不能赋值，也不能做运算，不能做比较
*/

/*
//浮点型

   var float_value_1 float32

   float_value_1 = 10.2
   float_value_2 := 10.2 // 如果不加小数点，float_value_2 会被推导为整型而不是浮点型,自动推导为float64
  // float_value_3 := 1.1E-10
  // float_value_1 = float_value_2  // float_value_2 是 float64 类型，会导致编译报错
   float_value_1 = float32(float_value_2) //强制类型转换。实际开发中，应该尽可能地使用 float64 类型，因为 math 包中所有有关数学运算的函数都会要求接收这个类型

//浮点数的精度
  float_value_4 := 0.1
  float_value_5 := 0.7
  float_value_6 := float_value_4 + float_value_5 // 0.7999999999999999
//浮点数的比较
   p := 0.00001
   // 判断 float_vlalue_1 与 float_value_2 是否相等
   if math.Dim(float64(float_value_1), float_value_2) < p { //解决方案是一种近似判断，通过一个可以接受的最小误差值 p，约定如果两个浮点数的差值在此精度的误差范围之内，则判定这两个浮点数相等。这个解决方案也是其他语言判断浮点数相等所采用的通用方案，PHP 也是这么做的
       fmt.Println("float_value_1 和 float_value_2 相等")
   }
*/

/*
复数类型

   var complex_value_1 complex64

   complex_value_1 = 1.10 + 10i          // 由两个 float32 实数构成的复数类型
   complex_value_2 := 1.10 + 10i         // 和浮点型一样，默认自动推导的实数类型是 float64，所以 complex_value_2 是 complex128 类型
   complex_value_3 := complex(1.10, 10)  // 与 complex_value_2 等价

*/
/*
//字符串
	//默认是通过 UTF-8 编码的字符序列，当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节，比如中文编码通常需要 3 个字节
	var str string         // 声明字符串变量
	str = "Hello World"    // 变量初始化
	//str_2 := "你好，学院君   // 也可以同时进行声明和初始化
	ch := str[0] // 取字符串的第一个字符

	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))  //格式化输出
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	//不可变值类型 在 Go 语言中，字符串是一种不可变值类型，一旦初始化之后，它的内容不能被修改
	//str := "Hello world"
	//str[0] = 'X' // 编译错误

	//字符串操作
	str = str + ", 学院君"
	str += ", 学院君"  // 上述语句也可以简写为这样，效果完全一样

	str = str +   //如果字符串长度较长，需要换行，则「+」必须出现在上一行的末尾
		", 学院君"

	fmt.Println(str)

//字符类型
	//在 Go 语言中支持两个字符类型，一个是 byte（实际上是 uint8 的别名），代表 UTF-8 字符串的单个字节的值；另一个是 rune，代表单个 Unicode 字符。
*/
/*
//数组
	//Go 语言中，数组是固定长度的、同一类型的数据集合。数组中包含的每个数据项被称为数组元素，一个数组包含的元素个数被称为数组的长度
	//数组长度在定义后就不可更改，在声明时可以指定数组长度为一个常量或者一个常量表达式（常量表达式是指在编译期即可计算结果的表达式）
	//var a [8]byte // 长度为8的数组，每个元素为一个字节
	//var b [3][3]int // 二维数组（9宫格）
	//var c [3][3][3]float64 // 三维数组（立体的9宫格）
	//var d = [3]int{1, 2, 3}  // 声明时初始化
	//var e = new([3]string)   // 通过 new 初始化
	//
	//a := [5]int{1,2,3,4,5}
	//a := [...]int{1, 2, 3} //在编译器会自动计算出长度3
	//a := [5]int{1, 2, 3} //[1 2 3 0 0] 如果没有填满，则空位会通过对应的元素类型空值填充
	//a := [5]int{1: 3, 3: 7} //[0 3 0 7 0] 初始化指定下标位置的元素值

	arr := [5]int{1,2,3,4,5}
	//a1, a2 := arr[0], arr[len(arr) - 1] //访问数组元素。a1 的值是 1，a2 的值是 5

	//遍历数组
	for i:=0;i<len(arr);i++{
		//fmt.Println(arr[i])
		fmt.Println("Element", i, "of arr is", arr[i])
		aa:=fmt.Sprintf("%s %d %s %d", "Element",i,"of arr is",arr[i])
		fmt.Println(aa)
	}
	for i,v:=range arr{   //类似php foreach
		fmt.Println("Element", i, "of arr is", v)
	}

	arr[0] = 100 //设置数组元素
*/
/*
//切片

	//是一个可变长度的、同一类型元素集合。切片长度随元素增长增长，但不随元素减少而减少。切片从底层看是对数组做了一层简单封装。
	//var slice []string = []string{"a", "b", "c"}  //声明并初始化
	//创建数组切片
	//months := [...]string{"January", "February", "March", "April", "May", "June"}
	//q2:=months[3:6] //1,基于数组创建切片  长度len(q2)  容量cap(q2)
	//q2_1:=q2[:1]  //2,基于切片创建切片
	//mySlice1 := make([]int, 5)  //3，直接创建。初始元素个数为 5 的数组切片，元素类型为整型，初始值为 0，容量为 5
	//mySlice2 := make([]int, 2, 3) //         初始元素个数为 5 的整型数组切片，初始值为 [0 0 0 0 0]，并预留 10 个元素的存储空间（容量为10）
	//mySlice3 := []int{1, 2, 3, 4, 5} //       直接创建并初始化包含 5 个元素的数组切片

	//newSlice := append(mySlice2, 1, 2, 3,4) //切片增加新元素
	//
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制 slice1 的前3个元素到 slice2 中
	//copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice3 = slice3[:len(slice3) - 5]  // 删除 slice3 尾部5个元素
	slice3 = slice3[5:]  // 删除 slice3 头部 5 个元素
*/

/*
//字典类型
	var testMap map[string]int  //字典声明
	testMap = map[string]int{  //字典初始化
		"one": 1,
		"two": 2,
		"three": 3,
	}

	k := "two"
	va, ok := testMap[k]  //查找元素
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, va)
	}
	delete(testMap, "four")  //删除元素

	var testMap1 = make(map[string]int) //make创建字典
	testMap1["one"] = 1  //元素赋值
	fmt.Println(testMap1)
*/

/*
//指针
	a := 100
	var ptr *int  //指针类型。 声明。表示指向存储 int 类型值的指针
	ptr = &a      //初始化
	//ptr1 := &a    //声明并初始化
	fmt.Println(ptr)  //0xc0000a2000
	fmt.Println(*ptr)  //100
*/

//条件语句
/*
	if condition1 {
		// do something
	} else if condition2 {
		// do something else
	} else {
		// catch-all or default
	}
*/

//分支语句

/*
	score := 100
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80 && score < 90:
		fmt.Println("Grade: B")
	case score >= 70 && score < 80:
		fmt.Println("Grade: C")
	case score >= 60 && score < 70:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	score := 100
	switch score {  //在与 case 分支值判等的时候，才可以这么做(变量名放在switch后面)
	case 90, 100:
		fmt.Println("Grade: A")
	case 80:
		fmt.Println("Grade: B")
	case 70:
		fmt.Println("Grade: C")
	case 65:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
*/
//循环语句
/*
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum := 0
	i := 0
	for {    //无限循环
		i++
		if i > 100 {
			break
		}
		sum += i
	}
	fmt.Println(sum)


	for k, v := range a { //for-range 结构
		fmt.Println(k, v)
	}


	sum := 0
	i := 0
	for i < 100 {   //基于条件判断进行循环


		i++
		sum += i
	}
	fmt.Println(sum)


	arr := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	ITERATOR1:  //break跳到指定标签处
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				break ITERATOR1
			}
			fmt.Println(num)
		}
	}

	arr := [][]int{{1,2,3},{4,5,6},{7,8,9}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr[i][j]
			if j > 1 {
				goto EXIT   //goto语句，跳转到本函数内某个标签
			}
			fmt.Println(num)
		}

	EXIT:
		fmt.Println("Exit.")
*/

//函数
/*
	//func add(a, b int) int  { //普通函数，两个int型参数，返回值为int
	//	return a + b
	//}
	//add(1, 2) //函数调用
	//Add(1, 2)  //首字母小写的函数只能在同一个包中访问，首字母大写的函数才可以在其他包中调用，Go 文件中定义的全局变量也是如此

	//按值传参。
	//把 x、y 变量作为参数传递到 add 函数时，这两个变量会拷贝出一个副本赋值给 a、b 变量作为参数，因此，在 add 函数中调整 a、b 变量的值并不会影响原变量 x、y 的值
	func add(a, b int) int  {
		a *= 2
		b *= 3
		return a + b
	}
	func main()  {
		x, y := 1, 2
		z := add(x, y)
		fmt.Printf("add(%d, %d) = %d\n", x, y, z) //add(1, 2) = 8
	}

	//引用传参
	func add(a, b *int) int {
		*a *= 2
		*b *= 3
		return *a + *b
	}
	func main()  {
		x, y := 1, 2
		z := add(&x, &y)
		fmt.Printf("add(%d, %d) = %d\n", x, y, z) //add(2, 6) = 8
	}

	//多返回值
	func add(a, b *int) (int, error) {
		if (*a < 0 || *b < 0) {
			err := errors.New("只支持非负整数相加")
			return 0, err
		}
		*a *= 2
		*b *= 3
		return *a + *b, nil
	}

	func main()  {
		x, y := -1, 2
		z, err := add(&x, &y)
		if (err != nil) {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("add(%d, %d) = %d\n", x, y, z)
	}

	//返回值命名
	func add(a, b *int) (c int, err error) {
		if (*a < 0 || *b < 0) {
			err = errors.New("只支持非负整数相加")
			return
		}
		*a *= 2
		*b *= 3
		c = *a + *b
		return
	}
*/

/*
//变长参数
	func myfunc(numbers ...int) {  //代表接受不定数量参数，类型全部是int
		for _, number := range numbers {
			fmt.Println(number)
		}
	}
	myfunc(1, 2, 3, 4, 5) //传递参数

	slice := []int{1, 2, 3, 4, 5}
	myfunc(slice...)  //还可以传递一个数组切片。在末尾需加上...表示变长参数
	myfunc(slice[1:3]...)

//任意类型的变长参数
	func myPrintf(args ...interface{}) { //空接口类型可以用于表示任意类型。支持任意类型，任意长度
		for _, arg := range args {
			switch reflect.TypeOf(arg).Kind() {
				case reflect.Int:
					fmt.Println(arg, "is an int value.")
				case reflect.String:
					fmt.Printf("\"%s\" is a string value.\n", arg)
				case reflect.Array:
					fmt.Println(arg, "is an array type.")
				default:
					fmt.Println(arg, "is an unknown type.")
			}
		}
	}
	func main() {
		myPrintf(1, "1", [1]int{1}, true)
	}
	1 is an int value.
	"1" is a string value.
	[1] is an array type.
	false is an unknown type.
*/
/*
//匿名函数
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))  // 调用匿名函数 add
	func(a, b int) {
		fmt.Println(a + b)
	} (1, 2) // 花括号后直接跟参数列表表示直接调用函数

//系统内置函数
	名称	说明
	close	用于在管道通信中关闭一个管道
	len、cap	len 用于返回某个类型的长度（字符串、数组、切片、字典和管道），cap 则是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道）
	new、make	new 和 make 均用于分配内存，new 用于值类型和用户自定义的类型（类），make 用于内置引用类型（切片、字典和管道）。它们在使用时将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针，也可以用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，所以 make 不仅分配内存地址还会初始化对应类型。
	copy、append	分别用于切片的复制和动态添加元素
	panic、recover	两者均用于错误处理机制
	print、println	打印函数，在实际开发中建议使用 fmt 包
	complex、real、imag	用于复数类型的创建和操作
*/
/*
//类定义，初始化
//689行
	student := NewStudent(1, "学院君", 100)
	fmt.Println(student)
	student.SetName("学院君小号")
	fmt.Println("Name:", student.GetName()) //使用类成员方法
	fmt.Println(student) //{id: 1, name: 学院君, male: false, score: 100.000000}。自动调用string方法，以字符串打印类的实例
//类的继承和方法重写
//714行
	animal:=Animal{"狗"}
	dog:=Dog{animal}
	fmt.Println(dog.GetName(), "叫声:", dog.Call(), "喜爱的食物:", dog.FavorFood())
*/
a:=123;
fmt.Println( reflect.TypeOf(a))



}
func GetName() (username,err string){
	return "zhangsan","no"
}
func add(a, b int) int  {
	a *= 2
	b *= 3
	return a + b
}

type Student struct {
	id uint
	name string
	male bool
	score float64
}
func NewStudent(id uint, name string, score float64) *Student {  //类初始化
	return &Student{id: id, name:name, score:score}
}
func (s Student) GetName() string  { //类添加成员方法
	return s.name
}
func (s *Student) SetName(name string) { //SetXXX 方法需要在函数内部修改成员变量的值，并且作用到该函数作用域以外，所以需要传入指针类型
	s.name = name
}
func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}


type Animal struct {
	name string
}
func (a Animal) Call() string {
	return "动物的叫声..."
}
func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}
func (a Animal) GetName() string  {
	return a.name
}

//继承Animal父类
type Dog struct {
	Animal
}
func (d Dog) FavorFood() string {
	return "骨头"
}
func (d Dog) Call() string {
	return "汪汪汪"
}