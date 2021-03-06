package gostreamer

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstStructure C.GstStructure

func (s *C.GstStructure) impl() *GstStructure {
	return (*GstStructure)(s)
}

func (s *GstStructure) native() *C.GstStructure {
	return (*C.GstStructure)(s)
}
