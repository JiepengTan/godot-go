// Package gdextensionwrapper generates C code to wrap all of the gdextension
// methods to call functions on the gdextension_api_structs to work
// around the Cgo C function pointer limitation.
package generator

import (
	_ "embed"
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
	WriteLine("")
	return tempStrBuilder.String()
}
