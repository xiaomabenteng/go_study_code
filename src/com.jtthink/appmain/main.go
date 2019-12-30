package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"time"
)

func test() {
	panic("exp")
}

func main()  {

	//c:=time.After(time.Second*3) // 等待指定时间后，向返回的chan里面写入当前时间。此函数不阻塞，他返回值只一个只读channel
	//fmt.Println("不阻塞")
	//fmt.Println(<-c) //三秒后出现结果
	////定义只读的channel c:=make(<-chan int)
	////定义只写的channel c:make(chan<- int)

	//i:=0
	//abc:fmt.Println(i)   //abc是标签 标签语法
	//time.Sleep(time.Second*2)
	//i++
	//goto abc            //使用goto语句定义的一个死循环




url:="https://news.cnblogs.com/n/page/%d/"

c:=make(chan map[int][]byte)
for i:=1;i<=10;i++{

	go func(index int) {
		defer func() {
			if err:=recover();err!=nil{ //捕获异常并处理
				fmt.Println(err)
			}
		}()
		url:=fmt.Sprintf(url,index) //格式化字符串，生成每一页页码url
		res,_:=http.Get(url) //http请求
		cnt,_:=ioutil.ReadAll(res.Body) //获取body
		c<-map[int][]byte{index:cnt} //将序号，和网页内容写入channel

	}(i)
}
	//for getcnt:=range c { //主线程干的事，不断地range channel 直到channel关闭
	//	for i,v:=range getcnt{
	//		ioutil.WriteFile(fmt.Sprintf("./files/%d",i),v,666) //输出文件
	//	}
	//}

	result:=map[int][]byte{}
	myloop:for{                //标签语法
		select {   			// select 类似switch case。只是用于channel通信（要么send,要么receive 就是读取通道或者写入通道）
			case result=<-c:  //能从通道读取出数据就走这个case
				for i,v:=range result{
					ioutil.WriteFile(fmt.Sprintf("./files/%d",i),v,666) //输出文件
				}
			case <-time.After(time.Second*3):
				break myloop
		}
	}

	//select 很类似switch case .只不过用于channel通信（要么send要么receive）
	//	譬如select {
	//	case i<-c:
	//		xxxoo
	//	case c<-123:
	//		xxxoo
	//	default:
	//	}
	//	1、按顺序判断，如果只有一个case通过，则执行该case
	//	2、如果多个case都通过，则随机选一个执行
	//	3、如果都没通过，则查找default；如果没有default，则阻塞





}
