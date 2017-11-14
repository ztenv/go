package config

import (
	"shlog"
)

type Context struct {
	WorkDir  string
	LibDir   string
	VCDir    string
	OutDir   string
	CPUCount int
	IsCompileAll string

	Log shlog.ILogger
}

func (this *Context) Print() {
	this.Log.Info("Workdir:%s", this.WorkDir)
	if(len(this.LibDir)==0) {
		this.Log.Warn("LibDir:LibDir is null and will be read from lbm\\solution\\src\\makefile_template")
	}else{
		this.Log.Info("LibDir:%s", this.LibDir)
	}
	this.Log.Info("VCDir:%s", this.VCDir)
	this.Log.Info("OutDir:%s", this.OutDir)
	this.Log.Info("CPUCount:%d", this.CPUCount)
	this.Log.Info("IsCompileAll:%s",this.IsCompileAll)
}
