package gostreamer

import (
//	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstBus C.GstBus

func (s *C.GstBus) impl() *GstBus {
	return (*GstBus)(s)
}

func (b *GstBus) native() *C.GstBus {
	return (*C.GstBus)(b)
}
