package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)
type all interface{} //声明一个空接口，
func test(str ...string)  { //不定参数设置
	for _,v:= range str {
		fmt.Println(v)
	}
}
func main()  {
	//mycore.ShowName()


	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4")
	if err != nil{
		fmt.Println("连接错"+err.Error())
		return
	}
	rows,err:=db.Query("select user_id,user_name from user limit 0,10")
	if err != nil{
		fmt.Println("查询错误"+err.Error())
		return
	}
	columns,err:=rows.Columns()
	fmt.Println(columns)
	if err != nil{
		fmt.Println("获取列错误"+err.Error())
		return
	}
	allRows:=make([]interface{}, 0)
	oneRow := make([]interface{}, len(columns))
	fmt.Println(oneRow)
	//userModels:=[]models.UserModel{}
	//for rows.Next(){
	//	temp:=models.UserModel{}
	//	rows.Scan(&temp.Uid,&temp.Uname)
	//	userModels=append(userModels,temp)
	//}

	for rows.Next(){

		rows.Scan(oneRow)
		allRows=append(allRows,oneRow)
	}


	fmt.Println(allRows)
}
