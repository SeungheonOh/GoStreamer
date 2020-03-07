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

type GstObject C.GstObject

func (s *C.GstObject) impl() *GstObject {
	return (*GstObject)(s)
}

func (o *GstObject) native() *C.GstObject {
	return (*C.GstObject)(o)
}

func (o *GstObject) SetName(name string) bool {
	s := C.CString(name)
	defer C.g_free((C.gpointer)(unsafe.Pointer(s)))
	return C.gst_object_set_name(o.native(), (*C.gchar)(s)) != 0
}

func (o *GstObject) GetName() string {
	s := C.gst_object_get_name(o.native())
	defer C.g_free((C.gpointer)(unsafe.Pointer(s)))
	if s == nil {
		return ""
	}
	return C.GoString((*C.char)(s))
}

func (o *GstObject) SetParent(parent *GstObject) bool {
	return C.gst_object_set_parent(o.native(), parent.native()) != 0
}

func (o *GstObject) GetParent() *GstObject {
	p := C.gst_object_get_parent(o.native())
	return (*GstObject)(p)
}

func (o *GstObject) GetPathString() string {
	s := C.gst_object_get_path_string(o.native())
	defer C.g_free((C.gpointer)(unsafe.Pointer(s)))
	if s == nil {
		return ""
	}
	return C.GoString((*C.char)(s))
}

func (o *GstObject) Unparent(parent *GstObject) {
	C.gst_object_unparent(parent.native())
}

func (o *GstObject) Unref() {
	C.gst_object_unref((C.gpointer)(o.native()))
}
