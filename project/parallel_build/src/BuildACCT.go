package main

import (
	"fmt"
	"service"
	"time"
)
var _VERSION_ string="0.0.0.0"
func main() {
	fmt.Println("verion:="+_VERSION_)
	startTime := time.Now()
	fmt.Printf("Building start at:%s\n", startTime.Format("2006-01-02 15:04:05"))
	{
		srv := &service.Service{}
		defer srv.Clean()
		if srv.Init() != 0 {
			fmt.Println("service Init error,please check your config");
		}
		srv.Run()
		stopTime := time.Now()
		fmt.Printf("Building stop at:%s\n", stopTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("time used:%d seconds\n", stopTime.Local().Unix()-startTime.Local().Unix())
	}
	fmt.Println("verion:="+_VERSION_)
}
