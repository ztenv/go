package main

import (
	"fmt"
	"service"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Printf("Building start at:%s\n", startTime.Format("2006-01-02 15:04:05"))

	service := &service.Service{}
	service.Init()
	service.Run()

	stopTime := time.Now()
	fmt.Printf("Building stop at:%s\n", stopTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("time used:%d seconds\n", stopTime.Local().Unix()-startTime.Local().Unix())
}
