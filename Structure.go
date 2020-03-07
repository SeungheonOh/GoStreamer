package main

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstStructure C.GstStructure

func (s *GstStructure) native() *C.GstStructure {
	return (*C.GstStructure)(s)
}
