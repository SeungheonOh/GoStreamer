package main

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

func (c *GstContext) native() *C.GstContext {
	return (*C.GstContext)(c)
}
