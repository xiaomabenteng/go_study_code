package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"sync"
)

func test() {
	panic("exp")
}

func main()  {

	////sync包 有三个方法 Add(int),done(),wait()
	////内部创建一个计数器，Add用来添加goroutine的个数。done用来执行数量减1，好比Add(-1)。wait用来阻塞主线程，知道所有协程执行完毕
	//var WG sync.WaitGroup
	//for i:=1;i<=5;i++{
	//	go func(index int) {
	//		defer func() {
	//			WG.Done()
	//		}()
	//		time.Sleep(time.Second*1)
	//		fmt.Println(index,"执行完成")
	//	}(i)
	//	WG.Add(1)
	//}
	//fmt.Println("任务总数",runtime.NumGoroutine())
	//WG.Wait()
	//fmt.Println("全部执行完成")


//使用WG 计数器，确认协程都执行完毕后再继续执行主线程
url:="https://news.cnblogs.com/n/page/%d/"
var WG sync.WaitGroup
for i:=1;i<=3;i++{
	go func(index int) {
		defer func() {
			if err:=recover();err!=nil{ //捕获异常并处理
				fmt.Println(err)
			}
			WG.Done()
		}()
		url:=fmt.Sprintf(url,index) //格式化字符串，生成每一页页码url
		res,_:=http.Get(url) //http请求
		cnt,_:=ioutil.ReadAll(res.Body) //获取body
		ioutil.WriteFile(fmt.Sprintf("./files/%d",index),cnt,666) //输出文件

	}(i)
	WG.Add(1)
}

	WG.Wait()
	fmt.Println("抓取完成")




}
