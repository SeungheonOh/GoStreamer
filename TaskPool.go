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

type GstTaskPool struct {
	GstObject
}

func (t *C.GstTaskPool) impl() *GstTaskPool {
	return (*GstTaskPool)(unsafe.Pointer(t))
}

func (t *GstTaskPool) native() *C.GstTaskPool {
	return (*C.GstTaskPool)(unsafe.Pointer(t))
}

// Constructors

func TaskPoolNew() *GstTaskPool {
	return C.gst_task_pool_new().impl()
}

func (p *GstTaskPool) Cleanup() {
	C.gst_task_pool_cleanup(p.native())
}

//Join: gpointer

func (p *GstTaskPool) Prepare() {
	//TODO: GError
	C.gst_task_pool_prepare(p.native(), nil)
}

// Push: callback, gpointer

//func (p *GstTaskPool)
