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

c:=make(chan map[int][]byte)
for i:=1;i<=3;i++{

	go func(index int) {
		url:=fmt.Sprintf(url,index) //格式化字符串，生成每一页页码url
		res,_:=http.Get(url) //http请求
		cnt,_:=ioutil.ReadAll(res.Body) //获取body
		c<-map[int][]byte{index:cnt} //将序号，和网页内容写入channel
		if index==3{
			close(c)
		}
	}(i)
}
	for getcnt:=range c { //主线程干的事，不断地range channel 直到channel关闭
		for i,v:=range getcnt{
			ioutil.WriteFile(fmt.Sprintf("./files/%d",i),v,666) //输出文件
		}
	}






}
