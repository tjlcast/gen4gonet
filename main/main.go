package main

import (
	"flag"
	"fmt"
	"github.com/tjlcast/gen4gonet/builder"
	"github.com/tjlcast/gen4gonet/parser_gener"
	"github.com/tjlcast/go_common/file_utils"
	"github.com/tjlcast/go_common/log_utils"
	"os"
	"path"
)

var (
	curPath string
	err     error

	srvFilePath = flag.String("f", "", "The to parse file.")
	workType    = flag.String("g", "", "The type to generate.")
)

func main() {
	curPath, err = os.Getwd()
	flag.Parse()

	filePath := path.Join(curPath, *srvFilePath)
	exist := file_utils.IsExist(filePath)
	if !exist {
		log_utils.Logger.Error("Fail to open: ", filePath)
	}

	bean := parser_gener.ParseAClzFile(filePath)
	var tpl string
	switch *workType {
	case "base":
		tpl = builder.Base(bean)
	case "rpcs":
		tpl = builder.BuildRpcSrvApi(bean)
	case "rpcc":
		tpl = builder.BuildRpcCliApi(bean)
	case "rests":
		tpl = builder.BuildRestApi(bean)
	default:
		log_utils.Logger.Error("Not support -g: " + *workType)
		os.Exit(-1)
	}

	fmt.Println(tpl)
}
