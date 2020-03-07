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

type GstCapsFeatures C.GstCapsFeatures

func (s *C.GstCapsFeatures) impl() *GstCapsFeatures {
	return (*GstCapsFeatures)(s)
}

func (b *GstCapsFeatures) native() *C.GstCapsFeatures {
	return (*C.GstCapsFeatures)(b)
}
