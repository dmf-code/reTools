package tplparser

import (
	"html/template"
	"reTools/library/generate/astparser"
	"reTools/library/go-fs"
)

type ControllerParser struct {
	TplContent	string // 模板内容
}

func NewControllerParser() *ControllerParser {
	data, err := fs.Read("D:\\go\\reTools\\library\\generate\\template\\controller.tpl")

	if err != nil {
		panic(err)
	}

	return &ControllerParser{TplContent: data}
}

func (this *ControllerParser) Parse(pi *astparser.PInterface, dir string, target string) {
	tpl := template.New("controller").Funcs()
}
