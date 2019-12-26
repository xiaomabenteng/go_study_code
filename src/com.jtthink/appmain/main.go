package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)


func main()  {


url:="https://news.cnblogs.com/n/page/%d/"

for i:=1;i<=3;i++{
	url:=fmt.Sprintf(url,i) //格式化字符串
	res,_:=http.Get(url) //http请求
	cnt,_:=ioutil.ReadAll(res.Body) //获取body
	ioutil.WriteFile(fmt.Sprintf("./files/%d",i),cnt,666) //输出文件
}








}
