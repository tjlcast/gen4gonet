package builder

import (
	"fmt"
	"github.com/tjlcast/gen4gonet/parser_gener"
	"strings"
)

var (
	RESTAPISURFIX = "RestApi"

	REST = "Rest"
)

func BuildRestApi(bean *parser_gener.Struct) string {
	bean.Fields.Append(parser_gener.NewVarTypePair("innService", "*"+bean.Name))

	FilterAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		for _, anno := range f.Annos {
			switch anno.Type {
			case "POST":
				return true
			case "GET":
				return true
			case "PUT":
				return true
			case "DELETE":
				return true
			}
		}
		return false
	})

	restBeanName := bean.Name + RESTAPISURFIX
	UpdateClzName(bean, restBeanName)

	var RegBody []string
	RegTpl := `router.AddRestHandle(net_utils.REST_%s, %s, s.%s)`
	RangeAllFunc(bean, func(b *parser_gener.Struct, f *parser_gener.StructFunc) bool {
		f.Name = f.Name + REST

		f.Args = []*parser_gener.VarTypePair{parser_gener.NewVarTypePair("c", "*gin.Context")}
		f.Rets = nil

		// 查看注解
		for _, anno := range f.Annos {
			switch anno.Type {
			case "GET":
				RegBody = append(RegBody, fmt.Sprintf(RegTpl, anno.Type, anno.Value, f.Name))
			case "POST":
				RegBody = append(RegBody, fmt.Sprintf(RegTpl, anno.Type, anno.Value, f.Name))
			case "DELETE":
				RegBody = append(RegBody, fmt.Sprintf(RegTpl, anno.Type, anno.Value, f.Name))
			case "PUT":
				RegBody = append(RegBody, fmt.Sprintf(RegTpl, anno.Type, anno.Value, f.Name))
			}
		}
		return true
	})

	registerFunc := parser_gener.NewStructFunc()
	registerFunc.Self = bean.Name
	registerFunc.Name = "Register"
	registerFunc.Args = []*parser_gener.VarTypePair{parser_gener.NewVarTypePair("router", "*net_utils.RouterMulti")}
	registerFunc.FuncBody = strings.Join(RegBody, "\n")
	registerFunc.Rets = []*parser_gener.VarType{parser_gener.NewVarType("error")}
	bean.Funcs.Append(registerFunc)

	var header = `
import "github.com/tjlcast/go_common/net_utils"
import "github.com/gin-gonic/gin"
	`
	tpl := bean.Tpl()

	return header + tpl
}


