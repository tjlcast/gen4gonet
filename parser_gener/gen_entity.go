package parser_gener

import "strings"

var (
	EntityTpl = `
type strutc {{name}} {
	{{fields}}
}
`
)

type Entity struct {
	Name string

	Fields  VarTypePairArr
	NameMap map[string]*VarTypePair
}

func NewEntity(name string, fields VarTypePairArr, nameMap map[string]*VarTypePair) *Entity {
	return &Entity{Name: name, Fields: fields, NameMap: nameMap}
}

func (e *Entity) Tpl() string {
	tpl := strings.ReplaceAll(EntityTpl, "{{name}}", e.Name)
	tpl = strings.ReplaceAll(tpl, "{{fields}}", e.Fields.Tpl())
	return tpl
}

func (e *Entity) PtrName() string {
	return "*" + e.Name
}
