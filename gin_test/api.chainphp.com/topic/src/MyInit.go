package src

import (
	"log"
	"os"
	"os/signal"
)

var ServeSignChan chan os.Signal

func init()  {
	ServeSignChan=make(chan os.Signal)
}

func ServerNotify()  {
	signal.Notify(ServeSignChan,os.Interrupt) //只监听Interrupt信号
	<-ServeSignChan
}
func ShuDownServer(err error)  {
	log.Println(err)
	ServeSignChan<-os.Interrupt
}
