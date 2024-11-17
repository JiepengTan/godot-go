package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	_ "embed"

	"github.com/JiepengTan/godotgo/cmd/codegen/extensionapiparser"
	"github.com/JiepengTan/godotgo/cmd/codegen/gdextensionparser"
	"github.com/JiepengTan/godotgo/cmd/codegen/gdextensionparser/clang"
	"github.com/JiepengTan/godotgo/cmd/codegen/generator"
)

var (
	//go:embed data/extension_api.json
	extension_api_json string
	//go:embed data/gdextension_interface.h
	gdextension_interface_h string
)
var (
	verbose          bool
	cleanAll         bool
	cleanGdextension bool
	cleanTypes       bool
	cleanClasses     bool
	genClangAPI      bool
	genExtensionAPI  bool
	packagePath      string
	godotPath        string
	parsedASTPath    string
	buildConfig      string
)

func init() {
	absPath, _ := filepath.Abs(".")
	var (
		defaultBuildConfig string
	)
	if strings.Contains(runtime.GOARCH, "32") {
		defaultBuildConfig = "float_32"
	} else {
		defaultBuildConfig = "float_64"
	}
	verbose = true
	genClangAPI = true
	genExtensionAPI = false
	packagePath = absPath
	godotPath = "godot"
	parsedASTPath = "_debug_parsed_ast.json"
	buildConfig = defaultBuildConfig
}

func generateCode() error {
	var (
		ast  clang.CHeaderFileAST
		eapi extensionapiparser.ExtensionApi
		err  error
	)
	if verbose {
		println(fmt.Sprintf(`build configuration "%s" selected`, buildConfig))
	}

	// generate go wrap code
	if genClangAPI {
		ast, err = gdextensionparser.GenerateGDExtensionInterfaceAST(gdextension_interface_h, packagePath, parsedASTPath)
		if err != nil {
			panic(err)
		}
	}
	if genExtensionAPI {
		eapi, err = extensionapiparser.GenerateExtensionAPI(extension_api_json, buildConfig)
		if err != nil {
			panic(err)
		}
		if eapi.Classes != nil {
			println("eapi is not nil")
		}
	}
	if genClangAPI {
		if verbose {
			println("Generating gdextension C wrapper functions...")
		}
		generator.Setup(packagePath, ast)
		generator.GenerateFFI()
		generator.GenerateAPI()
	}

	if verbose {
		println("cli tool done")
	}
	return nil
}
func execGoFmt(filePath string) {
	cmd := exec.Command("gofmt", "-w", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Panic(fmt.Errorf("error running gofmt: \n%s\n%w", output, err))
	}
}

func execGoImports(filePath string) {
	cmd := exec.Command("goimports", "-w", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Print(fmt.Errorf("error running goimports: \n%s\n%w", output, err))
	}
}

func main() {
	generateCode()
}
