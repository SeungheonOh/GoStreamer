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

type GstElementFactory C.GstElementFactory

type GstElementFactoryListType C.GstElementFactoryListType

func (s *C.GstElementFactory) impl() *GstElementFactory {
	return (*GstElementFactory)(s)
}

func (e *GstElementFactory) native() *C.GstElementFactory {
	return (*C.GstElementFactory)(e)
}

// Function
func ElementFactoryFind(name string) *GstElementFactory {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_element_factory_find((*C.gchar)(n)).impl()
}

/*
func ElementFactoryListFilter()
*/

func ElementFactoryMake(factory, name string) *GstElement {
	f := C.CString(factory)
	n := C.CString(name)
	defer C.free(unsafe.Pointer(f))
	defer C.free(unsafe.Pointer(n))
	return C.gst_element_factory_make((*C.gchar)(f), (*C.gchar)(n)).impl()
}

// Method

func (e *GstElementFactory) CanSinkAllCaps(caps *GstCaps) bool {
	return C.gst_element_factory_can_sink_all_caps(e.native(), caps.native()) != 0
}

func (e *GstElementFactory) CanSrcAllCaps(caps *GstCaps) bool {
	return C.gst_element_factory_can_src_all_caps(e.native(), caps.native()) != 0
}

func (e *GstElementFactory) CanSrcAnyCaps(caps *GstCaps) bool {
	return C.gst_element_factory_can_src_any_caps(e.native(), caps.native()) != 0
}

func (e *GstElementFactory) Create(name string) *GstElement {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_element_factory_create(e.native(), (*C.gchar)(n)).impl()
}

func (e *GstElementFactory) GetElementType() uint { // GType == uint
	return uint(C.gst_element_factory_get_element_type(e.native()))
}

func (e *GstElementFactory) GetMetadata(key string) string {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	return C.GoString(C.gst_element_factory_get_metadata(e.native(), (*C.gchar)(k)))
}

/*
func (e *GstElementFactory) GetMetadatKeys() []string {
	keys := C.gst_element_factory_get_metadata_keys(e.native())
	// IDK
}
*/

func (e *GstElementFactory) GetNumPadTemplates() int {
	return int(C.gst_element_factory_get_num_pad_templates(e.native()))
}

/*
func (e *GstElementFactory) GetStaticPadTemplates() {
}
func (e *GstElementFactory) GetURIProtocols() {
}
*/

func (e *GstElementFactory) GetURIType() GstURIType {
	return GstURIType(C.gst_element_factory_get_uri_type(e.native()))
}

func (e *GstElementFactory) HasInterface(interfacename string) bool {
	i := C.CString(interfacename)
	defer C.free(unsafe.Pointer(i))
	return C.gst_element_factory_has_interface(e.native(), (*C.gchar)(i)) != 0
}

func (e *GstElementFactory) ListIsType(t GstElementFactoryListType) bool {
	return C.gst_element_factory_list_is_type(e.native(), C.GstElementFactoryListType(t)) != 0
}
