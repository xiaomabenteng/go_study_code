package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type all interface{} //声明一个空接口，

func main()  {

	var i all=123 //空接口可以赋值为任意类型
	fmt.Println(i)

	var ii interface{} = "zhangsan" //匿名空接口
	fmt.Println(ii)

	list:=make([]interface{},2) //声明一个空类型，类型是空接口
	list[0]=1
	list[1]="lisi"
	fmt.Println(list)

//// // 链接mysql数据库
//	db,_ := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4")
//
//	rows,_:=db.Query("select user_id,user_name from user")
//	fmt.Println(rows.Columns())

	//// var newsModel models.NewsModel=models.NewsModel{} // 两种变量申明方式
	//newsModel:=models.NewsModel{} // 两种变量申明方式
	//for rows.Next(){
	//	//var uid int
	//	//var uname string
	//	//rows.Scan(&uid,&uname) //赋值查询结果到变量
	//	//fmt.Println(uid,uname)
	//
	//	rows.Scan(&newsModel.Id,&newsModel.Title) //查询结果赋值到结构体
	//	fmt.Println(newsModel)
	//}

	//newsModels:=[]models.NewsModel{} //声明切片，返回一个新闻列表
	//for rows.Next(){
	//	tmp:=models.NewsModel{}
	//	rows.Scan(&tmp.Id,&tmp.Title) //查询结果赋值到结构体
	//	newsModels=append(newsModels,tmp)
	//}
	//fmt.Println(newsModels)

}
