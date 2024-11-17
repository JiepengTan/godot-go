package main

import (
	"fmt"
	"os"
)

func main() {
	CheckEnvironment()
	targetDir = "."
	if len(os.Args) > 2 {
		targetDir = os.Args[2]
	}
	if len(os.Args) <= 1 {
		ShowHelpInfo()
		return
	}

	switch os.Args[1] {
	case "help", "version":
		ShowHelpInfo()
		return
	case "clear":
		Clear()
		return
	case "stopweb":
		StopWebServer()
		return
	case "updatemod":
		UpdateMod()
		return
	case "init":
		PrepareGoEnv()
		SetupEnv()
		return
	}
	err := SetupEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	switch os.Args[1] {
	case "run", "editor", "export", "build":
		BuildDll()
	case "runweb", "buildweb":
		BuildDll()
		CheckExportWeb()
		BuildWasm()
	}
	switch os.Args[1] {
	case "run":
		err = Run("")
	case "editor":
		err = Run("-e")
	case "runweb":
		err = RunWebServer()
	case "exportweb":
		err = ExportWeb()
	case "export":
		err = Export()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
