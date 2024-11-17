package main

import (
	"unsafe"

	gdx "github.com/JiepengTan/godotgo"
)

import "C"

func main() {
	RegisterTypes()
	gdx.LinkEngine(gdx.EngineCallbacks{
		OnEngineStart:   onStart,
		OnEngineUpdate:  onUpdate,
		OnEngineDestroy: onDestory,
	})
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	println("hello godot link!")
	return 0
}

func RegisterTypes() {

}

func onStart() {
	println("hello world!")
}

func onUpdate(delta float32) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("goodbye world!")
}
