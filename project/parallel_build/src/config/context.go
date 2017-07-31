package config

import(
	"fmt"
)

type Context struct {
	WorkDir string
	LibDir string
	VCDir string
	OutDir string
	CPUCount int
}

func (this *Context) Print()  {
	fmt.Printf("Workdir:%s\n",this.WorkDir)
	fmt.Printf("LibDir:%s\n",this.LibDir)
	fmt.Printf("VCDir:%s\n",this.VCDir)
	fmt.Printf("OutDir%s\n",this.OutDir)
	fmt.Printf("CPUCount:%d\n",this.CPUCount)
}
