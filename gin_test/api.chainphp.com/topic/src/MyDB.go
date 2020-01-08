package src

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DBHelper *gorm.DB
var Err error
func InitDB()  {
	//注意下面这个错误invalid memory address or nil pointer。。。 ：=快速初始化并赋值。生成的变量和func外面的变量变成两个了
	//DBHelper,Err := gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local")
	//DBHelper,Err = gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local") //家里的数据库
	DBHelper,Err = gorm.Open("mysql", "root:123456@tcp(192.168.10.133:3306)/test?charset=utf8mb4&parseTime=True&loc=Local") //公司的数据库配置
	if Err != nil{
		//fmt.Println(Err)
		//log.Fatal("DB初始化错误",Err) //<遇到错误退出程序第一种方式，>向控制台输出错误信息，并停止程序
		ShuDownServer(Err) //<遇到错误退出程序第二种方式，监听信号>
	}
	DBHelper.LogMode(true)
	//使用gorm自带简易连接池
	DBHelper.DB().SetMaxIdleConns(10)
	DBHelper.DB().SetMaxOpenConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)

}