package gostreamer

import (
	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstFormat C.GstFormat

func (f *C.GstFormat) impl() *GstFormat {
	return (*GstFormat)(unsafe.Pointer(f))
}

func (f *GstFormat) native() *C.GstFormat {
	return (*C.GstFormat)(unsafe.Pointer(f))
}
