package main

import (
	"fmt"
	"net/http"
	"time"
)
type MyHandler struct {

}

func (*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	fmt.Println(request.RequestURI)
	if request.RequestURI=="/woshi"{
		writer.Write([]byte("zhangsan"))
	}else{
		writer.Write([]byte("hello myhandler"))
	}

}

func main()  {


	mymux:=http.NewServeMux()                // 自定义路由管理
	mymux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {  //
		writer.Write([]byte("index"))
	})
	mymux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {  //

		getUserName:=request.URL.Query().Get("username")
		if getUserName!=""{
			 c:=&http.Cookie{Name:"username",Value:getUserName}
			 http.SetCookie(writer,c)
		}
		writer.Write([]byte("登录页"))
	})
	mymux.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {  //
		c:=&http.Cookie{Name:"username",Value:"",Expires:time.Now().AddDate(-1,0,0)} //设置cookie过期
		http.SetCookie(writer,c)

		writer.Header().Set("Content-type","text/html")
		writer.Write([]byte("unlogin....."))
		script:=`<script>
		setTimeout(()=>{self.location='/'},2000)
</script>`
		writer.Write([]byte(script))


	})
	server:=http.Server{Addr:":8099",Handler:mymux}
	err:=server.ListenAndServe()
	if err!=nil{
		fmt.Println(err)
	}
}
