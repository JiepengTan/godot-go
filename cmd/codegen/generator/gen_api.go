// Package gdextensionwrapper generates C code to wrap all of the gdextension
// methods to call functions on the gdextension_api_structs to work
// around the Cgo C function pointer limitation.
package generator

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/JiepengTan/godotgo/cmd/codegen/gdextensionparser/clang"
)

var (
	//go:embed api/api.go.tmpl
	api_go string
)

func GenerateAPI() {
	relDir = GenerateRelDir + "api"
	codeGenAPIGo()
}
func codeGenAPIGo() error {
	return RenderCode("api.gen.go", api_go, template.FuncMap{
		"genAPIGo": genAPIGo,
	})
}
func genAPIGo(ast clang.CHeaderFileAST) string {
	tempStrBuilder = strings.Builder{}
	funcs := ast.CollectGDExtensionInterfaceFunctions()
	for _, f := range funcs {
		params := ""
		for j, a := range f.Arguments {
			name := a.Name
			if a.Name == "" {
				name = "p_func"
			}
			params += name + " " + ToGoTypeString(a.Type)
			if a.Name == "" {
				params += fmt.Sprintf("/*%s*/", a.Type.CStyleString())
			}
			if j != len(f.Arguments)-1 {
				params += ","
			}
		}

		retStr := "" + toGoTypeString(f.ReturnType.GoString())
		if retStr == "void" {
			retStr = ""
		}
		WriteLine("%s func(%s) %s", strings.Replace(f.Name, "GDExtensionInterface", "", 1), params, retStr)
	}
	return tempStrBuilder.String()
}

var (
	typeMap = map[string]string{
		"size_t":                     "int64",
		"int32_t":                    "int32",
		"uint32_t":                   "uint32",
		"int64_t":                    "int64",
		"uint64_t":                   "uint64",
		"GDObjectInstanceID":         "int64",
		"GDExtensionInt":             "int64",
		"GDExtensionVariantType":     "uint32/*VariantType*/",
		"GDExtensionVariantOperator": "uint32/*VariantOperator*/",
		"GDExtensionBool":            "bool",
		"char32_t":                   "rune",
		"const char *":               "string",
	}
)

func ToGoTypeString(t clang.Type) string {
	return toGoTypeString(t.GoString())
}
func toGoTypeString(typeName string) string {
	if value, exists := typeMap[typeName]; exists {
		return value
	}
	return typeName
}
