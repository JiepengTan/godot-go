package ffi

import (
	"unsafe"

	gdx "github.com/JiepengTan/godotgo"
)
import "C"

const (
	GDExtensionInitializationLevelCore    GDExtensionInitializationLevel = 0
	GDExtensionInitializationLevelServers GDExtensionInitializationLevel = 1
	GDExtensionInitializationLevelScene   GDExtensionInitializationLevel = 2
	GDExtensionInitializationLevelEditor  GDExtensionInitializationLevel = 3
)

var (
	dlsymGD   func(string) unsafe.Pointer
	callbacks gdx.CallbackInfo
)

//go:linkname main main.main
func main()

func Link() bool {
	return false
}
func Linked() {
}
func BindCallback(info gdx.CallbackInfo) {
	callbacks = info
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	dlsymGD = func(s string) unsafe.Pointer {
		ptr := getProcAddress(lookupFunc, s)
		if ptr == nil {
			println("can not getProcAddress ", s)
		}
		return ptr
	}
	api.loadProcAddresses()
	bindFFI()
	init := (*initialization)(configuration)
	*init = initialization{}
	init.minimum_initialization_level = initializationLevel(GDExtensionInitializationLevelScene)
	doInitialization(init)
	return 1
}
