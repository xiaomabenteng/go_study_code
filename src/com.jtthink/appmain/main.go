package main

import (
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//func test() (ret int) { //函数返回值可以只写类型int,或者写上变量名和类型
//	//a:=1
//	return                // 此时可以直接写 return
//	return ret
//}

//func test() (ret int) {  // 返回值是1。因为该函数执行顺序是，ret被初始化为0，然后变量a赋值给ret,返回ret
//	a:=1
//	defer func() {
//		a++
//	}()
//	return a
//}
func test() (ret int) {  // 返回值是2。因为该函数执行顺序是，ret被初始化为0，return先执行把变量a赋值给ret,接着defer执行，最后返回ret
	a:=1
	defer func() {
		ret++
	}()
	return a
}
func main()  {
	fmt.Println(test())
//
//
//url:="https://news.cnblogs.com/n/page/%d/"
//
//c:=make(chan map[int][]byte)
//for i:=1;i<=3;i++{
//
//	go func(index int) {
//		url:=fmt.Sprintf(url,index) //格式化字符串，生成每一页页码url
//		res,_:=http.Get(url) //http请求
//		cnt,_:=ioutil.ReadAll(res.Body) //获取body
//		c<-map[int][]byte{index:cnt} //将序号，和网页内容写入channel
//		if index==3{
//			close(c)
//		}
//	}(i)
//}
//	for getcnt:=range c { //主线程干的事，不断地range channel 直到channel关闭
//		for i,v:=range getcnt{
//			ioutil.WriteFile(fmt.Sprintf("./files/%d",i),v,666) //输出文件
//		}
//	}






}
