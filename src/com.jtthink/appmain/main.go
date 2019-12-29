package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)

func test() {
	panic("exp")
}

func main()  {



url:="https://news.cnblogs.com/n/page/%d/"

c:=make(chan map[int][]byte)
for i:=1;i<=3;i++{

	go func(index int) {
		defer func() {
			if err:=recover();err!=nil{ //捕获异常并处理
				fmt.Println(err)
			}
			if index==3 { //关闭通道，否则主函数一直在执行
				close(c)
			}
		}()
		url:=fmt.Sprintf(url,index) //格式化字符串，生成每一页页码url
		res,_:=http.Get(url) //http请求
		cnt,_:=ioutil.ReadAll(res.Body) //获取body
		if index==3 {  //模拟在执行第三页网页抓取时抛出异常
			test()
		}
		c<-map[int][]byte{index:cnt} //将序号，和网页内容写入channel

	}(i)
}
	for getcnt:=range c { //主线程干的事，不断地range channel 直到channel关闭
		for i,v:=range getcnt{
			ioutil.WriteFile(fmt.Sprintf("./files/%d",i),v,666) //输出文件
		}
	}






}
