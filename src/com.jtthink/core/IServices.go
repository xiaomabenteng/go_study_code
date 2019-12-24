package core

type IServices interface {
	Get(id int) string
}

var myService IServices

func SetService(service IServices)   {
	myService=service
}

func GetService() IServices {
	return myService
}