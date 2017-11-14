package linker

import (
	"container/list"
	"os"
	"path/filepath"
	"strings"
	"config"
	"shlog"
)

type IFileManager interface {
	Init(context *config.Context) int
	GetFileList() *list.List
	GetLbmDir() *Lbm_dir
	Load() int
	ReLoad() int
	UnInit() int
}
type Lbm_dir struct {
	base_lbm_dir string
	bkps_lbm_dir string
	sett_lbm_idr string
}

type FileManager struct {
	context *config.Context
	Work_Dir string
	lbmdir   *Lbm_dir
	LBM_File *list.List
	log shlog.ILogger
}

func (fm *FileManager) Init(context *config.Context) int {
	fm.context=context
	fm.log=fm.context.Log
	fm.Work_Dir = fm.context.WorkDir
	fm.LBM_File = list.New()
	fm.lbmdir = &Lbm_dir{base_lbm_dir: filepath.Join(fm.Work_Dir, "src\\base\\lbm\\"),
		bkps_lbm_dir: filepath.Join(fm.Work_Dir, "src\\bkps\\lbm\\"),
		sett_lbm_idr: filepath.Join(fm.Work_Dir, "src\\sett\\lbm\\")}
	fm.log.Info("work_dir=%s", fm.Work_Dir)

	fm.log.Info("lbm dirs:")
	fm.log.Info(fm.lbmdir.base_lbm_dir)
	fm.log.Info(fm.lbmdir.bkps_lbm_dir)
	fm.log.Info(fm.lbmdir.sett_lbm_idr)
	return 0
}

func(fm *FileManager) GetLbmDir() *Lbm_dir{
	return fm.lbmdir
}

func (fm *FileManager) GetFileList() *list.List {
	return fm.LBM_File
}

func (fm *FileManager) scanFile(file_dir string) {
	err := filepath.Walk(file_dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".cpp") {
			//fmt.Println("lbm_file:",path)
			path=strings.Replace(path,".cpp",".obj",1)
			fm.LBM_File.PushBack(path)
		}
		return nil
	})
	if err != nil {
		fm.log.Error("scanFile err:%s", err.Error())
	}
}

func (fm *FileManager) scanLBMFile() int{
	fm.scanFile(fm.lbmdir.base_lbm_dir)
	fm.scanFile(fm.lbmdir.bkps_lbm_dir)
	fm.scanFile(fm.lbmdir.sett_lbm_idr)
	fm.log.Info("Scaned %d LBM files", fm.LBM_File.Len())
	return fm.LBM_File.Len()
}

func (fm *FileManager) Load() int {
	fm.log.Info("FileManager is loading...")
	res:=fm.scanLBMFile()
	fm.log.Info("FileManager loaded")
	if res>0{
		return 0
	}else{
		return -1
	}
}

func (fm *FileManager) ReLoad() int {
	fm.LBM_File = list.New()
	fm.log.Info("FileManager is reloading...")
	res:=fm.scanLBMFile()
	fm.log.Info("FileManager reloaded")
	if res>0{
		return 0
	}else{
		return -1
	}
}
func (fm *FileManager) UnInit() int {
	fm.log.Info("FileManager is unloading...")
	fm.LBM_File = list.New()
	fm.log.Info("FileManager unloaded")
	return 0
}
