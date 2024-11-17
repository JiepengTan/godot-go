package main

import (
	"unsafe"

	gdx "github.com/JiepengTan/godotgo"
)

import "C"

func main() {
	gdx.LinkEngine(gdx.EngineCallbacks{})
	println("hello world")
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	println("hello godot link!")
	return 0
}
