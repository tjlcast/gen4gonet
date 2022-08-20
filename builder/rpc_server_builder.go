package builder

import (
	"github.com/tjlcast/gen4gonet/parser_gener"
)

var (
	SRVRPCAPISURFIX = "SrvRpcApi"
	SRVRPC = "SrvRpc"
)

func BuildRpcSrvApi(bean *parser_gener.Struct) string {
	bean.Fields.Append(parser_gener.NewVarTypePair("innService", "*"+bean.Name))

	FilterAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		for _, anno := range f.Annos {
			switch anno.Type {
			case "RPC":
				return true
			}
		}
		return false
	})

	restBeanName := bean.Name + SRVRPCAPISURFIX
	UpdateClzName(bean, restBeanName)



	RangeAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		f.Name = f.Name + SRVRPC

		// 查看注解
		for _, anno := range f.Annos {
			switch anno.Type {
			case "RPC":
				doc := anno.Value
				arr := parser_gener.ParseVTPairArr(doc)
				f.Args = arr
				f.Rets = []*parser_gener.VarType{parser_gener.NewVarType("error"),}
			}
		}
		return true
	})

	registerFunc := parser_gener.NewStructFunc()
	registerFunc.Self = bean.Name
	registerFunc.Name = "Register"
	registerFunc.Args = []*parser_gener.VarTypePair{parser_gener.NewVarTypePair("router", "*net_utils.RouterMulti")}
	registerFunc.FuncBody = "router.AddRpcHandle(s)"
	registerFunc.Rets = []*parser_gener.VarType{parser_gener.NewVarType("error")}
	bean.Funcs.Append(registerFunc)

	var header = `
import "github.com/tjlcast/go_common/net_utils"
import "github.com/gin-gonic/gin"
	`
	tpl := bean.Tpl()

	return header + tpl
}

