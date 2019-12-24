package main

import (
	_ "com.jtthink/Servicesb"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type all interface{} //声明一个空接口，

func main()  {

	////m:=map[string]string{"name":"zhangsan"} //map类型 第一种初始化方式
	//m:=make(map[string]string) //第二种初始化方式
	//m["username"]="lisi"
	//delete(m,"username") //删除一个键
	//fmt.Println(m)

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

	allRows:=make([]interface{},0) //定义大切片
	for rows.Next(){
		//var uid int
		//var uname string
		//rows.Scan(&uid,&uname)
		////oneRow:=make([]interface{},2) //第一种方式
		////oneRow[0]=uid
		////oneRow[1]=uname
		////allRows=append(allRows,oneRow)
		//
		//allRows=append(allRows,[]interface{}{uid,uname}) //第二种简化方式


		oneRow:=make([]interface{},2)
		rows.Scan(&oneRow[0],&oneRow[1]) //每一行数据
		//for i:=0;i<len(oneRow) ;i++  { 第一种for方式
		//	fmt.Println(i,oneRow[i]) //返回的是ascii值
		//}
		for i,val:=range oneRow{  //第二种for方式
		v,ok:=val.([]byte) //类型断言
		if(ok){
			//fmt.Println(i,string(v)) //类型断言后才能转换
			oneRow[i]=string(v)
		}

		}
		//b:=[]byte{97,98,99}
		//fmt.Println(string(b)) //字节数组转成string
		allRows=append(allRows,oneRow) //每一行塞进大切片


	}
	fmt.Println(allRows)

}
