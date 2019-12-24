package main

import (
	_ "com.jtthink/Servicesb"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {

	//arr:=[5]int{1,2,3} //数组
	//fmt.Println(arr)
	//fmt.Println(len(arr),cap(arr))
	//
	//s:=[]int{1,2,3} //切片,代值初始化
	//fmt.Println(s)
	//fmt.Println(len(s),cap(s))
	//
	//var ss[]int  //声明切片
	//ss=make([]int,5) //初始化切片
	//fmt.Println(ss)
	//
	//sss:=make([]int,5) //声明切片
	//fmt.Println(sss)

	//arr:=[5]int{1,2,3,4,5}
	//s:=arr[0:4] //从数组生成切片，左闭右开，>=0 && <4
	//fmt.Println(s)
	//fmt.Println(len(s),cap(s)) //长度是实际元素个数，容量是数组长度
	//s=append(s,1,2,4,5) //超过容量的添加会自动扩容，
	//fmt.Println(len(s),cap(s))






	//var service core.IServices=s.UserService{}
	//Println(service.Get(1))

	//core.SetService(servicesb.UserService{})
	//fmt.Println(core.GetService().Get(1))


	//fmt.Println(core.GetService().Get(1))

// // 链接mysql数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4")
	if err != nil{
		fmt.Println("连接出错"+err.Error())
		return
	}
	rows,err:=db.Query("select user_id,user_name from user")
	if err != nil{
		fmt.Println("查询出错"+err.Error())
		return
	}
	fmt.Println(rows.Columns())

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
