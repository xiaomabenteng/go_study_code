package main

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"sync"
	"time"
)

func GetRqPath(rq string)  string{
	r:=regexp.MustCompile(`^GET\s(.*?)\sHTTP`)
	if r.MatchString(rq) {
		return r.FindStringSubmatch(rq)[1]
	} else{
		return "/"
	}
}
func main()  {

//	str:=`GET /abc.php HTTP/1.1
//Host: 127.0.0.1:8099
//Connection: keep-alive
//Cache-Control: max-age=0
//Upgrade-Insecure-Requests: 1`
	//	fmt.Println(GetRqPath(str))


	rand.Seed(time.Now().Unix()) //设置随机因子。默认值是1，不设置的话，每次程序重新运行，生成的随机数是一样的

	//模拟需求，倒计时两秒内，生成一些幸运用户ID
	ctx,_:=context.WithTimeout(context.Background(),time.Second*2) //主线程设置2秒后超时取消的上下文context

	var WG sync.WaitGroup
	WG.Add(1)

	go GetUsers(ctx,&WG)
	WG.Wait()

	fmt.Println("生成幸运用户成功")

}

func GetUsers(ctx context.Context,wg *sync.WaitGroup)  {
	fmt.Println("开始生成幸运用户")
	users:=make([]int,0)

	for  {
		select{
			case <-ctx.Done(): //此chan可读，代表父context发起 取消操作
				fmt.Println(users)
				wg.Done()
				return
			default:
				users=append(users,getUserId(1000,100000))
		}
	}

}

func getUserId(min int,max int)  int{
	return rand.Intn(max-min)+min
}
