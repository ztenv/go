package service

import (
	"compiler"
	"config"
	"flag"
	"fmt"
	"linker"
	"os"
	"runtime"
	"time"
	"strings"
	"path/filepath"
)

type ISerice interface {
	Init() int
	Run() int
	Clean() int

	parseArgs()(*string,*string,*string,*string,*string)
	buildDir(work_dir *string,lib_dir *string,vc_dir *string,out_dir *string) int
}

type Service struct {
	compiler compiler.ICompiler
	linker   linker.ILinker
	context  *config.Context
	clean Icleaner
}

func (this *Service) parseArgs()(*string,*string,*string,*string,*string){
	workdir, _ := os.Getwd()
	work_dir := flag.String("WorkDir", filepath.Join(workdir,"..\\"), "--WorkDir=[WorkDir:absolute path or relative path]")
	lib_dir := flag.String("LibDir","" , "--LibDir=[KBSS LibDir:default from the makefile's KCBP_DIR]")
	vc_dir := flag.String("VCDir", "", "--VCDir=[vcvarsall.bat file dir:absolute path or relative path]")
	out_dir := flag.String("OutDir", "", "--OutDir=[OutDir:absolute path or relative path]")
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
		fmt.Printf("work_dir:%s does not exists,please check the WorkDir again!\n",*work_dir)
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
		fmt.Printf("vc_dir:%s does not exist,please check the VCDir again!\n",*vc_dir)
		return -2
	}

	if len(*out_dir)!=0 {
		if strings.Index(*out_dir, ":") == -1 {
			*out_dir = filepath.Join(workdir, *out_dir)
		}
		_, err = os.Stat(*out_dir)
		if err != nil {
			fmt.Printf("out_dir:%s does not exist,will be creating...\n", *out_dir)
			err = os.MkdirAll(*out_dir, 777)
			if err == nil {
				fmt.Printf("out_dir:%s created successfully!\n", *out_dir)
			} else {
				fmt.Printf("out_dir:%s created error:%s\n", *out_dir, err.Error())
				return -3
			}
		}
	}
	return 0
}

func (this *Service) Init() int {
	work_dir,lib_dir,vc_dir,out_dir,compile_all:=this.parseArgs()
	if res:=this.buildDir(work_dir,lib_dir,vc_dir,out_dir);res!=0{
		return res
	}
	this.context = &config.Context{WorkDir: *work_dir,
		LibDir:   *lib_dir,
		VCDir:    *vc_dir,
		OutDir:   *out_dir,
		CPUCount: runtime.NumCPU(),
		IsCompileAll:strings.ToLower(*compile_all)}
	this.context.Print()

	this.compiler = new(compiler.Compiler)
	this.compiler.Init(this.context)
	this.linker = new(linker.Linker)
	this.linker.Init(this.context)
	return 0
}

func (this *Service) Run() int {

	startTime := time.Now()
	this.compiler.Start()
	this.compiler.Wait()
	this.compiler.Stop()
	stopTime := time.Now()
	fmt.Printf("compile used time:%d seconds\n", stopTime.Local().Unix()-startTime.Local().Unix())

	startTime = time.Now()
	this.linker.Start()
	this.linker.Wait()
	this.linker.Stop()
	stopTime = time.Now()
	fmt.Printf("link used time:%d seconds\n", stopTime.Local().Unix()-startTime.Local().Unix())

	this.clean=&cleaner{}
	this.clean.Init(this.context)
	return 0
}
func (this *Service)Clean()int{
	return this.clean.CleanInterFiles()
}
