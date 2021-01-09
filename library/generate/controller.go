package generate

import (
	"flag"
	"log"
	"os"
	"reTools/library/generate/astparser"
	"reTools/library/generate/tplparser"
	"reTools/library/go-fs"
)

func init() {

}

type ControllerCommand struct {
	IsCreate	*bool	//是否创建代码，以后可能还会有删除
	IFile		*string //接口文件
	Dir			*string //目标文件夹
	CommandSet
	ControllerCommandSet *flag.FlagSet
}

func NewControllerCmd() *ControllerCommand {
	is_create := false

	return &ControllerCommand{IsCreate: &is_create, //由于IsCreate是指针，需要赋值，不然后面判断会出错
		ControllerCommandSet: flag.NewFlagSet("controller args", flag.ExitOnError)}
}

func (this *ControllerCommand) Init() {
	// go run main.go service -c
	if len(os.Args) > 2 && os.Args[1] == "controller" {
		this.IsCreate = this.ControllerCommandSet.Bool("c", false, "create controller")
		this.IFile = this.ControllerCommandSet.String("i", "", "interface file (app\\api)")
		this.Dir = this.ControllerCommandSet.String("d", "app/controller",
			"generates to the specified folder (default:app\\service)")
		err := this.ControllerCommandSet.Parse(os.Args[:2])

		if err != nil {
			log.Println(err)
		}
	}
}

func (this *ControllerCommand) Run() {
	if *this.IsCreate { // -c 有
		if err := fs.FileExists(*this.IFile + ".go"); err == nil {
			p := astparser.NewParser(*this.IFile + ".go")
			infs := p.ParseInterfaces() // 切片
			tplParser := tplparser.NewControllerParser() //模板解析类，专门处理 controller 解析和生成

			for _, pi := range infs {
				tplParser.Parse()
			}
		}
	}
}
