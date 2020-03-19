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

type GstQuery C.GstQuery

type GstQueryType C.GstQueryType
type GstBufferPool C.GstBufferPool           //TODO: move this
type GstSchedulingFlags C.GstSchedulingFlags //TODO:MOVETHIS

func (s *C.GstQuery) impl() *GstQuery {
	return (*GstQuery)(unsafe.Pointer(s))
}

func (q *GstQuery) native() *C.GstQuery {
	return (*C.GstQuery)(unsafe.Pointer(q))
}

// Constructors

func QueryNewAcceptCaps(caps *GstCaps) *GstQuery {
	return C.gst_query_new_accept_caps(caps.native()).impl()
}

func QueryNewAllocation(caps *GstCaps, need_pool bool) *GstQuery {
	n := 0
	if need_pool {
		n = 1
	}
	return C.gst_query_new_allocation(caps.native(), C.gboolean(n)).impl()
}

func QueryNewBitrate() *GstQuery {
	return C.gst_query_new_bitrate().impl()
}

func QueryNewBuffering(format GstFormat) *GstQuery {
	return C.gst_query_new_buffering(C.GstFormat(format)).impl()
}

func QueryNewCaps(filter *GstCaps) *GstQuery {
	return C.gst_query_new_caps(filter.native()).impl()
}

func QueryNewContext(context_type string) *GstQuery {
	c := C.CString(context_type)
	defer C.free(unsafe.Pointer(c))
	return C.gst_query_new_context((*C.gchar)(c)).impl()
}

func QueryNewConvert(src_format, dest_format GstFormat, value int64) *GstQuery {
	return C.gst_query_new_convert(C.GstFormat(src_format), C.gint64(value), C.GstFormat(dest_format)).impl()
}

func QueryNewCustom(t GstQueryType, structure *GstStructure) *GstQuery {
	return C.gst_query_new_custom(C.GstQueryType(t), structure.native()).impl()
}

func QueryNewDrain() *GstQuery {
	return C.gst_query_new_drain().impl()
}

func QueryNewDuration(format GstFormat) *GstQuery {
	return C.gst_query_new_duration(C.GstFormat(format)).impl()
}

func QueryNewFormats() *GstQuery {
	return C.gst_query_new_formats().impl()
}

func QueryNewLatency() *GstQuery {
	return C.gst_query_new_latency().impl()
}

func QueryNewPosition(format GstFormat) *GstQuery {
	return C.gst_query_new_position(C.GstFormat(format)).impl()
}

func QueryNewScheduling() *GstQuery {
	return C.gst_query_new_scheduling().impl()
}

func QueryNewSeeking(format GstFormat) *GstQuery {
	return C.gst_query_new_seeking(C.GstFormat(format)).impl()
}

func QueryNewSegment(format GstFormat) *GstQuery {
	return C.gst_query_new_segment(C.GstFormat(format)).impl()
}

func QueryNewURI() *GstQuery {
	return C.gst_query_new_uri().impl()
}

// Methods

//AddAllocationMeta: Gtype

/*
func (q *GstQuery) AddAllocationPool(pool *GstBufferPool, size, min_buffer, max_buffer uint) {
	C.gst_query_add_allocation_pool(q.native(), pool.native(), C.guint(size), C.guint(min_buffer), C.guint(max_buffer))
}
*/

func (q *GstQuery) AddBufferingRange(start, stop uint) bool {
	return C.gst_query_add_buffering_range(q.native(), C.long(start), C.long(stop)) != 0
}

func (q *GstQuery) AddSchedulingMode(mode GstPadMode) {
	C.gst_query_add_scheduling_mode(q.native(), C.GstPadMode(mode))
}

func (q *GstQuery) Copy() *GstQuery {
	return C.gst_query_copy(q.native()).impl()
}

//FindAllocationMeta: GType

func (q *GstQuery) GetNAllocationMetas() uint {
	return uint(C.gst_query_get_n_allocation_metas(q.native()))
}

func (q *GstQuery) GetNAllocationParams() uint {
	return uint(C.gst_query_get_n_allocation_params(q.native()))
}

func (q *GstQuery) GetNAllocationPools() uint {
	return uint(C.gst_query_get_n_allocation_pools(q.native()))
}

func (q *GstQuery) GetNBufferingRanges() uint {
	return uint(C.gst_query_get_n_buffering_ranges(q.native()))
}

func (q *GstQuery) GetNSchedulingModes() uint {
	return uint(C.gst_query_get_n_scheduling_modes(q.native()))
}

func (q *GstQuery) GetStructure() *GstStructure {
	return C.gst_query_get_structure(q.native()).impl()
}

func (q *GstQuery) HasSchedulingMode(mode GstPadMode) bool {
	return C.gst_query_has_scheduling_mode(q.native(), C.GstPadMode(mode)) != 0
}

func (q *GstQuery) HasSchedulingModeWithFlags(mode GstPadMode, flags GstSchedulingFlags) bool {
	return C.gst_query_has_scheduling_mode_with_flags(q.native(), C.GstPadMode(mode), C.GstSchedulingFlags(flags)) != 0
}

func (q *GstQuery) ParseAcceptCaps() []*GstCaps {
	//TODO: Rethink this
	var first *GstCaps
	var caps []*GstCaps

	start := unsafe.Pointer(&first)          // **C.GstCaps
	size := unsafe.Sizeof((*C.GstCaps)(nil)) // size of a pointer

	C.gst_query_parse_accept_caps(q.native(), (**C.GstCaps)(start))

	for i := 0; ; i++ {
		item := (*GstCaps)(unsafe.Pointer(uintptr(start) + size*uintptr(i)))
		if item == nil {
			break
		}
		caps = append(caps, item)
	}
	return caps
}

func (q *GstQuery) ParseAcceptCapsResult() bool {
	var result C.gboolean
	C.gst_query_parse_accept_caps_result(q.native(), &result)
	return result != 0
}

func (q *GstQuery) ParseAllocation() {
}

//func (q *GstQuery)
