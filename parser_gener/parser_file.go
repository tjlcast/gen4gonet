package parser_gener

import (
	"github.com/tjlcast/go_common/file_utils"
	"strings"
)

type Template interface {
	Tpl() string
}

func ParseAClzFile(path string)*Struct {
	lines, e := file_utils.ReadFileLines(path)
	if e != nil {
		panic(e)
	}

	var parseStruct *Struct
	var parseFuncs []*StructFunc

	var curAnnoGroup = []*Anno{}
	for ; len(lines) != 0; lines = lines[1:] {
		line := lines[0]
		line = strings.TrimSpace(line)

		anno, b := GetAnno(line)
		if b {
			curAnnoGroup = append(curAnnoGroup, anno)
			continue
		} else {
			if strings.HasPrefix(line, TYPE) {
				parseStruct, e = ParseStruct(line)
				if e != nil {
					panic(e)
				}
				parseStruct.Annos = curAnnoGroup
			} else if strings.HasPrefix(line, FUNC) {
				structFunc, e := ParseFunc(&lines)
				if e != nil {
					panic(e)
				}
				structFunc.Annos = curAnnoGroup
				parseFuncs = append(parseFuncs, structFunc)
			}
			// 清空已经扫描到的注解
			curAnnoGroup = []*Anno{}
		}
	}

	for _, f := range parseFuncs {
		if parseStruct.Name == f.GetSelfEntity() {
			parseStruct.Funcs = append(parseStruct.Funcs, f)
		}
	}

	// fmt.Println(parseStruct.Tpl())
	return parseStruct
}
