package builder

import (
	"github.com/tjlcast/gen4gonet/parser_gener"
)

func UpdateClzName(bean *parser_gener.Struct, newName string) {
	bean.Name = newName

	for _, f := range bean.Funcs {
		f.Self = newName
	}
}

func RangeAllFunc(bean *parser_gener.Struct, doFunc func(*parser_gener.Struct, *parser_gener.StructFunc) bool) {
	for _, f := range bean.Funcs {
		goon := doFunc(bean, f)
		if !goon {
			return
		}
	}
}

func FilterAllFunc(bean *parser_gener.Struct, doFunc func(*parser_gener.Struct, *parser_gener.StructFunc) bool) {
	var newFuncArr parser_gener.StructFuncArr
	for _, f := range bean.Funcs {
		goon := doFunc(bean, f)
		if goon {
			newFuncArr.Append(f)
		}
	}
	bean.Funcs = newFuncArr
}

func Base(bean *parser_gener.Struct) string {
	return bean.Tpl()
}
