package main

import "regexp"

func GetRequestPath(rq string)  string{
	r:=regexp.MustCompile(`^GET\s(.*?)\sHTTP`)
	if r.MatchString(rq) {
		return r.FindStringSubmatch(rq)[1]
	} else{
		return "/"
	}
}