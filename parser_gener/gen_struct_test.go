package parser_gener

import (
	"fmt"
	"testing"
)

func TestParseStruct(t *testing.T) {
	doc := "type dafdasf struct "

	parseStruct, _ := ParseStruct(doc)
	fmt.Println(parseStruct.Tpl())
}
