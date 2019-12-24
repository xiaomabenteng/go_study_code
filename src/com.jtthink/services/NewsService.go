package services

import "com.jtthink/core"

type NewsService struct {
	
}

func (ns NewsService) Get(id int) string {
	return "单新闻"
}

func init()  {

	core.SetService(UserService{})
}