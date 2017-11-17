package service
import(
	"os/exec"
	"github.com/mahonia"
	"os"
	"config"
	"bytes"
	"shlog"
)

type Icleaner interface{
	Init(context *config.Context) int
	CleanInterFiles() int
	UnInit() int
}

type cleaner struct {
	context *config.Context
	decoder    mahonia.Decoder
	log shlog.ILogger
}

func (this *cleaner)Init(context *config.Context) int{
	this.context=context
	this.log=this.context.Log
	this.decoder = mahonia.NewDecoder("gb18030")
	return 0
}
func (this *cleaner)UnInit()int{
	this.context=nil
	return 0
}

func (this *cleaner)CleanInterFiles()int{
	this.log.Info("cleaning outdir:%s",this.context.OutDir)
	os.Chdir(this.context.OutDir)
	in:=bytes.NewBuffer(nil)
	cmd:=exec.Command("cmd","/K","del *.lib\n")
	cmd.Stdin=in
	in.WriteString("del *.exp\n del *.dll.manifest\n")
	out,err:=cmd.Output()
	res:=0
	if err!=nil{
		this.log.Error("CleanInterFiles error:%s",this.decoder.ConvertString(err.Error()))
		res=-1
	}else{
		this.log.Info("CleanInterFiles:%s",this.decoder.ConvertString(string(out)))
		res=0
	}
	this.log.Info("Outdir is cleaned")
	return res
}
