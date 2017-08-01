package compiler

import (
	"bufio"
	"bytes"
	"config"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type ICompiler interface {
	Init(context *config.Context) int
	Start() int
	Wait() int
	Stop() int
}

type Compiler struct {
	context   *config.Context
	waitGroup *sync.WaitGroup
	decoder   mahonia.Decoder
}

func (this *Compiler) formatMakefile() int {
	if file, err := os.Open(filepath.Join(this.context.WorkDir, "src\\makefile")); err == nil {
		defer file.Close()
		if wr, werr := os.Create(filepath.Join(this.context.WorkDir, "src\\makefile.new")); werr == nil {
			defer wr.Close()
			reader := bufio.NewReader(file)
			for {
				line, _, err := reader.ReadLine()
				if err != nil || io.EOF == err {
					break
				}
				linestr := string(strings.Replace(string(line), "cl ", "cl /MP"+strconv.Itoa(runtime.NumCPU())+" ", -1))
				if strings.Index(linestr, "SRC_DIR") != -1 && strings.Index(linestr, "lbm") != -1 && strings.LastIndex(linestr, "base.lib") != -1 {
					break
				}
				wr.WriteString(linestr + "\n")
			}
		} else {
			fmt.Printf("create file:%s error:%s\n", filepath.Join(this.context.WorkDir, "src\\makefile.new"), werr.Error())
			return -1
		}
	} else {
		fmt.Printf("open file:%s error:%s\n", filepath.Join(this.context.WorkDir, "src\\makefile"), err.Error())
		return -1
	}
	return 0
}

func (this *Compiler) processMakefile() int {
	os.Chdir(filepath.Join(this.context.WorkDir, "src"))
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("cmd", "/K", "gen_makefile_new.bat")
	cmd.Stdin = in
	out, err := cmd.Output()
	if err == nil {
		fmt.Printf("create makefile :%s", this.decoder.ConvertString(string(out)))
	} else {
		fmt.Printf("create makefile error:%s", this.decoder.ConvertString(err.Error()))
	}
	this.formatMakefile()
	os.Remove(filepath.Join(this.context.WorkDir, "src\\makefile"))
	os.Rename(filepath.Join(this.context.WorkDir, "src\\makefile.new"), filepath.Join(this.context.WorkDir, "src\\makefile"))
	in.Reset()
	cmd = exec.Command("cmd", "/K", "del_temp_file.bat")
	cmd.Stdin = in
	in.WriteString("\\cp_makefile.bat\n")
	out, err = cmd.Output()

	if err == nil {
		fmt.Printf("copy makefile:%s", this.decoder.ConvertString(string(out)))
	} else {
		fmt.Printf("copy makefile error:%s", this.decoder.ConvertString(err.Error()))
		return -1
	}
	return 0
}

func (this *Compiler) Init(context *config.Context) int {
	this.context = context
	this.waitGroup = new(sync.WaitGroup)
	this.decoder = mahonia.NewDecoder("gb18030")
	this.processMakefile()
	return 0
}

func (this *Compiler) compile() int {

	defer this.waitGroup.Done()

	os.Chdir(filepath.Join(this.context.WorkDir, "src"))
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("cmd", "/K", filepath.Join(this.context.VCDir, "vcvarsall.bat"))
	cmd.Stdin = in
	in.WriteString("make\n")

	out, err := cmd.Output()
	if err == nil {
		fmt.Printf("compile info:%s", this.decoder.ConvertString(string(out)))
	} else {
		fmt.Printf("compile error:%s", this.decoder.ConvertString(err.Error()))
	}

	return 0
}

func (this *Compiler) Start() int {
	fmt.Println("Compiler is starting....")
	this.waitGroup.Add(1)
	go this.compile()
	fmt.Println("Compiler started")
	return 0
}

func (this *Compiler) Wait() int {
	fmt.Println("Compiler is waiting...")
	this.waitGroup.Wait()
	fmt.Println("Compiler waited")
	return 0
}

func (this *Compiler) Stop() int {
	fmt.Println("Compiler is stopping...")
	fmt.Println("Compiler stopped")
	return 0
}
