package main

import (
	"embed"
	"errors"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	_ "embed"

	"github.com/JiepengTan/godotgo/cmd/gdg/pkg/impl"
	"github.com/JiepengTan/godotgo/cmd/gdg/pkg/util"
)

var (
	//go:embed template/project/*
	proejct_fs embed.FS

	//go:embed template/version
	version string

	//go:embed template/project/.godot/gdg_web_server.py
	gdg_web_server_py string

	//go:embed template/project/.godot/extension_list.cfg
	extension_list_cfg string

	//go:embed template/project/go/go.mod.txt
	go_mode_txt string

	//go:embed template/project/go/go.mod.local
	go_mode_local string

	//go:embed template/project/go/main.go
	main_go string

	//go:embed template/project/project.godot
	project_godot string

	//go:embed template/project/gdg.gdextension
	gdg_gdextension string

	//go:embed template/project/main.tscn
	main_tscn string

	//go:embed template/project/export_presets.cfg
	export_presets_cfg string

	//go:embed template/project/.gitignore
	gitignore string
)

var (
	targetDir   string
	cmdPath     string
	projectPath string
	libPath     string
	serverPort  int = 8005
	binPostfix      = ""
	goPath          = ""
)

func ShowHelpInfo() {
	showHelpInfo("gdg")
}
func showHelpInfo(cmdName string) {
	msg := `
Usage:

    #CMDNAME <command> [path]      

The commands are:

    - init            # Create a #CMDNAME project in the current directory
    - run             # Run the current project
    - editor          # Open the current project in editor mode
    - build           # Build the dynamic library
    - export          # Export the PC package (macOS, Windows, Linux) (TODO)
    - runweb          # Launch the web server
    - buildweb        # Build for WebAssembly (WASM)
    - exportweb       # Export the web package
    - clear           # Clear the project 

 eg:

    #CMDNAME init                      # create a project in current path
    #CMDNAME init ./test/demo01        # create a project at path ./test/demo01 
	`
	fmt.Println(strings.ReplaceAll(msg, "#CMDNAME", cmdName))
}
func UpdateMod() {
	impl.UpdateMod()
}
func CheckEnvironment() {
	impl.CheckEnvironment()
}

func PrepareGoEnv() {
	if !util.IsFileExist(targetDir) {
		err := util.CopyDir(proejct_fs, "template/project", targetDir)
		if err != nil {
			os.Exit(1)
			return
		}
		util.SetupFile(false, targetDir+"/go/go.mod", go_mode_local)
	} else {
		os.MkdirAll(targetDir, 0755)
		if err := util.SetupFile(false, targetDir+"/go/go.mod", go_mode_txt); err != nil {
			panic(err)
		}
		if err := util.SetupFile(false, targetDir+"/go/main.go", main_go); err != nil {
			panic(err)
		}
	}
	rawDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Chdir(targetDir)
	err = util.RunGolang(nil, "mod", "tidy")
	os.Chdir(rawDir)
	if err != nil {
		println("gdg project create failed ", targetDir)
		panic(err)
	} else {
		println("gdg project create succ ", targetDir, "\n=====you can type  'gdg run "+targetDir+"'  to run the project======")
	}

}
func SetupEnv() error {
	var GOOS, GOARCH = runtime.GOOS, runtime.GOARCH
	if os.Getenv("GOOS") != "" {
		GOOS = os.Getenv("GOOS")
	}
	if os.Getenv("GOARCH") != "" {
		GOARCH = os.Getenv("GOARCH")
	}
	if GOARCH != "amd64" && GOARCH != "arm64" {
		return errors.New("gdg requires an amd64, or an arm64 system")
	}
	var err error
	binPostfix, cmdPath, err = impl.CheckAndGetAppPath(version)
	if err != nil {
		return fmt.Errorf("gdg requires Godot to be installed as a binary at $GOPATH/bin/: %w", err)
	}
	projectPath, _ = filepath.Abs(targetDir)
	goPath, _ = filepath.Abs(projectPath + "/go")

	wd := goPath

	for wd := wd; true; wd = filepath.Dir(wd) {
		if wd == "/" {
			return fmt.Errorf("gdg requires your project to have a go.mod file")
		}
		_, err := os.Stat(wd + "/go.mod")
		if err == nil {
			break
		} else if os.IsNotExist(err) {
			continue
		} else {
			return err
		}
	}

	var libraryName = fmt.Sprintf("gdg-%v-%v", GOOS, GOARCH)
	switch GOOS {
	case "windows":
		libraryName += ".dll"
	case "darwin":
		libraryName += ".dylib"
	default:
		libraryName += ".so"
	}
	libPath, _ = filepath.Abs(path.Join(projectPath, "lib", libraryName))

	_, err = os.Stat(projectPath + "/.godot")
	hasInited := !os.IsNotExist(err)
	os.MkdirAll(projectPath+"/.godot", 0755)
	util.SetupFile(false, projectPath+"/main.tscn", main_tscn)
	util.SetupFile(false, projectPath+"/project.godot", project_godot)
	util.SetupFile(false, projectPath+"/.gitignore", gitignore)
	util.SetupFile(true, projectPath+"/gdg.gdextension", gdg_gdextension)
	util.SetupFile(false, projectPath+"/.godot/extension_list.cfg", extension_list_cfg)
	if !hasInited {
		BuildDll()
		ImportProj()
	}
	return nil
}

