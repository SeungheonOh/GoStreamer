package gostreamer

import (
	"os"
	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
#include <stdio.h>
*/
import "C"

func GstInit(argv *[]string) {
	if argv == nil {
		C.gst_init(nil, nil)
		return
	}
	argc := len(*argv)

	c_argv := C.malloc(C.ulong(8 * (argc)))
	for i, s := range *argv {
		arg_s := (*C.char)(C.CString(s))
		if arg_s == nil {
			continue
		}
		*(**C.char)(unsafe.Pointer(uintptr(c_argv) + uintptr(i*8))) = arg_s
	}

	C.gst_init((*C.int)(unsafe.Pointer(&argc)),
		(***C.char)(unsafe.Pointer(&c_argv)))
}

func GstEasyInit() {
	argv := os.Args
	GstInit(&argv)
}

func GstDeinit() {
	C.gst_deinit()
}

func GstInitCheck(argv *[]string) bool {
	if argv == nil {
		return false
	}
	argc := len(*argv)

	c_argv := C.malloc(C.ulong(8 * (argc)))
	for i, s := range *argv {
		arg_s := (*C.char)(C.CString(s))
		if arg_s == nil {
			continue
		}
		*(**C.char)(unsafe.Pointer(uintptr(c_argv) + uintptr(i*8))) = arg_s
	}

	return C.gst_init_check((*C.int)(unsafe.Pointer(&argc)),
		(***C.char)(unsafe.Pointer(&c_argv)), nil) != 0
}

func GstIsInitialized() bool {
	return C.gst_is_initialized() != 0
}

func GstVersion() (int, int, int, int) {
	var major, minor, micro, nano C.guint
	C.gst_version(&major, &minor, &micro, &nano)
	return int(major), int(minor), int(micro), int(nano)
}

func GstVersionString() string {
	return C.GoString(C.gst_version_string())
}
