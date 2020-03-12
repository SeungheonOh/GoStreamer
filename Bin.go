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

type GstBin struct {
	GstElement
}

func (b *C.GstBin) impl() *GstBin {
	return (*GstBin)(unsafe.Pointer(b))
}

func (b *GstBin) native() *C.GstBin {
	return (*C.GstBin)(unsafe.Pointer(b))
}

func (b *GstBin) New(name string) *GstElement {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_bin_new((*C.gchar)(n)).impl()
}

func (b *GstBin) Add(elements ...*GstElement) bool {
	for _, e := range elements {
		if C.gst_bin_add(b.native(), e.native()) == 0 {
			return false
		}
	}
	return true
}

func (b *GstBin) FindUnlinkedPad(direction GstPadDirection) *GstPad {
	return C.gst_bin_find_unlinked_pad(b.native(), C.GstPadDirection(direction)).impl()
}

/*
func (b *GstBin) GetByInterface
*/

func (b *GstBin) GetByNameRecurseUp(name string) *GstElement {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_bin_get_by_name_recurse_up(b.native(), (*C.gchar)(n)).impl()
}

func (b *GstBin) GetSuppressed_flags() GstElementFlags {
	return GstElementFlags(C.gst_bin_get_suppressed_flags(b.native()))
}

// A lot of iterate_??? methods (Skipped)

func (b *GstBin) RecalculateLatency() bool {
	return C.gst_bin_recalculate_latency(b.native()) != 0
}

func (b *GstBin) Remove(elements ...*GstElement) {
	for _, e := range elements {
		C.gst_bin_remove(b.native(), e.native())
	}
}

func (b *GstBin) SetSuppressedFlags(flags GstElementFlags) {
	C.gst_bin_set_suppressed_flags(b.native(), C.GstElementFlags(flags))
}

func (b *GstBin) SyncChildrenStates() bool {
	return C.gst_bin_sync_children_states(b.native()) != 0
}
