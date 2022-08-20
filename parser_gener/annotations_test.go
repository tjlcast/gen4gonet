package parser_gener

import (
	"fmt"
	"testing"
)

func TestGetAnno(t *testing.T) {
	doc1 := "// @SERVER"
	anno, b := GetAnno(doc1)
	if b {
		fmt.Println(anno)
	}

	doc2 := "// @SERVER()"
	anno, b = GetAnno(doc2)
	if b {
		fmt.Println(anno)
	}

	doc3 := "// @SERVER (vaf)"
	anno, b = GetAnno(doc3)
	if b {
		fmt.Println(anno)
	}
}
