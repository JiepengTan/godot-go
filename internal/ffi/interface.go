package ffi

//#include "ffi_wrapper.gen.h"
import "C"

func ToGdBool(value bool) C.GDExtensionBool {
	if value {
		return 1
	} else {
		return 0
	}
}
