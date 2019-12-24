package main

import (
	_ "com.jtthink/Servicesb"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type all interface{} //声明一个空接口，
func test(str ...string)  { //不定参数设置
	for _,v:= range str {
		fmt.Println(v)
	}
}

func main()  {
	//test("aa","bb","ccc")
	//
	//list:=[]string{"dd","ee","ff"}
	//test(list...) //将list打散传进去，类似解构

//
// // 链接mysql数据库
	db,_ := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4")

	rows,_:=db.Query("select user_id,user_name,mobile from user")
	fmt.Println(rows.Columns())
	columns,_:=rows.Columns()
	allRows:=make([]interface{},0) //定义大切片
	filedMap:=make(map[string]string)
	oneRow:=make([]interface{},len(columns))
	scanRow:=make([]interface{},len(columns))
	for rows.Next(){


		for i,_ :=range oneRow {
			scanRow[i]=&oneRow[i]
		}
		//rows.Scan(&oneRow[0],&oneRow[1]) //每一行数据(scan里面需要是地址)
		rows.Scan(scanRow...) //每一行数据

		for i,val:=range oneRow{  //第二种for方式
		v,ok:=val.([]byte) //类型断言
		if(ok){
			//fmt.Println(i,string(v)) //类型断言后才能转换
			filedMap[columns[i]]=string(v)
		}

		}
		//b:=[]byte{97,98,99}
		//fmt.Println(string(b)) //字节数组转成string
		allRows=append(allRows,filedMap) //每一行塞进大切片


	}
	fmt.Println(allRows)

}
