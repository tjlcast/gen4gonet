package parser_gener

import (
	"fmt"
	"strings"
)

type StructFuncArr []*StructFunc

func (s *StructFuncArr) Append(sFunc *StructFunc) {
	*s = append(*s, sFunc)
}

type StructFunc struct {
	Self string

	Name string
	Args VarTypePairArr
	Rets VarTypeArr

	Annos []*Anno

	FuncBody string
}

func NewStructFunc() *StructFunc {
	return &StructFunc{
		FuncBody: DefaultFuncBody,
	}
}

func (s *StructFunc) GetSelfEntity() string {
	if strings.Contains(s.Self, "*") {
		return s.Self[strings.Index(s.Self, "*")+1:]
	}
	return s.Self
}

func ParseFunc(lines *[]string) (*StructFunc, error) {
	line := (*lines)[0]
	line = strings.TrimSpace(line[4:])

	if strings.HasPrefix(line, "(") {
		return ParseStructFunc(lines)
	} else {
		return ParsePureFunc(lines)
	}
}

func ParseStructFunc(lines *[]string) (*StructFunc, error) {
	var funcSign = ""

	for ; len(*lines) != 0 && !strings.Contains((*lines)[0], "{"); *lines = (*lines)[1:] {
		funcSign += (*lines)[0]
	}
	if len(*lines) != 0 {
		funcSign += (*lines)[0]
		*lines = (*lines)[1:]
	}

	funcSign = strings.TrimSpace(funcSign)
	funcSign = funcSign[0:strings.Index(funcSign, "{")]
	funcSign = strings.TrimSpace(funcSign)
	funcSign = strings.ReplaceAll(funcSign, "\t", "")

	// "func (xxx) xxx (xxx)"
	submatch := StructFuncRegex.FindStringSubmatch(funcSign)

	structFunc := NewStructFunc()
	structFunc.Self = submatch[1]
	structFunc.Name = submatch[2]
	structFunc.Args = ParseVTPairArr(submatch[3])
	structFunc.Rets = ParseVTArr(submatch[4])

	// fmt.Println(structFunc.Tpl())

	return structFunc, nil
}

func ParsePureFunc(lines *[]string) (*StructFunc, error) {
	var funcSign string

	for ; len(*lines) != 0 && !strings.Contains((*lines)[0], "{"); *lines = (*lines)[1:] {
		funcSign += (*lines)[0]
	}
	if len(*lines) != 0 {
		funcSign += (*lines)[0]
		*lines = (*lines)[1:]
	}

	funcSign = strings.TrimSpace(funcSign)
	funcSign = funcSign[0:strings.Index(funcSign, "{")]
	funcSign = strings.TrimSpace(funcSign)
	funcSign = strings.ReplaceAll(funcSign, "\t", "")

	// "func (xxx) xxx (xxx)"
	submatch := PureFuncRegex.FindStringSubmatch(funcSign)

	pureFunc := NewStructFunc()
	pureFunc.Name = submatch[1]
	pureFunc.Args = ParseVTPairArr(submatch[2])
	pureFunc.Rets = ParseVTArr(submatch[3])

	return pureFunc, nil
}

func (s *StructFunc) Tpl() string {
	pureTpl := `
func %s(%s) %s {
	%s
}
`

	structTpl := `
func (s *%s)%s(%s) %s {
	%s
}
`
	if s.Self == "" {
		return fmt.Sprintf(pureTpl, s.Name, s.Args.Tpl(), s.Rets.Tpl(), s.FuncBody)
	} else {
		return fmt.Sprintf(structTpl, s.GetSelfEntity(), s.Name, s.Args.Tpl(), s.Rets.Tpl(), s.FuncBody)
	}
}
