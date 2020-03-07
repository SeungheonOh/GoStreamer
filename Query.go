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

type GstQuery C.GstQuery

func (q *GstQuery) native() *C.GstQuery {
	return (*C.GstQuery)(q)
}
