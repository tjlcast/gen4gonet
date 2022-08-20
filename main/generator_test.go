package main

import (
	"fmt"
	"github.com/tjlcast/gen4gonet/parser_gener"
	"github.com/tjlcast/gen4gonet/builder"
	"os"
	"path"
	"testing"
)

func getPath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, "server.txt")
}

func TestGen1(t *testing.T) {
	path := getPath()
	bean := parser_gener.ParseAClzFile(path)

	var newTpl string
	newTpl = builder.Base(bean)
	fmt.Println(newTpl)
}

func TestGen2(t *testing.T) {
	path := getPath()
	bean := parser_gener.ParseAClzFile(path)

	var newTpl string
	newTpl = builder.BuildRestApi(bean)
	fmt.Println(newTpl)
}

func TestGen3(t *testing.T) {
	path := getPath()
	bean := parser_gener.ParseAClzFile(path)

	var newTpl string
	newTpl = builder.BuildRpcCliApi(bean)
	fmt.Println(newTpl)
}

func TestGen4(t *testing.T) {
	path := getPath()
	bean := parser_gener.ParseAClzFile(path)

	var newTpl string
	newTpl = builder.BuildRpcSrvApi(bean)
	fmt.Println(newTpl)
}
