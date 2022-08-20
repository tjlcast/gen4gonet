package parser_gener

import (
	"fmt"
	"strings"
)

type VarTypePair struct {
	vName string
	vType string
}

func (s *VarTypePair) Name() string {
	return s.vName
}

func (s *VarTypePair) Type() string {
	return s.vType
}

func NewVarTypePair(vName string, vType string) *VarTypePair {
	return &VarTypePair{vName: vName, vType: vType}
}

func (s *VarTypePair) Tpl() string {
	tpl := "%s %s"
	tpl = fmt.Sprintf(tpl, s.vName, s.vType)
	return tpl
}

type VarTypePairArr []*VarTypePair

func (s *VarTypePairArr) Append(pair *VarTypePair) {
	*s = append(*s, pair)
}

func ParseVTPairArr(doc string) VarTypePairArr {
	var arr VarTypePairArr
	if doc == "" {
		return arr
	}

	doc = strings.TrimSpace(doc)
	doc = strings.TrimLeft(doc, "(")
	doc = strings.TrimRight(doc, ")")

	doc = strings.ReplaceAll(doc, "\n", "")
	doc = strings.ReplaceAll(doc, "\r", "")

	split := strings.Split(doc, ",")
	for _, term := range split {
		term := strings.TrimSpace(term)
		submatch := VarTypeRegex.FindStringSubmatch(term)
		arr = append(arr, NewVarTypePair(submatch[1], submatch[2]))
	}
	return arr
}

func (s VarTypePairArr) Tpl() string {
	var arr []string
	for _, vt := range s {
		arr = append(arr, vt.Tpl())
	}
	return strings.Join(arr, ", ")
}

func (s VarTypePairArr) FieldsTpl() string {
	if s == nil {
		return ""
	}

	var arr []string
	for _, vt := range s {
		arr = append(arr, fmt.Sprintf("%s %s", vt.vName, vt.vType))
	}
	return strings.Join(arr, "\n")
}

type VarType struct {
	vName string
}

func NewVarType(vName string) *VarType {
	return &VarType{vName: vName}
}

func (s *VarType) Tpl() string {
	tpl := "%s"
	tpl = fmt.Sprintf(tpl, s.vName)
	return tpl
}

type VarTypeArr []*VarType

func ParseVTArr(doc string) VarTypeArr {
	var arr VarTypeArr
	if doc == "" {
		return arr
	}

	doc = strings.TrimSpace(doc)
	doc = strings.TrimLeft(doc, "(")
	doc = strings.TrimRight(doc, ")")

	doc = strings.ReplaceAll(doc, "\n", "")
	doc = strings.ReplaceAll(doc, "\r", "")

	split := strings.Split(doc, ",")
	for _, term := range split {
		term := strings.TrimSpace(term)
		arr = append(arr, NewVarType(term))
	}
	return arr
}

func (s VarTypeArr) Tpl() string {
	if s == nil || len(s) == 0 {
		return ""
	}

	var arr []string
	for _, vt := range s {
		arr = append(arr, vt.Tpl())
	}
	return strings.Join(arr, ", ")
}
