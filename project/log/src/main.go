package main

import "logger"

func main() {
	var logger logger.ILogger = &logger.Logger{}
	logger.Init("logdemo.log")
	defer logger.UnInit()
	logger.Debug("aaaa1")
	logger.Info("aaaa2")
	logger.Warn("aaaa3")
	logger.Error("bbbbb:%s", "*****")
	logger.Fatal("aaaaaaafff")
	logger.Debug("aaaa4")
}
