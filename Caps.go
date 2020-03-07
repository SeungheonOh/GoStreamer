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

type GstCaps C.GstCaps

func (p *GstCaps) native() *C.GstCaps {
	return (*C.GstCaps)(p)
}
