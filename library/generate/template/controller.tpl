package controller

{{MakeImports .Imports}}

{{$interface:=.Name}}

// {{.Name}} 接口实现类

type {{MakeControllerName $interface}} struct {

}

func New{{MakeControllerName $interface}} () *{{MakeControllerName $interface}} {
    return &{{MakeControllerName $interface}}{}
}

{{range .MethodsNew}}
func (this *{{MakeControllerName $interface}}) {{.FunName}}({{MakeParams .FunParams}}) {{MakeResults .FunResult}} {
    // your codes
}

{{end}}
