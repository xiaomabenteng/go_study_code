package main

import (
	"fmt"
	"regexp"
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

	str:=`GET /abc.php HTTP/1.1
Host: 127.0.0.1:8099
Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1`

	fmt.Println(GetRqPath(str))


}
