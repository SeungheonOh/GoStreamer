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

type GstEvent C.GstEvent

func (e *GstEvent) native() *C.GstEvent {
	return (*C.GstEvent)(e)
}
