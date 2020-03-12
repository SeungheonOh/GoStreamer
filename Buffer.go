package gostreamer

import (
//"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstBuffer C.GstBuffer
type GstBufferCopyFlags C.GstBufferCopyFlags
type GstBufferFlags C.GstBufferFlags

func (s *C.GstBuffer) impl() *GstBuffer {
	return (*GstBuffer)(s)
}

func (p *GstBuffer) native() *C.GstBuffer {
	return (*C.GstBuffer)(p)
}
