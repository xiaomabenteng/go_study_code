package servicesb

import "com.jtthink/core"

type UserService struct {
	
}

func (us UserService) Get(id int )  string{
	return "b 单用户"
}

func init()  {

	core.SetService(UserService{})
}
