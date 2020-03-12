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

type GstCaps C.GstCaps

// GstCapsIntersectMode
type GstCapsIntersectMode C.GstCapsIntersectMode

const (
	GST_CAPS_INTERSECT_ZIG_ZAG GstCapsIntersectMode = 0
	GST_CAPS_INTERSECT_FIRST   GstCapsIntersectMode = 1
)

func (s *C.GstCaps) impl() *GstCaps {
	return (*GstCaps)(s)
}

func (p *GstCaps) native() *C.GstCaps {
	return (*C.GstCaps)(p)
}

// Constructors
func CapsNewAny() *GstCaps {
	return C.gst_caps_new_any().impl()
}

func CapsNewEmpty() *GstCaps {
	return C.gst_caps_new_empty().impl()
}

func CapsNewEmptySimple(media_type string) *GstCaps {
	m := C.CString(media_type)
	defer C.free(unsafe.Pointer(m))
	return (*GstCaps)(C.gst_caps_new_empty_simple(m))
}

/*
func CapsNewFull(struct1 ...*GstStructure) *GstCaps {
}
*/

func CapsFromString(str string) *GstCaps {
	s := C.CString(str)
	defer C.free(unsafe.Pointer(s))
	return C.gst_caps_from_string(s).impl()
}

// Methods

func (c *GstCaps) Append(c2 *GstCaps) {
	C.gst_caps_append(c.native(), c2.native())
}

func (c *GstCaps) AppendStructure(structure *GstStructure) {
	C.gst_caps_append_structure(c.native(), structure.native())
}

func (c *GstCaps) AppendStructureFull(structure *GstStructure, features *GstCapsFeatures) {
	C.gst_caps_append_structure_full(c.native(), structure.native(), features.native())
}

func (c *GstCaps) CanIntersect(c2 *GstCaps) bool {
	return C.gst_caps_can_intersect(c.native(), c2.native()) != 0
}

func (c *GstCaps) Copy() *GstCaps {
	return C.gst_caps_copy(c.native()).impl()
}

func (c *GstCaps) CopyNth(nth uint) *GstCaps {
	return C.gst_caps_copy_nth(c.native(), C.guint(nth)).impl()
}

func (c *GstCaps) Fixate() *GstCaps {
	return C.gst_caps_fixate(c.native()).impl()
}

func (c *GstCaps) GetFeatures(index uint) *GstCapsFeatures {
	return C.gst_caps_get_features(c.native(), C.guint(index)).impl()
}

func (c *GstCaps) GetSize() uint {
	return uint(C.gst_caps_get_size(c.native()))
}

func (c *GstCaps) GetStructure(index uint) *GstStructure {
	return C.gst_caps_get_structure(c.native(), C.guint(index)).impl()
}

func (c *GstCaps) Intersect(c2 *GstCaps) *GstCaps {
	return C.gst_caps_intersect(c.native(), c2.native()).impl()
}

func (c *GstCaps) IntersectFull(c2 *GstCaps, mod GstCapsIntersectMode) *GstCaps {
	return C.gst_caps_intersect_full(c.native(), c2.native(), C.GstCapsIntersectMode(mod)).impl()
}

func (c *GstCaps) IsAlwaysCampatiable(c2 *GstCaps) bool {
	return C.gst_caps_is_always_compatible(c.native(), c2.native()) != 0
}

func (c *GstCaps) IsAny() bool {
	return C.gst_caps_is_any(c.native()) != 0
}

func (c *GstCaps) IsEmpty() bool {
	return C.gst_caps_is_empty(c.native()) != 0
}

func (c *GstCaps) IsEqual(c2 *GstCaps) bool {
	return C.gst_caps_is_equal(c.native(), c2.native()) != 0
}

func (c *GstCaps) IsEqualFixed(c2 *GstCaps) bool {
	return C.gst_caps_is_equal_fixed(c.native(), c2.native()) != 0
}

func (c *GstCaps) IsFixed() bool {
	return C.gst_caps_is_fixed(c.native()) != 0
}

func (c *GstCaps) IsStrictlyEqual(c2 *GstCaps) bool {
	return C.gst_caps_is_strictly_equal(c.native(), c2.native()) != 0
}

func (c *GstCaps) IsSubset(c2 *GstCaps) bool {
	return C.gst_caps_is_subset(c.native(), c2.native()) != 0
}

func (c *GstCaps) IsSubsetStructure(structure *GstStructure) bool {
	return C.gst_caps_is_subset_structure(c.native(), structure.native()) != 0
}

func (c *GstCaps) IsSubsetStructureFull(structure *GstStructure, features *GstCapsFeatures) bool {
	return C.gst_caps_is_subset_structure_full(c.native(), structure.native(), features.native()) != 0
}

func (c *GstCaps) Merge(c2 *GstCaps) *GstCaps {
	return C.gst_caps_merge(c.native(), c2.native()).impl()
}

func (c *GstCaps) MergeStructure(structure *GstStructure) *GstCaps {
	return C.gst_caps_merge_structure(c.native(), structure.native()).impl()
}

func (c *GstCaps) MergeStructureFull(structure *GstStructure, features *GstCapsFeatures) *GstCaps {
	return C.gst_caps_merge_structure_full(c.native(), structure.native(), features.native()).impl()
}

func (c *GstCaps) Normalize() *GstCaps {
	return C.gst_caps_normalize(c.native()).impl()
}

func (c *GstCaps) Ref() *GstCaps {
	return C.gst_caps_ref(c.native()).impl()
}

func (c *GstCaps) RemoveStructure(index uint) {
	C.gst_caps_remove_structure(c.native(), C.guint(index))
}

func (c *GstCaps) SetFeatures(index uint, features *GstCapsFeatures) {
	C.gst_caps_set_features(c.native(), C.guint(index), features.native())
}

func (c *GstCaps) SetFeaturesSimple(features *GstCapsFeatures) {
	C.gst_caps_set_features_simple(c.native(), features.native())
}

/*
func (c *GstCaps) SetSimple()
*/

/*
func (c *GetCaps) SetValue(field string, )
*/

func (c *GstCaps) Simplify() *GstCaps {
	return C.gst_caps_simplify(c.native()).impl()
}

func (c *GstCaps) StealStructure(index uint) *GstStructure {
	return C.gst_caps_steal_structure(c.native(), C.guint(index)).impl()
}

func (c *GstCaps) Subtract(subtrahend *GstCaps) *GstCaps {
	return C.gst_caps_subtract(c.native(), subtrahend.native()).impl()
}

func (c *GstCaps) ToString() string {
	return C.GoString(C.gst_caps_to_string(c.native()))
}

func (c *GstCaps) Truncate() *GstCaps {
	return C.gst_caps_truncate(c.native()).impl()
}

func (c *GstCaps) Unref() {
	C.gst_caps_unref(c.native())
}
