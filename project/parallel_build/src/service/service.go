package service

import (
	"compiler"
	"config"
	"flag"
	"linker"
	"os"
	"runtime"
	"time"
	"strings"
	"path/filepath"
	"shlog"
)

type ISerice interface {
	Init(log shlog.ILogger) int
	Run() int
	Clean() int

	parseArgs()(*string,*string,*string,*string,*string)
	buildDir(work_dir *string,lib_dir *string,vc_dir *string,out_dir *string) int
}

type Service struct {
	//copysrc copysrc.ICopysrc
	compiler compiler.ICompiler
	linker   linker.ILinker
	context  *config.Context
	log shlog.ILogger
	clean Icleaner
}

func (this *Service) parseArgs()(*string,*string,*string,*string,*string){
	workdir, _ := os.Getwd()
	work_dir := flag.String("WorkDir", filepath.Join(workdir,"..\\"), "--WorkDir=[WorkDir:absolute path or relative path]")
	lib_dir := flag.String("LibDir","" , "--LibDir=[KBSS LibDir:default will be set :$(KCBP_DIR)\\lib\\]")
	vc_dir := flag.String("VCDir", "", "--VCDir=[vcvarsall.bat file dir:absolute path or relative path]")
	out_dir := flag.String("OutDir", "", "--OutDir=[OutDir:absolute path or relative path.default " +
		"will be set :$(KCBP_DIR)\\kbsslbm\\]")
	compile_all:=flag.String("CompileAll","true","--CompileAll=[true|false]")
	flag.Parse()
	return work_dir,lib_dir,vc_dir,out_dir,compile_all
}

func (this *Service)buildDir(work_dir *string,lib_dir *string,vc_dir *string,out_dir *string)int{
	workdir,_:=os.Getwd()
	if strings.Index(*work_dir,":")==-1{
		*work_dir=filepath.Join(workdir,*work_dir)
	}
	_,err:=os.Stat(*work_dir)
	if err!=nil{
		this.log.Error("open work_dir:%s error:%s,please check the WorkDir again!",*work_dir,err.Error())
		return -1
	}
	if len(*lib_dir)>2 && strings.Index(*lib_dir,":")==-1{
		*lib_dir=filepath.Join(workdir,*lib_dir)
	}
	if strings.Index(*vc_dir,":")==-1{
		*vc_dir=filepath.Join(workdir,*vc_dir)
	}
	_,err=os.Stat(*vc_dir)
	if err!=nil{
		this.log.Error("open vc_dir:%s error:%s",*vc_dir,err.Error())
		return -2
	}

	if len(*out_dir)!=0 {
		if strings.Index(*out_dir, ":") == -1 {
			*out_dir = filepath.Join(workdir, *out_dir)
		}
		_, err = os.Stat(*out_dir)
		if err != nil {
			this.log.Warn("out_dir:%s does not exist,will be creating...", *out_dir)
			err = os.MkdirAll(*out_dir, 777)
			if err == nil {
				this.log.Info("out_dir:%s created successfully!\n", *out_dir)
			} else {
				this.log.Error("out_dir:%s created error:%s\n", *out_dir, err.Error())
				return -3
			}
		}
	}
	return 0
}

func (this *Service) Init(log shlog.ILogger) int {
	var res int=0
	this.log=log
	work_dir,lib_dir,vc_dir,out_dir,compile_all:=this.parseArgs()
	if res=this.buildDir(work_dir,lib_dir,vc_dir,out_dir);res!=0{
		log.Error("buildDir error,please check")
		return res
	}
	this.context = &config.Context{WorkDir: *work_dir,
		LibDir:   *lib_dir,
		VCDir:    *vc_dir,
		OutDir:   *out_dir,
		CPUCount: runtime.NumCPU(),
		IsCompileAll:strings.ToLower(*compile_all),
		Log:log}
	this.context.Print()

	//this.copysrc=new(copysrc.Copysrc)
	//this.copysrc.Init(this.context)
	//this.copysrc.Run()

	this.compiler = new(compiler.Compiler)
	res=this.compiler.Init(this.context)
	this.linker = new(linker.Linker)
	res|=this.linker.Init(this.context)
	return res
}

func (this *Service) Run() int {
	var res int=0
	startTime := time.Now()
	if res=this.compiler.Start();res==0 {
		res |= this.compiler.Wait()
		res |= this.compiler.Stop()
	}
	stopTime := time.Now()
	this.log.Info("compile used time:%d seconds", stopTime.Local().Unix()-startTime.Local().Unix())

	if res==0 {
		startTime = time.Now()
		if res |= this.linker.Start();res==0 {
			res |= this.linker.Wait()
			res |= this.linker.Stop()
		}
		stopTime = time.Now()
		this.log.Info("link used time:%d seconds", stopTime.Local().Unix()-startTime.Local().Unix())
	}

	this.clean=&cleaner{}
	this.clean.Init(this.context)
	return res
}
func (this *Service)Clean()int {
	if this.clean != nil{
		return this.clean.CleanInterFiles()
	}
	return -1
}
