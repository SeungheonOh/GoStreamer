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

type GstTask C.GstTask
type GstTaskState C.GstTaskState

func (t *C.GstTask) impl() *GstTask {
	return (*GstTask)(unsafe.Pointer(t))
}

func (t *GstTask) native() *C.GstTask {
	return (*C.GstTask)(unsafe.Pointer(t))
}
