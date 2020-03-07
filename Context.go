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

type GstContext C.GstContext

func (s *C.GstContext) impl() *GstContext {
	return (*GstContext)(s)
}

func (c *GstContext) native() *C.GstContext {
	return (*C.GstContext)(c)
}
