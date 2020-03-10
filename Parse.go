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

type GstParseContext C.GstParseContext

type GstParseFlags C.GstParseFlags

const (
	GST_PARSE_FLAG_NONE                    GstParseFlags = 0
	GST_PARSE_FLAG_FATAL_ERRORS            GstParseFlags = 1
	GST_PARSE_FLAG_NO_SINGLE_ELEMENT_BINDS GstParseFlags = 2
	GST_PARSE_FLAG_PLACE_IN_BING           GstParseFlags = 4
)

func (p *C.GstParseContext) impl() *GstParseContext {
	return (*GstParseContext)(p)
}

func (p *GstParseContext) native() *C.GstParseContext {
	return (*C.GstParseContext)(p)
}

func ParseContextNew() *GstParseContext {
	return C.gst_parse_context_new().impl()
}

func (p *GstParseContext) Copy() *GstParseContext {
	return C.gst_parse_context_copy(p.native()).impl()
}

func (p *GstParseContext) Free() {
	C.gst_parse_context_free(p.native())
}

func (p *GstParseContext) GetMissingElements() []string {
	ret := []string{}
	cp := unsafe.Pointer(C.gst_parse_context_get_missing_elements(p.native()))
	defer C.g_strfreev((**C.gchar)(cp))
	i := 0
	for {
		if unsafe.Pointer(uintptr(cp)+uintptr(i)) == nil {
			break
		}
		ret = append(ret, C.GoString((*C.char)(unsafe.Pointer((uintptr(cp) + uintptr(i))))))
		i++
	}
	return ret
}

func ParseLaunch(pipeline_description string) *GstElement {
	p := C.CString(pipeline_description)
	defer C.free(unsafe.Pointer(p))
	return C.gst_parse_launch((*C.gchar)(p), nil).impl()
}

func ParseLaunchFull(pipeline_description string, context *GstParseContext, flags GstParseFlags) *GstElement {
	p := C.CString(pipeline_description)
	defer C.free(unsafe.Pointer(p))
	return C.gst_parse_launch_full((*C.gchar)(p), context.native(), C.GstParseFlags(flags), nil).impl()
}
