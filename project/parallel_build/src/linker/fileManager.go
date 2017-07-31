package linker
import (
	"fmt"
	"path/filepath"
	"container/list"
	"strings"
	"os"
)
type IFileManager interface {
	Init(work_dir string) int
	GetFileList() *list.List
	Load() int
	ReLoad() int
	UnInit()int
}
type lbm_dir struct{
	base_lbm_dir string
	bkps_lbm_dir string
	sett_lbm_idr string
}

func (lbm *lbm_dir)print() {
	fmt.Println("lbm dirs:")
	fmt.Println(lbm.base_lbm_dir)
	fmt.Println(lbm.bkps_lbm_dir)
	fmt.Println(lbm.sett_lbm_idr)
}

type FileManager struct{
	Work_Dir string
	lbmdir *lbm_dir
	LBM_File *list.List
}

func (fm *FileManager)Init(work_dir string)int{
	fm.Work_Dir=work_dir
	fm.LBM_File=list.New()
	fm.lbmdir=&lbm_dir{base_lbm_dir:filepath.Join(fm.Work_Dir,"src\\base\\lbm\\"),
		bkps_lbm_dir:filepath.Join(fm.Work_Dir,"src\\bkps\\lbm\\"),
		sett_lbm_idr:filepath.Join(fm.Work_Dir,"src\\sett\\lbm\\")}
	fmt.Println("work_dir=",fm.Work_Dir)
	fm.lbmdir.print()
	return 0
}

func (fm *FileManager)GetFileList()*list.List{
	return fm.LBM_File
}

func (fm *FileManager) scanFile(file_dir string){
	err:=filepath.Walk(file_dir,func(path string,f os.FileInfo,err error) error{
		if f==nil{
			return err
		}
		if f.IsDir(){
			return nil
		}
		if(strings.HasSuffix(path,".obj")){
			//fmt.Println("lbm_file:",path)
			fm.LBM_File.PushBack(path)
		}
		return nil
	})
	if(err!=nil){
		fmt.Println("err:",err)
	}
}

func (fm *FileManager) scanLBMFile(){
	fm.scanFile(fm.lbmdir.base_lbm_dir)
	fm.scanFile(fm.lbmdir.bkps_lbm_dir)
	fm.scanFile(fm.lbmdir.sett_lbm_idr)
	fmt.Printf("Scaned %d LBM files\n",fm.LBM_File.Len())
}

func (fm *FileManager)Load()int {
	fmt.Println("FileManager is loading...")
	fm.scanLBMFile()
	fmt.Println("FileManager loaded")
	return  0
}

func (fm *FileManager)ReLoad()int{
	fm.LBM_File=list.New()
	fmt.Println("FileManager is reloading...")
	fm.scanLBMFile()
	fmt.Println("FileManager reloaded")
	return  0
}
func (fm *FileManager)UnInit() int{
	fmt.Println("FileManager is unloading...")
	fm.LBM_File=list.New()
	fmt.Println("FileManager unloaded")
	return  0
}
