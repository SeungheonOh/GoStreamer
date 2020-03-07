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

type GstCaps C.GstCaps

func (p *GstCaps) native() *C.GstCaps {
	return (*C.GstCaps)(p)
}

// Constructors
func CapsNewAny() *GstCaps {
	c := C.gst_caps_new_any()
	return (*GstCaps)(c)
}

func CapsNewEmpty() *GstCaps {
	//c := C.gst_caps_new_any()
	return (*GstCaps)(C.gst_caps_new_any())
}

func CapsNewEmptySimple(media_type string) *GstCaps {
	c := C.gst_caps_new_any()
	return (*GstCaps)(c)
}
