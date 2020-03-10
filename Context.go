package gostreamer

import (
	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstContext C.GstContext

func (c *C.GstContext) impl() *GstContext {
	return (*GstContext)(c)
}

func (c *GstContext) native() *C.GstContext {
	return (*C.GstContext)(c)
}

func ContextNew(context_type string, persistent bool) *GstContext {
	p := 0
	if persistent {
		p = 1
	}
	c := C.CString(context_type)
	defer C.free(unsafe.Pointer(c))
	return C.gst_context_new((*C.gchar)(c), C.gboolean(p)).impl()
}

func (c *GstContext) GetContextType() string {
	return C.GoString(C.gst_context_get_context_type(c.native()))
}

func (c *GstContext) GetStructure() *GstStructure {
	return C.gst_context_get_structure(c.native()).impl()
}

func (c *GstContext) HasContextType(context_type string) bool {
	ct := C.CString(context_type)
	defer C.free(unsafe.Pointer(ct))
	return C.gst_context_has_context_type(c.native(), (*C.gchar)(ct)) != 0
}

func (c *GstContext) IsPersistent() bool {
	return C.gst_context_is_persistent(c.native()) != 0
}

func (c *GstContext) WritableStructure() *GstStructure {
	return C.gst_context_writable_structure(c.native()).impl()
}
