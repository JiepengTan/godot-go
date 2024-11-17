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
	//go:embed ffi/ffi_wrapper.h.tmpl
	ffi_wrapper_h string

	//go:embed ffi/ffi_wrapper.go.tmpl
	ffi_wrapper_go string

	//go:embed ffi/ffi.go.tmpl
	ffi_go string
)

func GenerateFFI() {
	relDir = GenerateRelDir + "ffi"
	codeGenFFIWrapperHeader()
	codeGenFFIWrapperGo()
	codeGenFFIGo()
}

func codeGenFFIWrapperGo() error {
	return RenderCode("ffi_wrapper.gen.go", ffi_wrapper_go, template.FuncMap{
		"genFFIWrapperGo": genFFIWrapperGo,
	})
}

func codeGenFFIWrapperHeader() error {
	return RenderCode("ffi_wrapper.gen.h", ffi_wrapper_h, template.FuncMap{
		"genFFIWrapperHeader": genFFIWrapperHeader,
	})
}
func codeGenFFIGo() error {
	return RenderCode("ffi.gen.go", ffi_go, template.FuncMap{
		"genFFIGo": genFFIGo,
	})
}
func genFFIGo(ast clang.CHeaderFileAST) string {
	tempStrBuilder = strings.Builder{}
	return tempStrBuilder.String()
}

func genFFIWrapperGo(ast clang.CHeaderFileAST) string {
	tempStrBuilder = strings.Builder{}
	return tempStrBuilder.String()
}

func genFFIWrapperHeader(ast clang.CHeaderFileAST) string {
	tempStrBuilder = strings.Builder{}
	// deal functions
	funcs := ast.CollectFunctions()
	for _, f := range funcs {
		//int cgo_PtrSetter(const pointer fn,GDExtensionTypePtr* p_base)
		Write("%s cgo_%s(", f.ReturnType.CStyleString(), f.Name)
		Write("const %s fn", f.Name)
		for j, a := range f.Arguments {
			Write(",%s", a.CStyleString(j))
		}
		WriteLine(") {")

		//return fn(p_base) };
		if f.ReturnType.CStyleString() != "void" {
			Write("\treturn")
		}
		Write(" fn(")
		for j, a := range f.Arguments {
			Write("%s", a.ResolvedName(j))
			if j != len(f.Arguments)-1 {
				Write(",")
			}
		}
		Write(");")
		WriteLine("\n}")
	}
	WriteLine("\n\n\n// -------------------- Structs ------------------------- ")

	stucts := ast.CollectStructs()
	{
		for _, t := range stucts {
			if len(t.CollectFunctions()) == 0 {
				continue
			}
			for _, f := range t.CollectFunctions() {
				// void cgo_GDExtensionInitialization_initialize(const GDExtensionInitialization * p_struct, void *  userdata, GDExtensionInitializationLevel p_level){
				Write("%s cgo_%s_%s(", f.ReturnType.CStyleString(), t.Name, f.Name)
				Write("const %s * p_struct", t.Name)
				for j, a := range f.Arguments {
					Write(",%s", a.CStyleString(j))
				}
				WriteLine(") {")

				//p_struct->initialize(userdata, p_level);
				if f.ReturnType.CStyleString() != "void" {
					Write("\treturn")
				}
				Write(" p_struct->%s(", f.Name)
				for j, a := range f.Arguments {
					Write("%s", a.ResolvedName(j))
					if j != len(f.Arguments)-1 {
						Write(",")
					}
				}
				Write(");")
				WriteLine("\n}")

			}
		}
	}

	return tempStrBuilder.String()
}
