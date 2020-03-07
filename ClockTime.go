package gostreamer

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstClockTime C.GstClockTime

func (s *C.GstClockTime) impl() *GstClockTime {
	return (*GstClockTime)(s)
}

func (c *GstClockTime) native() *C.GstClockTime {
	return (*C.GstClockTime)(c)
}
