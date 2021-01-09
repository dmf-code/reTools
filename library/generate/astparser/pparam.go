package astparser

import (
	"go/ast"
	"sort"
)

type FieldMap map[int]string

//排序取出key，默认从大到小
func (this FieldMap) Keys() []int {
	keys :=make([]int, 0)
	for key := range this {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

//根据排序后的key取出value集合
func (this FieldMap) Values() []interface{} {
	ret := make([]interface{}, 0)
	keys := this.Keys()
	for _, key := range keys {
		ret = append(ret, this[key])
	}

	return ret
}

type PParam struct {
	Name		string
	FieldPos	FieldMap // 记录每个exp的起始pos
}

func NewPParam(name string, t ast.Expr) *PParam {
	ret := &PParam{name, make(map[int]string, 0)}
	ast.Walk(ret, t)

	return ret
}

func (this *PParam) Visit(node ast.Node) ast.Visitor {
	switch expr := node.(type) {
	case *ast.Ident:
		this.FieldPos[int(expr.Pos())] =expr.Name // 记录起始位置
	case *ast.ArrayType: // 切片类型
		this.FieldPos[int(expr.Pos())] = "[]"
	case *ast.InterfaceType: //interface类型
		this.FieldPos[int(expr.Pos())] = "interface{}"
	case *ast.MapType: // map类型
		this.FieldPos[int(expr.Map)] = "map"
		this.FieldPos[int(expr.Key.Pos()-1)] = "["
		this.FieldPos[int(expr.Key.End())] = "]"
	case *ast.SelectorExpr:
		this.FieldPos[int(expr.End())] = "."
		break
	case *ast.StarExpr: // 指针
		this.FieldPos[int(expr.Pos())] = "*"
	case *ast.ChanType: // chan类型
		this.FieldPos[int(expr.Pos())] = "chan " // 空格必须要
	}
	return this
}
