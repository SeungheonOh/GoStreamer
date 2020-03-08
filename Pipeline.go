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

type GstPipeline C.GstPipeline

func (p *C.GstPipeline) impl() *GstPipeline {
	return (*GstPipeline)(p)
}

func (p *GstPipeline) native() *C.GstPipeline {
	return (*C.GstPipeline)(p)
}

func PipelineNew(name string) *GstElement {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_pipeline_new((*C.gchar)(n)).impl()
}
