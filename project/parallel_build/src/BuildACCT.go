package main

import (
	"fmt"
	"service"
	"time"
	"shlog"
	"os"
)
var _VERSION_ string="0.0.0.0"
func main() {
	log,err:=initLog()
	if err!=nil{
		fmt.Printf("init log error:%s",err.Error())
		os.Exit(1)
	}
	var res int=0
	defer log.UnInit()
	log.Info("Version:=%s",_VERSION_)
	startTime:= time.Now()
	log.Info("Building starts at:%s",startTime.Format("2006-01-02 15:04:05"))
	{
		buildService:= &service.Service{}
		defer buildService.Clean()//出错不清理生成的.lib文件
		if res=buildService.Init(log);res != 0 {
			log.Fatal("Service init error,please check your config")
		}else {
			res = buildService.Run()
			stopTime := time.Now()
			log.Info("Building stop at:%s", stopTime.Format("2006-01-02 15:04:05"))
			log.Info("Building used time:%d seconds", stopTime.Local().Unix()-startTime.Local().Unix())
		}
	}
	if(res!=0) {
		log.UnInit()//出错，则显示调用日志卸载，否则golang将不会调用
		os.Exit(res)
	}
}

func initLog() (shlog.ILogger,error){
	var loger shlog.ILogger=&shlog.Logger{}
	res:=loger.Init("BuildACCT"+time.Now().Format("20060102150405")+".log")
	return loger,res
}
