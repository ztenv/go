package main

import (
	"fmt"
	"service"
	"time"
	"shlog"
)
var _VERSION_ string="0.0.0.0"
func main() {
	log,err:=initLog()
	if err!=nil{
		fmt.Printf("init log error:%s",err.Error())
		return
	}
	defer log.UnInit()
	log.Info("Version:=%s",_VERSION_)
	startTime := time.Now()
	log.Info("Building start at:%s",startTime.Format("2006-01-02 15:04:05"))
	{
		srv := &service.Service{}
		defer srv.Clean()
		if srv.Init(log) != 0 {
			log.Fatal("Service init error,please check your config")
		}
		srv.Run()
		stopTime := time.Now()
		log.Info("Building stop at:%s",stopTime.Format("2006-01-02 15:04:05"))
		log.Info("Build time used:%d seconds",stopTime.Local().Unix()-startTime.Local().Unix())
	}
}

func initLog() (shlog.ILogger,error){
	var loger shlog.ILogger=&shlog.Logger{}
	res:=loger.Init("BuildACCT"+time.Now().Format("20060102150405")+".log")
	return loger,res
}
