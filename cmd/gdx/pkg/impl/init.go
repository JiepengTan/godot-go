package impl

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	_ "embed"

	"github.com/JiepengTan/godotgo/cmd/gdx/pkg/util"
)

func findFirstMatchingFile(dir, pattern, exclude string) string {
	var foundFile string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if match, _ := filepath.Match(pattern, info.Name()); match {
			if !strings.Contains(info.Name(), exclude) {
				foundFile = path
				return filepath.SkipDir
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return foundFile
}

// Install SCons and Ninja using pip
func installPythonPackages() {
	util.RunCommand(nil, "pip", "install", "scons==4.7.0")
}

// Get the Go environment variable
func getGoEnv() string {
	cmd := exec.Command("go", "env", "GOPATH")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to get GOPATH: %v", err)
	}
	return string(output)
}

func downloadPack(dstDir, tagName, postfix string) error {
	urlHeader := "https://github.com/JiepengTan/godot/releases/download/"
	fileName := tagName + postfix
	url := urlHeader + tagName + "/" + fileName
	// download pc
	err := util.DownloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	// download web
	fileName = tagName + "_web.zip"
	url = urlHeader + tagName + "/" + fileName
	err = util.DownloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	// download webpack
	fileName = tagName + "_webpack.zip"
	url = urlHeader + tagName + "/" + fileName
	err = util.DownloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	return err
}

// Helper function to check if a specific Python command exists
func checkAppInstalled(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func CheckEnvironment() {
	if !checkAppInstalled("python3") && !checkAppInstalled("python") {
		fmt.Println("Python is not installed. Please install Python first, python version should >= 3.8")
		os.Exit(1)
	}
}

func CheckAndGetAppPath(version string) (string, string, error) {
	binPostfix := ""
	if runtime.GOOS == "windows" {
		binPostfix = "_win.exe"
	} else if runtime.GOOS == "darwin" {
		binPostfix = "_darwin"
	} else if runtime.GOOS == "linux" {
		binPostfix = "_linux"
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	tagName := "gdx" + version
	dstFileName := tagName + binPostfix
	gdx, err := exec.LookPath(dstFileName)
	if err == nil {
		if _, err := exec.Command(gdx, "--version").CombinedOutput(); err == nil {
			return binPostfix, gdx, nil
		}
	}

	dstDir := gopath + "/bin"
	cmdPath := path.Join(dstDir, dstFileName)
	info, err := os.Stat(cmdPath)
	if os.IsNotExist(err) {
		println("Downloading gdx pack...")
		err := downloadPack(dstDir, tagName, binPostfix)
		if err != nil {
			print("downloadPack error:" + err.Error())
			return binPostfix, cmdPath, err
		}
		if err := os.Chmod(cmdPath, 0755); err != nil {
			return binPostfix, cmdPath, err
		}
	} else if err != nil {
		return binPostfix, "", err
	} else {
		if info.Mode()&0111 == 0 {
			if err := os.Chmod(cmdPath, 0755); err != nil {
				return binPostfix, cmdPath, err
			}
		}
	}
	return binPostfix, cmdPath, nil
}
