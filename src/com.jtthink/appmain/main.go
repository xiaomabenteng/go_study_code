package main

import (
	"bufio"
	_ "com.jtthink/Servicesb"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)


func main()  {

	////定义一个变量，并输出到控制台
	//var myname bytes.Buffer
	//myname.Write([]byte("zhangsan"))
	//myname.WriteTo(os.Stdout)

	////写入文件
	//file,_:=os.OpenFile("./test/test",os.O_RDWR | os.O_CREATE | os.O_APPEND,666 )
	//defer file.Close()
	//file.Write([]byte("zhangsan"))

	////每隔三秒写入文件。频繁的对磁盘读写会造成不必要的性能开支和对磁盘影响
	//file,_:=os.OpenFile("./test/test",os.O_RDWR | os.O_CREATE | os.O_APPEND | os.O_TRUNC,666 )
	//defer file.Close()
	//for i:=1;i<10;i++ {
	//	file.Write([]byte("zhangsan\n"))
	//	time.Sleep(time.Second*2)
	//}

	//每隔三秒写入文件。
	file,_:=os.OpenFile("./test/test",os.O_RDWR | os.O_CREATE | os.O_APPEND | os.O_TRUNC,666 )
	defer file.Close()
	fw:=bufio.NewWriter(file) //使用bufio包，自带缓冲区默认4096，当达到一定量才进行下一步操作
	for i:=1;i<3;i++ {
		fw.Write([]byte("zhangsan\n"))
		time.Sleep(time.Second*2)
	}
	fw.Flush() //如果一直未达到设置缓冲大小。手动刷新缓冲，把内容写入

}
