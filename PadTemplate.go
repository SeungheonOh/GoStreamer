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

type GstPadTemplate C.GstPadTemplate

func (s *C.GstPadTemplate) impl() *GstPadTemplate {
	return (*GstPadTemplate)(s)
}

func (p *GstPadTemplate) native() *C.GstPadTemplate {
	return (*C.GstPadTemplate)(p)
}
