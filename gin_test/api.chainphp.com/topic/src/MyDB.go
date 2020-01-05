package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var DBHelper *gorm.DB
var Err error
func init()  {
	//注意下面这个错误invalid memory address or nil pointer。。。 ：=快速初始化并赋值。生成的变量和func外面的变量变成两个了
	//DBHelper,Err := gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local")
	DBHelper,Err = gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if Err != nil{
		fmt.Println(Err)
	}
	DBHelper.LogMode(true)
	//使用自带简易连接池
	DBHelper.DB().SetMaxIdleConns(10)
	DBHelper.DB().SetMaxOpenConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)

}