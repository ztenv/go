package main

import "shlog"

func main() {
	var logger shlog.ILogger = &shlog.Logger{}
	logger.Init("logdemo.log")
	defer logger.UnInit()
	logger.Debug("%d+%d=%d", 1, 2, 3)
	logger.Info("hello %s", "world")
	logger.Warn("I %s %s", "see", "you")
	logger.Error("do you %s", "known")
	logger.Fatal("%s************%s***********%d", "######", "$$$$$$", 100)
}
