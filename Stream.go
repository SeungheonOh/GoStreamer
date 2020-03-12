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

type GstStream C.GstStream

func (s *C.GstStream) impl() *GstStream {
	return (*GstStream)(unsafe.Pointer(s))
}

func (s *GstStream) native() *C.GstStream {
	return (*C.GstStream)(unsafe.Pointer(s))
}
