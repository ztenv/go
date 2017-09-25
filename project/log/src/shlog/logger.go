package shlog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type ILogger interface {
	Init(logfile string) int
	UnInit()
	Debug(format string, msg ...interface{})
	Info(format string, msg ...interface{})
	Warn(format string, msg ...interface{})
	Error(format string, msg ...interface{})
	Fatal(format string, msg ...interface{})
}
type Logger struct {
	fileName string
	logFile  *os.File

	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

func (this *Logger) Init(logfile string) int {
	this.fileName = logfile
	lf, _err := os.Create(this.fileName)
	if _err != nil {
		fmt.Printf("create log file :%s error:", this.fileName, _err.Error())
		return -1
	}
	this.logFile = lf
	writers := []io.Writer{os.Stdout, lf}
	mw := io.MultiWriter(writers...)
	this.debugLogger = log.New(mw, "[debug]", log.Ldate|log.Lmicroseconds)
	this.infoLogger = log.New(mw, "[info]", log.Ldate|log.Lmicroseconds)
	this.warnLogger = log.New(mw, "[warn]", log.Ldate|log.Lmicroseconds)
	this.errorLogger = log.New(mw, "[error]", log.Ldate|log.Lmicroseconds)
	this.fatalLogger = log.New(mw, "[fatal]", log.Ldate|log.Lmicroseconds)
	return 0
}

func (this *Logger) UnInit() {
	_err := this.logFile.Close()
	if _err != nil {
		fmt.Println("Close file:%s error:%s", this.fileName, _err.Error())
	}
}

func (this *Logger) Debug(format string, msg ...interface{}) {
	if msg != nil {
		this.debugLogger.Printf(format+"\n", msg...)
	} else {
		this.debugLogger.Println(format)
	}
}
func (this *Logger) Info(format string, msg ...interface{}) {
	if msg != nil {
		this.infoLogger.Printf(format+"\n", msg...)
	} else {
		this.infoLogger.Println(format)
	}
}
func (this *Logger) Warn(format string, msg ...interface{}) {
	if msg != nil {
		this.warnLogger.Printf(format+"\n", msg...)
	} else {
		this.warnLogger.Println(format)
	}
}
func (this *Logger) Error(format string, msg ...interface{}) {
	if msg != nil {
		this.errorLogger.Printf(format+"\n", msg...)
	} else {
		this.errorLogger.Println(format)
	}
}
func (this *Logger) Fatal(format string, msg ...interface{}) {
	if msg != nil {
		this.fatalLogger.Printf(format+"\n", msg...)
	} else {
		this.fatalLogger.Println(format)
	}
}
