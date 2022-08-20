package parser_gener

import "regexp"

var (
	FUNC = "func"
	TYPE = "type"

	// anno
	// @ http://c.biancheng.net/view/5124.html
	AnnoKVRegex = regexp.MustCompile("(.+)\\((.*)\\)")
	AnnoKRegex  = regexp.MustCompile("(.+)")

	// struct
	StructRegex = regexp.MustCompile("type (\\w*) struct")

	// var type pair
	VarTypeRegex = regexp.MustCompile("(.+) +(.+)")

	// function
	StructFuncRegex = regexp.MustCompile("func *?\\((.+?)\\) *?(\\w*?) *?\\((.*?)\\)(.*)")
	PureFuncRegex   = regexp.MustCompile("func *(\\w*?) *?\\((.*?)\\)(.*)")

	// Default func body
	DefaultFuncBody = `
	// todo
	return nil`


)