func ExportWebEditor() error {
	gopath := build.Default.GOPATH
	editorZipPath := path.Join(gopath, "bin", "gdg"+version+"_web.zip")
	dstPath := path.Join(projectPath, ".builds/web")
	os.MkdirAll(dstPath, os.ModePerm)
	if util.IsFileExist(editorZipPath) {
		util.Unzip(editorZipPath, dstPath)
	} else {
		return errors.New("editor zip file not found: " + editorZipPath)
	}
	os.Rename(path.Join(dstPath, "godot.editor.html"), path.Join(dstPath, "index.html"))
	return nil
}
func CheckExportWeb() error {
	if !util.IsFileExist(path.Join(projectPath, ".builds/web")) {
		return ExportWeb()
	}
	return nil
}

func ExportWeb() error {
	BuildDll()
	// Delete gdextension
	os.RemoveAll(filepath.Join(projectPath, "lib"))
	os.Remove(filepath.Join(projectPath, ".godot", "extension_list.cfg"))
	os.Remove(filepath.Join(projectPath, "gdg.gdextension"))
	// Copy template files
	util.SetupFile(false, filepath.Join(projectPath, "export_presets.cfg"), export_presets_cfg)

	BuildWasm()
	err := ExportBuild("Web")
	return err
}
func Clear() {
	projectPath := targetDir
	os.RemoveAll(filepath.Join(projectPath, "lib"))
	os.RemoveAll(filepath.Join(projectPath, ".godot"))
	os.RemoveAll(filepath.Join(projectPath, ".build"))
}
func StopWebServer() {
	if runtime.GOOS == "windows" {
		content := "taskkill /F /IM python.exe\r\ntaskkill /F /IM pythonw.exe\r\n"
		tempFileName := "temp_kill.bat"
		os.WriteFile(tempFileName, []byte(content), 0644)
		cmd := exec.Command("cmd.exe", "/C", tempFileName)
		cmd.Run()
		os.Remove(tempFileName)
	} else {
		cmd := exec.Command("pkill", "-f", "gdg_web_server.py")
		cmd.Run()
	}
}
func RunWebServer() error {
	if !util.IsFileExist(filepath.Join(projectPath, ".builds", "web")) {
		ExportWeb()
	}
	port := serverPort
	StopWebServer()
	scriptPath := filepath.Join(projectPath, ".godot", "gdg_web_server.py")
	executeDir := filepath.Join(projectPath, "../", ".builds/web")
	util.SetupFile(false, scriptPath, gdg_web_server_py)
	println("web server running at http://localhost:" + fmt.Sprint(port))
	cmd := exec.Command("python", scriptPath, "-r", executeDir, "-p", fmt.Sprint(port))
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting server: %v", err)
	}
	return nil
}

func Export() error {
	platform := "Win"
	args := "--headless --quit --export-debug " + platform
	println("run: ", cmdPath, projectPath, args)
	gdg := exec.Command(cmdPath, args)
	gdg.Dir = projectPath
	gdg.Stderr = os.Stderr
	gdg.Stdout = os.Stdout
	gdg.Stdin = os.Stdin
	return gdg.Run()
}

func BuildDll() {
	rawdir, _ := os.Getwd()
	os.Chdir(goPath)
	envVars := []string{"CGO_ENABLED=1"}
	util.RunGolang(envVars, "build", "-o", libPath, "-buildmode=c-shared")
	os.Chdir(rawdir)
}

func BuildWasm() {
	rawdir, _ := os.Getwd()
	dir := path.Join(projectPath, ".builds/web/")
	os.MkdirAll(dir, 0755)
	filePath := path.Join(dir, "gdg.wasm")
	os.Chdir(goPath)
	envVars := []string{"GOOS=js", "GOARCH=wasm"}
	util.RunGolang(envVars, "build", "-o", filePath)
	os.Chdir(rawdir)
}

func ExportBuild(platform string) error {
	println("start export: platform =", platform, " projectPath =", projectPath)
	os.MkdirAll(filepath.Join(projectPath, ".builds", strings.ToLower(platform)), os.ModePerm)
	cmd := exec.Command(cmdPath, "--headless", "--quit", "--path", projectPath, "--export-debug", platform)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error exporting to web:", err)
	}
	return err
}

func ImportProj() error {
	cmd := exec.Command(cmdPath, "--import", "--headless")
	cmd.Dir = projectPath
	err := cmd.Start()
	err = cmd.Wait()
	if err != nil {
		fmt.Println("ImportProj finished")
	} else {
		fmt.Println("ImportProj successfully")
	}
	return nil
}

func Run(args string) error {
	util.SetupFile(false, filepath.Join(projectPath, ".godot", "extension_list.cfg"), extension_list_cfg)
	util.SetupFile(false, filepath.Join(projectPath, "gdg.gdextension"), gdg_gdextension)

	println("run: ", cmdPath, projectPath, args)
	cmd := exec.Command(cmdPath, args)
	cmd.Dir = projectPath
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
