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
)

type ISerice interface {
	Init() int
	Run() int
}

type Service struct {
	compiler compiler.ICompiler
	linker   linker.ILinker
	context  *config.Context
}

func (this *Service) Init() int {
	workdir, _ := os.Getwd()
	work_dir := flag.String("WorkDir", workdir, "WorkDir")
	lib_dir := flag.String("LibDir", workdir, "LibDir")
	vc_dir := flag.String("VCDir", workdir, "VCDir")
	out_dir := flag.String("OutDir", workdir, "OutDir")
	flag.Parse()
	this.context = &config.Context{WorkDir: *work_dir,
		LibDir:   *lib_dir,
		VCDir:    *vc_dir,
		OutDir:   *out_dir,
		CPUCount: runtime.NumCPU()}

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
	this.linker.Start2()
	this.linker.Wait()
	this.linker.Stop()
	stopTime = time.Now()
	fmt.Printf("link used time:%d seconds\n", stopTime.Local().Unix()-startTime.Local().Unix())
	return 0
}
