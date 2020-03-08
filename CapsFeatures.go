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

type GstCapsFeatures C.GstCapsFeatures

func (s *C.GstCapsFeatures) impl() *GstCapsFeatures {
	return (*GstCapsFeatures)(s)
}

func (b *GstCapsFeatures) native() *C.GstCapsFeatures {
	return (*C.GstCapsFeatures)(b)
}

// Constructor

func CapsFeaturesNewAny() *GstCapsFeatures {
	return C.gst_caps_features_new_any().impl()
}

func CapsFeaturesNewEmpty() *GstCapsFeatures {
	return C.gst_caps_features_new_empty().impl()
}

func CapsFeaturesFromString(features string) *GstCapsFeatures {
	f := C.CString(features)
	defer C.free(unsafe.Pointer(f))
	return C.gst_caps_features_from_string((*C.gchar)(f)).impl()
}

// missing some TODO: Deal with va_list

func (c *GstCapsFeatures) Add(feature string) {
	f := C.CString(feature)
	defer C.free(unsafe.Pointer(f))
	C.gst_caps_features_add(c.native(), (*C.gchar)(f))
}

func (c *GstCapsFeatures) AddId(feature uint32) {
	C.gst_caps_features_add_id(c.native(), C.GQuark(feature))
}

func (c *GstCapsFeatures) Contain(feature string) bool {
	f := C.CString(feature)
	defer C.free(unsafe.Pointer(f))
	return C.gst_caps_features_contains(c.native(), (*C.gchar)(f)) != 0
}

func (c *GstCapsFeatures) ContainId(feature uint32) bool {
	return C.gst_caps_features_contains_id(c.native(), C.GQuark(feature)) != 0
}

func (c *GstCapsFeatures) Copy() *GstCapsFeatures {
	return C.gst_caps_features_copy(c.native()).impl()
}
func (c *GstCapsFeatures) Free() {
	C.gst_caps_features_free(c.native())
}
func (c *GstCapsFeatures) GetNth(nth int) string {
	return C.GoString(C.gst_caps_features_get_nth(c.native(), C.guint(nth)))
}

func (c *GstCapsFeatures) GetNthId(nth int) uint32 {
	return uint32(C.gst_caps_features_get_nth_id(c.native(), C.guint(nth)))
}

func (c *GstCapsFeatures) GetSize() int {
	return int(C.gst_caps_features_get_size(c.native()))
}

func (c *GstCapsFeatures) IsAny() bool {
	return C.gst_caps_features_is_any(c.native()) != 0
}

func (c *GstCapsFeatures) IsEqual(c2 *GstCapsFeatures) bool {
	return C.gst_caps_features_is_equal(c.native(), c2.native()) != 0
}

func (c *GstCapsFeatures) Remove(feature string) {
	f := C.CString(feature)
	defer C.free(unsafe.Pointer(f))
	C.gst_caps_features_remove(c.native(), (*C.gchar)(f))
}

func (c *GstCapsFeatures) RemoveId(feature uint32) {
	C.gst_caps_features_remove_id(c.native(), C.GQuark(feature))
}

func (c *GstCapsFeatures) SetParentRefcount(refcount int) bool {
	return C.gst_caps_features_set_parent_refcount(c.native(), (*C.int)(unsafe.Pointer(&refcount))) != 0 // <- :/
}

func (c *GstCapsFeatures) ToString() string {
	return C.GoString(C.gst_caps_features_to_string(c.native()))
}
