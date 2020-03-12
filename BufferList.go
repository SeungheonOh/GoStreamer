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

type GstBufferList C.GstBufferList

func (b *C.GstBufferList) impl() *GstBufferList {
	return (*GstBufferList)(unsafe.Pointer(b))
}

func (b *GstBufferList) native() *C.GstBufferList {
	return (*C.GstBufferList)(unsafe.Pointer(b))
}
