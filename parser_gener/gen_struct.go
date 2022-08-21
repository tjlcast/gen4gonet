package parser_gener

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type Struct struct {
	Package string

	Imports string

	Name  string

	Fields VarTypePairArr

	Funcs StructFuncArr

	Annos []*Anno

	//
	StructBody string
	ConstructBody string
}

func NewStruct(sName string) *Struct {
	return &Struct{Package: "main", Name: sName, Funcs: []*StructFunc{}, Fields: []*VarTypePair{}}
}

func ParseStruct(doc string) (*Struct, error) {
	if !strings.Contains(doc, "struct") {
		return nil, errors.New("Parse error: There is no struct: " + doc)
	}

	submatch := StructRegex.FindStringSubmatch(doc)
	structName := submatch[1]

	return NewStruct(structName), nil
}

func (s *Struct) ConstructTpl() string {
	var args []string
	var fuzhi []string
	for _, field := range s.Fields {
		args = append(args, field.vName+"1 "+field.vType)
		fuzhi = append(fuzhi, field.vName+": "+field.vName+"1,")
	}

	tpl := `
func New{{name}}(%s) *{{name}} {
	return &{{name}}{%s}
}`
	tpl = strings.ReplaceAll(tpl, "{{name}}", s.Name)

	tpl = fmt.Sprintf(tpl, strings.Join(args, ","), strings.Join(fuzhi, ""))
	return tpl
}

func (s *Struct) AllFuncTpl() string {
	var funcsTpl []string
	for _, sfunc := range s.Funcs {
		funcsTpl = append(funcsTpl, sfunc.Tpl())
	}
	return strings.Join(funcsTpl, "\n")
}

func (s *Struct) Tpl() string {
	conTpl := s.ConstructTpl()
	allFuncTpl := s.AllFuncTpl()

	tpl := `
package %s

%s

/// Auto Generated
type %s struct {
	%s
}

// constructor
%s 

// funcs
%s 
`
	tpl = fmt.Sprintf(tpl, s.Package, s.Imports, s.Name, s.Fields.FieldsTpl(), conTpl, allFuncTpl)
	return tpl
}
