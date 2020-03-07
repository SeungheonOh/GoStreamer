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

type GstPadTemplate C.GstPadTemplate

func (p *GstPadTemplate) native() *C.GstPadTemplate {
	return (*C.GstPadTemplate)(p)
}
