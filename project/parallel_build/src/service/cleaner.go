package service
import(
	"fmt"
	"os/exec"
	"github.com/mahonia"
	"config"
	"os"
	"bytes"
)

type Icleaner interface{
	Init(context *config.Context) int
	CleanInterFiles() int
	UnInit() int
}

type cleaner struct {
	context *config.Context
	decoder    mahonia.Decoder
}

func (this *cleaner)Init(context *config.Context) int{
	this.context=context
	this.decoder = mahonia.NewDecoder("gb18030")
	return 0
}
func (this *cleaner)UnInit()int{
	this.context=nil
	return 0
}

func (this *cleaner)CleanInterFiles()int{
	os.Chdir(this.context.OutDir)
	in:=bytes.NewBuffer(nil)
	cmd:=exec.Command("cmd","/K","del *.lib\n")
	cmd.Stdin=in
	in.WriteString("del *.exp\n del *.dll.manifest\n")
	out,err:=cmd.Output()
	if err!=nil{
		fmt.Printf("CleanInterFiles error:%s",this.decoder.ConvertString(err.Error()))
		return -1
	}else{
		fmt.Printf("CleanInterFiles:%s",this.decoder.ConvertString(string(out)))
	}
	return 0
}
