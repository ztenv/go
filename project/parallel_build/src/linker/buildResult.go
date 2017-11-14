package linker

import (
	"config"
	"shlog"
	"path/filepath"
	"os"
	"strings"
)
//主要功能：扫描LBM的cpp、obj、dll文件，并生成一个列表，把列表输出到一个结果表中。
// 以后完成

type LbmNode struct{
	LbmName string
	LbmCpp string
	LbmObj string
	LbmDll string
}
func (this *LbmNode)SetCppName(cpp string){
	this.LbmCpp=cpp
}
func (this *LbmNode)SetObjName(obj string){
	this.LbmObj=obj
}
func (this *LbmNode)SetDllName(dll string){
	this.LbmDll=dll
}

type BuildResult struct{
	context *config.Context
	lbmMap map[string]*LbmNode

	log shlog.ILogger
}

func (this *BuildResult)Init(context *config.Context){
	this.context=context
	this.log=this.context.Log
}

func (this *BuildResult)addLbmNode(lbmName string){
	this.lbmMap[lbmName]=&LbmNode{LbmName:lbmName,LbmCpp:"",LbmObj:"",LbmDll:""}
}
func (this *BuildResult)getLbmNode(lbmName string) *LbmNode{
	node:=this.lbmMap[lbmName]
	if node==nil{
		this.addLbmNode(lbmName)
		node=this.lbmMap[lbmName]
	}
	return node
}
func (this *BuildResult)scanFile(file_dir string,suffix string){
	err := filepath.Walk(file_dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".cpp") {
			path=strings.Replace(path,".cpp",".obj",1)
		}
		return nil
	})
	if err != nil {
		this.log.Error("scanFile err:%s", err.Error())
	}
}
func (this *BuildResult)Build(fm FileManager){
	fl:=fm.GetFileList()
	for item:=fl.Front();item!=nil;item=item.Next(){
		this.lbmMap[item.Value.(string)]=&LbmNode{LbmName:item.Value.(string),LbmCpp:"",LbmObj:"",LbmDll:""}
	}
}

