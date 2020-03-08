package gostreamer

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstPad C.GstPad

type GstFlowReturn C.GstFlowReturn
type GstPadDirection C.GstPadDirection
type GstPadFlags C.GstPadFlags
type GstPadLinkCheck C.GstPadLinkCheck
type GstPadLinkReturn C.GstPadLinkReturn
type GstPadMode C.GstPadMode
type GstPadProbeReturn C.GstPadProbeReturn
type GstPadProbeType C.GstPadProbeType

func (s *C.GstPad) impl() *GstPad {
	return (*GstPad)(s)
}

func (p *GstPad) native() *C.GstPad {
	return (*C.GstPad)(p)
}
