package gostreamer

import (
	"os"
	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

func GstInit(argv *[]string) {
	if argv == nil {
		C.gst_init(nil, nil)
		return
	}
	argc := len(*argv)

	c_argv := C.malloc(C.ulong(4 * (argc)))
	for i, s := range *argv {
		arg_s := (*C.char)(C.CString(s))
		if arg_s == nil {
			continue
		}
		*(**C.char)(unsafe.Pointer(uintptr(c_argv) + uintptr(i))) = arg_s
	}

	C.gst_init((*C.int)(unsafe.Pointer(&argc)),
		(***C.char)(unsafe.Pointer(&c_argv)))

}

func GstEasyInit() {
	argv := os.Args
	GstInit(&argv)
}

func GstGetVersion() (int, int, int, int) {
	var major, minor, micro, nano C.guint
	C.gst_version(&major, &minor, &micro, &nano)
	return int(major), int(minor), int(micro), int(nano)
}
