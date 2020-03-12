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

type GstEvent C.GstEvent
type GstEventType C.GstEventType
type GstEventTypeFlags C.GstEventTypeFlags
type GstQOSType C.GstQOSType

func (e *C.GstEvent) impl() *GstEvent {
	return (*GstEvent)(unsafe.Pointer(e))
}

func (e *GstEvent) native() *C.GstEvent {
	return (*C.GstEvent)(unsafe.Pointer(e))
}
