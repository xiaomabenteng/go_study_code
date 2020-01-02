package main

import (
	"fmt"
	"math/rand"
	"regexp"
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
//
//	fmt.Println(GetRqPath(str))
	rand.Seed(time.Now().Unix()) //设置随机因子。默认值是1，不设置的话，每次程序重新运行，生成的随机数是一样的
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))


}
