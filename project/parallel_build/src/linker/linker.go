package linker

import (
	"bytes"
	"config"
	"fmt"
	"github.com/mahonia"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type ILinker interface {
	Init(context *config.Context) int
	Start() int
	Stop() int
	Wait() int
}

type Linker struct {
	context    *config.Context
	fm         IFileManager
	wait_group *sync.WaitGroup
	ch         chan int
	chNumber   int
	decoder    mahonia.Decoder
}

func (this *Linker) Init(context *config.Context) int {
	this.context = context
	this.fm = &FileManager{}
	this.fm.Init(this.context.WorkDir)
	this.wait_group = &sync.WaitGroup{}
	this.chNumber = this.context.CPUCount
	this.ch = make(chan int, this.chNumber)
	this.decoder = mahonia.NewDecoder("gb18030")

	return 0
}
func (this *Linker) Start() int {
	fmt.Println("link starting...")
	this.fm.Load()
	os.Chdir(this.context.WorkDir)
	fileList := this.fm.GetFileList()
	linkCount := 40 //int(math.Ceil(float64(this.fm.GetFileList().Len())/8.0))
	fileCount := fileList.Len()
	processCount := 0
	startCount := 0
	s := make([]string, 0, linkCount)
	for item := fileList.Front(); item != nil; item = item.Next() {
		s = append(s, item.Value.(string))
		if len(s)%linkCount == 0 {
			buildlist := make([]string, linkCount, linkCount)
			copy(buildlist, s)
			this.ch <- 1
			go this.build(buildlist)
			this.wait_group.Add(1)
			startCount++
			fmt.Printf("%d goroutine process %d started\n", startCount, len(buildlist))
			s = make([]string, 0, linkCount)
		}
		if processCount++; fileCount-processCount < 80 {
			linkCount = 20
		}
	}
	if len(s) > 0 {
		buildlist := make([]string, len(s), linkCount)
		copy(buildlist, s)
		this.ch <- 1
		go this.build(buildlist)
		this.wait_group.Add(1)
		startCount++
		fmt.Printf("%d goroutine process %d started\n", startCount, len(buildlist))
		s = make([]string, 0, linkCount)
	}
	fmt.Println("All goroutines  started")
	return 0
}

func (this *Linker) Stop() int {
	this.fm.UnInit()
	close(this.ch)
	return 0
}

func (this *Linker) Wait() int {
	this.wait_group.Wait()
	fmt.Println("link done")
	return 0
}

func (this *Linker) build(lbmlist []string) {
	defer this.wait_group.Done()
	in := bytes.NewBuffer(nil)
	os.Chdir(this.context.VCDir)
	cmd := exec.Command("cmd", "/K", "vcvarsall.bat\n")
	cmd.Stdin = in
	for _, lbm := range lbmlist {
		var output string
		if len(this.context.OutDir) > 0 {
			output = strings.Replace(filepath.Join(this.context.OutDir, string(lbm[strings.LastIndex(lbm, "\\")+1:])), ".obj", ".dll", 1)
		} else {
			output = strings.Replace(lbm, ".obj", ".dll", 1)
		}
		in.WriteString("link /INCREMENTAL:NO /NOLOGO /DLL /MANIFEST /MANIFESTUAC:\"level='asInvoker' uiAccess='" +
			"false'\"  /SUBSYSTEM:WINDOWS /OPT:REF /OPT:ICF /LTCG /DYNAMICBASE /NXCOMPAT /MACHINE:X86 /ERRORREPORT:PROMPT" +
			"  /LIBPATH:\"" + this.context.LibDir + "\" xsdkDBEngine.lib lbmapi.lib kcxpapi.lib encrypt.lib kcxpxa.lib KCBPPacketOpApi.lib  " +
			"GeneralLBMAPI.lib  odbc1pc.lib KCAS_AuthenticationCheck.lib KSTEncryptd.lib kernel32.lib user32.lib " +
			"gdi32.lib winspool.lib comdlg32.lib advapi32.lib shell32.lib ole32.lib oleaut32.lib uuid.lib odbc32.lib " +
			"odbccp32.lib odbcbcp.lib bkps.lib sett.lib common.lib base.lib " + lbm + " /OUT:" + output + "\n mt.exe -outputresource:"+
			output+";#2 -manifest "+output+".manifest\n")
	}
	out, err := cmd.Output()
	if err == nil {
		fmt.Println(this.decoder.ConvertString(string(out)))
	} else {
		fmt.Println("link error:", this.decoder.ConvertString(err.Error()))
	}
	<-this.ch
}
