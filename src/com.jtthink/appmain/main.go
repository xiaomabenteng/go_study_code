package main

import (
	_ "com.jtthink/Servicesb"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type all interface{} //声明一个空接口，

func main()  {

	//var i all=123 //空接口可以赋值为任意类型
	//fmt.Println(i)
	//
	//var ii interface{} = "zhangsan" //匿名空接口
	//fmt.Println(ii)
	//
	//list:=make([]interface{},2) //声明一个空类型，类型是空接口
	//list[0]=1
	//list[1]="lisi"
	//fmt.Println(list)

// // 链接mysql数据库
	db,_ := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4")

	rows,_:=db.Query("select user_id,user_name from user")
	fmt.Println(rows.Columns())

	allRows:=make([]interface{},0)
	for rows.Next(){
		var uid int
		var uname string
		rows.Scan(&uid,&uname)
		//oneRow:=make([]interface{},2) //第一种方式
		//oneRow[0]=uid
		//oneRow[1]=uname
		//allRows=append(allRows,oneRow)

		allRows=append(allRows,[]interface{}{uid,uname}) //第二种简化方式

	}
	fmt.Println(allRows)

}
