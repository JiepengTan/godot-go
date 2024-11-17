package main

import (
	gdx "github.com/JiepengTan/godotgo"
	"github.com/JiepengTan/godotgo/extension"
)

import "C"

func main() {
	extension.LinkEngine(gdx.EngineCallbacks{})
	println("hello world")
}
