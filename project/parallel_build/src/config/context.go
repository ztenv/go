package config

import (
	"shlog"
)

type Context struct {
	WorkDir  string
	LibDir   string
	VCDir    string
	Platform string
	OutDir   string
	CPUCount int
	IsCompileAll string

	Log shlog.ILogger
}

func (this *Context) Print() {
	this.Log.Info("Workdir:%s", this.WorkDir)
	if(len(this.LibDir)==0) {
		this.Log.Warn("LibDir:LibDir is null and will be as \"KCBP_DIR\\lib\\\",KCBP_DIR is in makefile_template file.")
	}else{
		this.Log.Info("LibDir:%s", this.LibDir)
	}
	this.Log.Info("VCDir:%s", this.VCDir)
	if(len(this.OutDir)==0){
		this.Log.Info("OutDir:OutDir is null and will be set as \"KCBP_DIR\\kbsslbm\\\",KCBP_DIR is in makefile_template file.")
	}else {
		this.Log.Info("OutDir:%s", this.OutDir)
	}
	this.Log.Info("Platform:%s",this.Platform)
	this.Log.Info("CPUCount:%d", this.CPUCount)
	this.Log.Info("IsCompileAll:%s",this.IsCompileAll)
}
