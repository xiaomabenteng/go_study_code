package main

import (
	"context"
	"net/http"
	"time"
)
/**
模拟一个需求，在url有query请求参数时候，就进行一个缓慢的查询操作。通过context超时时间，超过时间自动报超时

 */
func getCountData(c chan string)  {
	time.Sleep(time.Second*2)
	c<-"统计结果" //向通道内写入数据
}
func main()  {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("query")=="" {
			writer.Write([]byte("这是首页"))
		}else{
			ctx,cancel:=context.WithTimeout(request.Context(),time.Second*3) //三秒后向ctx.Done里面写入数据
			defer cancel()
			c:=make(chan string)
			go getCountData(c)
			select {
				case <-ctx.Done():
					writer.Write([]byte("超时"))
				case ret:=<-c:
					writer.Write([]byte(ret))
			}

		}

	})
	http.ListenAndServe(":8099",nil)

}
