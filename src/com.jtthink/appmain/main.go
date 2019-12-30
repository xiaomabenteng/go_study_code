package main

import (
	"bufio"
	_ "com.jtthink/Servicesb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"sync"
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

	////每隔三秒写入文件。
	//file,_:=os.OpenFile("./test/test",os.O_RDWR | os.O_CREATE | os.O_APPEND | os.O_TRUNC,666 )
	//defer file.Close()
	//fw:=bufio.NewWriter(file) //使用bufio包，自带缓冲区默认4096，当达到一定量才进行下一步操作
	//for i:=1;i<3;i++ {
	//	fw.Write([]byte("zhangsan\n"))
	//	time.Sleep(time.Second*2)
	//}
	//fw.Flush() //如果一直未达到设置缓冲大小。手动刷新缓冲，把内容写入


	//多协程，按顺序，按行读取文件
	file,_:=os.OpenFile("./test/test",os.O_RDWR,666 )
	defer file.Close()
	fw:=bufio.NewReader(file)
	var WG sync.WaitGroup
	var mutex sync.Mutex
	for i:=1;i<=2;i++{    //两个协程循环读取文件，协程运行顺序不确定，所以读取出来文件顺序不确定。加互斥锁可解决，
		go func(index int) {
			defer WG.Done()
			for  {
				mutex.Lock()
				str,err:=fw.ReadString('\n')
				if err!=nil{
					if err==io.EOF{
						mutex.Unlock() //文件读取完也要解锁。否则死锁报错
						break
					}
					fmt.Println(err)
				}
				time.Sleep(time.Millisecond*200)
				fmt.Printf("【协程%d】:%s",index,str)
				mutex.Unlock()
			}
		}(i)
	}
	WG.Add(2)
	WG.Wait()
	fmt.Println("读取完成")




}
