package copysrc

import(
	"config"
	"os/exec"
	"path/filepath"
	"github.com/mahonia"
	"os"
	"logger"
)

type ICopysrc interface {
	Init(context *config.Context) int
	Run() int
}

type Copysrc struct{
	context *config.Context
	decoder mahonia.Decoder

	log shlog.ILogger
}

func (this *Copysrc)Init(context *config.Context) int{
	this.context=context
	this.log=this.context.Log
	this.decoder=mahonia.NewDecoder("gb18030")
	return 0
}

func (this *Copysrc)Run() int{
	var cmdstr string="get_src_sql.bat"
	os.Chdir(filepath.Join(this.context.WorkDir,"..\\..\\"))
	cmd:=exec.Command("cmd","/K",cmdstr)

	out,err:=cmd.Output()
	if err==nil{
		this.log.Info("copysrc files file:%s",this.decoder.ConvertString(string(out)))
	}else{
		this.log.Error("copy src files(execute get_src_sql.bat) error:%s",err.Error())
	}
	os.Chdir(this.context.WorkDir)
	return 0
}


