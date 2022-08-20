package builder

import (
	"github.com/tjlcast/gen4gonet/parser_gener"
)

var (
	CLIRPCAPISURFIX = "CliRpcApi"
	CLIRPC = "CliRpc"
)

func BuildRpcCliApi(bean *parser_gener.Struct) string {
	oldBeanName := bean.Name
	bean.Fields.Append(parser_gener.NewVarTypePair("endpoint", "string"))

	FilterAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		for _, anno := range f.Annos {
			switch anno.Type {
			case "RPC":
				return true
			}
		}
		return false
	})

	restBeanName := bean.Name + CLIRPCAPISURFIX
	UpdateClzName(bean, restBeanName)

	RangeAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		srvFName := f.Name + SRVRPC
		f.Name = f.Name + CLIRPC

		// 查看注解
		for _, anno := range f.Annos {
			switch anno.Type {
			case "RPC":
				doc := anno.Value
				arr := parser_gener.ParseVTPairArr(doc)
				f.Args = arr

				req := arr[0]
				res := arr[1]

				f.FuncBody = `return net_utils.SendTcp(s.endpoint, "`+oldBeanName+`.`+srvFName+`", `+req.Name()+`, &`+res.Name()+`)`
				f.Rets = []*parser_gener.VarType{parser_gener.NewVarType("error"),}
			}
		}
		return true
	})

	var header = `
import "github.com/tjlcast/go_common/net_utils"
import "github.com/gin-gonic/gin"
	`
	tpl := bean.Tpl()

	return header + tpl
}

