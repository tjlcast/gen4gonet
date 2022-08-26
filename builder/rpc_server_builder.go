package builder

import (
	"github.com/tjlcast/gen4gonet/parser_gener"
	"github.com/tjlcast/go_common/string_utils"
)

var (
	SRVRPCAPISURFIX = "SrvRpcApi"
	SRVRPC          = "SrvRpc"
)

var (
	RpcRequests = make(map[string]*parser_gener.Entity)
	RpcReplys   = make(map[string]*parser_gener.Entity)
)

func BuildRpcSrvApi(bean *parser_gener.Struct) string {
	bean.Fields.Append(parser_gener.NewVarTypePair("innService", "*"+bean.Name))

	// 筛选有RPC注解的方法
	FilterAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		if parser_gener.IsAwsomeAnno("RPC", f.Annos) {
			return true
		}
		return false
	})

	restBeanName := bean.Name + SRVRPCAPISURFIX
	UpdateClzName(bean, restBeanName)

	RangeAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		f.Name = f.Name + SRVRPC

		// 查看注解
		// 默认，认为func的入参全为request，func的返回为reply;
		// anno 指导rpc方法的生成;
		if parser_gener.IsAwsomeAnno("RPC", f.Annos) {
			f.Rets = []*parser_gener.VarType{parser_gener.NewVarType("error"),}
		}

		requestEntity := CreateRequestEntity(f, f.Args)
		RpcRequests[f.Name] = requestEntity
		replyEntity := CreateReplyStruct(f, f.Rets)
		RpcReplys[f.Name] = replyEntity

		// 设置rpc-func的req与rep
		f.Comments = "\n//REQ"+requestEntity.Tpl()+"\n//REP"+replyEntity.Tpl()+"\n"
		params := parser_gener.VarTypePairArr{}
		params.Append(parser_gener.NewVarTypePair("req", requestEntity.Name))
		params.Append(parser_gener.NewVarTypePair("rep", replyEntity.PtrName()))
		f.Args = params
		f.FuncBody = "// todo 赋值、调用、返回"
		return true
	})

	registerFunc := parser_gener.NewStructFunc()
	registerFunc.Self = bean.Name
	registerFunc.Name = "Register"
	registerFunc.Args = []*parser_gener.VarTypePair{parser_gener.NewVarTypePair("router", "*net_utils.RouterMulti")}
	registerFunc.FuncBody = "router.AddRpcHandle(s)"
	registerFunc.Rets = []*parser_gener.VarType{parser_gener.NewVarType("error")}
	bean.Funcs.Append(registerFunc)

	bean.Imports = `
import "github.com/tjlcast/go_common/net_utils"
import "github.com/gin-gonic/gin"
	`
	tpl := bean.Tpl()

	return tpl
}

func CreateRequestEntity(f *parser_gener.StructFunc, fields parser_gener.VarTypePairArr) *parser_gener.Entity {
	nameMap := make(map[string]*parser_gener.VarTypePair)
	for _, argPair := range fields {
		nameMap[argPair.VName] = argPair
		argPair.VName = string_utils.FirstLower(argPair.VName)
	}
	return parser_gener.NewEntity(f.Name+"Req", fields, nameMap)
}

func CreateReplyStruct(f *parser_gener.StructFunc, rets parser_gener.VarTypeArr) *parser_gener.Entity {
	nameMap := make(map[string]*parser_gener.VarTypePair)
	var fields parser_gener.VarTypePairArr
	for _, ret := range rets {
		pair := parser_gener.NewVarTypePair(string_utils.FirstLower(ret.VName), ret.VName)
		nameMap[ret.VName] = pair
		fields.Append(pair)
	}
	return parser_gener.NewEntity(f.Name+"Rep", fields, nameMap)
}
