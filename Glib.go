package gostreamer

import (
//"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>

int Is_obj(GObject *o) {
	return G_IS_OBJECT(o);
}
*/
import "C"

//type GQuark C.GQuark //uint32

//type GObject C.GObject

/*
func (g *GObject) native() *C.GObject {
	return (*C.GObject)(unsafe.Pointer(g))
}

func (g *C.GObject) impl() *GObject {
	return (*GObject)(unsafe.Pointer(g))
}

func (g *GObject) ObjectSet(name, value string) {
	n := C.CString(name)
	s := C.CString(value)
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(s))
	C.g_object_set(g.native(), (*C.gchar)(n), (*C.gchar)(s), nil)
}
*/
