package gostreamer

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstPad C.GstPad

func (s *C.GstPad) impl() *GstPad {
	return (*GstPad)(s)
}

func (p *GstPad) native() *C.GstPad {
	return (*C.GstPad)(p)
}
