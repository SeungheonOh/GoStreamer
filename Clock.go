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

type GstClock C.GstClock

func (e *GstClock) native() *C.GstClock {
	return (*C.GstClock)(e)
}
