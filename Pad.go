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

type GstPad C.GstPad

type GstFlowReturn C.GstFlowReturn
type GstPadDirection C.GstPadDirection
type GstPadFlags C.GstPadFlags
type GstPadLinkCheck C.GstPadLinkCheck
type GstPadLinkReturn C.GstPadLinkReturn
type GstPadMode C.GstPadMode
type GstPadProbeReturn C.GstPadProbeReturn
type GstPadProbeType C.GstPadProbeType

const (
	GST_FLOW_CUSTOM_SUCCESS_2 GstFlowReturn = 102
	GST_FLOW_CUSTOM_SUCCESS_1 GstFlowReturn = 101
	GST_FLOW_CUSTOM_SUCCESS   GstFlowReturn = 100
	GST_FLOW_OK               GstFlowReturn = 0
	GST_FLOW_NOT_LINKED       GstFlowReturn = -1
	GST_FLOW_FLUSHING         GstFlowReturn = -2
	GST_FLOW_EOS              GstFlowReturn = -3
	GST_FLOW_NOT_NEGOTIATED   GstFlowReturn = -4
	GST_FLOW_ERROR            GstFlowReturn = -5
	GST_FLOW_NOT_SUPPORTED    GstFlowReturn = -6
	GST_FLOW_CUSTOM_ERROR     GstFlowReturn = -100
	GST_FLOW_CUSTOM_ERROR_1   GstFlowReturn = -101
	GST_FLOW_CUSTOM_ERROR_2   GstFlowReturn = -102

	GST_PAD_UNKNOWN GstPadDirection = 0
	GST_PAD_SRC     GstPadDirection = 1
	GST_PAD_SINK    GstPadDirection = 2

	GST_PAD_FLAG_BLOCKED          GstPadFlags = 16
	GST_PAD_FLAG_FLUSHING         GstPadFlags = 32
	GST_PAD_FLAG_EOS              GstPadFlags = 64
	GST_PAD_FLAG_BLOCKING         GstPadFlags = 128
	GST_PAD_FLAG_NEED_PARENT      GstPadFlags = 256
	GST_PAD_FLAG_NEED_RECONFIGURE GstPadFlags = 512
	GST_PAD_FLAG_PENDING_EVENTS   GstPadFlags = 1024
	GST_PAD_FLAG_FIXED_CAPS       GstPadFlags = 2048
	GST_PAD_FLAG_PROXY_CAPS       GstPadFlags = 4096
	GST_PAD_FLAG_PROXY_ALLOCATION GstPadFlags = 8192
	GST_PAD_FLAG_PROXY_SCHEDULING GstPadFlags = 16384
	GST_PAD_FLAG_ACCEPT_INTERSECT GstPadFlags = 32768
	GST_PAD_FLAG_ACCEPT_TEMPLATE  GstPadFlags = 65536
	GST_PAD_FLAG_LAST             GstPadFlags = 1048576

	GST_PAD_LINK_CHECK_NOTHING        GstPadLinkCheck = 0
	GST_PAD_LINK_CHECK_HIERARCHY      GstPadLinkCheck = 1
	GST_PAD_LINK_CHECK_TEMPLATE_CAPS  GstPadLinkCheck = 2
	GST_PAD_LINK_CHECK_CAPS           GstPadLinkCheck = 4
	GST_PAD_LINK_CHECK_NO_RECONFIGURE GstPadLinkCheck = 8
	GST_PAD_LINK_CHECK_DEFAULT        GstPadLinkCheck = 5

	GST_PAD_LINK_OK              GstPadLinkReturn = 0
	GST_PAD_LINK_WRONG_HIERARCHY GstPadLinkReturn = -1
	GST_PAD_LINK_WAS_LINKED      GstPadLinkReturn = -2
	GST_PAD_LINK_WRONG_DIRECTION GstPadLinkReturn = -3
	GST_PAD_LINK_NOFORMAT        GstPadLinkReturn = -4
	GST_PAD_LINK_NOSCHED         GstPadLinkReturn = -5
	GST_PAD_LINK_REFUSED         GstPadLinkReturn = -6

	GST_PAD_MODE_NONE GstPadMode = 0
	GST_PAD_MODE_PUSH GstPadMode = 1
	GST_PAD_MODE_PULL GstPadMode = 2

	GST_PAD_PROBE_DROP    GstPadProbeReturn = 0
	GST_PAD_PROBE_OK      GstPadProbeReturn = 1
	GST_PAD_PROBE_REMOVE  GstPadProbeReturn = 2
	GST_PAD_PROBE_PASS    GstPadProbeReturn = 3
	GST_PAD_PROBE_HANDLED GstPadProbeReturn = 4

	GST_PAD_PROBE_TYPE_INVALID          GstPadProbeType = 0
	GST_PAD_PROBE_TYPE_IDLE             GstPadProbeType = 1
	GST_PAD_PROBE_TYPE_BLOCK            GstPadProbeType = 2
	GST_PAD_PROBE_TYPE_BUFFER           GstPadProbeType = 16
	GST_PAD_PROBE_TYPE_BUFFER_LIST      GstPadProbeType = 32
	GST_PAD_PROBE_TYPE_EVENT_DOWNSTREAM GstPadProbeType = 64
	GST_PAD_PROBE_TYPE_EVENT_UPSTREAM   GstPadProbeType = 128
	GST_PAD_PROBE_TYPE_EVENT_FLUSH      GstPadProbeType = 256
	GST_PAD_PROBE_TYPE_QUERY_DOWNSTREAM GstPadProbeType = 512
	GST_PAD_PROBE_TYPE_QUERY_UPSTREAM   GstPadProbeType = 1024
	GST_PAD_PROBE_TYPE_PUSH             GstPadProbeType = 4096
	GST_PAD_PROBE_TYPE_PULL             GstPadProbeType = 8192
	GST_PAD_PROBE_TYPE_BLOCKING         GstPadProbeType = 3
	GST_PAD_PROBE_TYPE_DATA_DOWNSTREAM  GstPadProbeType = 112
	GST_PAD_PROBE_TYPE_DATA_UPSTREAM    GstPadProbeType = 128
	GST_PAD_PROBE_TYPE_DATA_BOTH        GstPadProbeType = 240
	GST_PAD_PROBE_TYPE_BLOCK_DOWNSTREAM GstPadProbeType = 114
	GST_PAD_PROBE_TYPE_BLOCK_UPSTREAM   GstPadProbeType = 130
	GST_PAD_PROBE_TYPE_EVENT_BOTH       GstPadProbeType = 192
	GST_PAD_PROBE_TYPE_QUERY_BOTH       GstPadProbeType = 1536
	GST_PAD_PROBE_TYPE_ALL_BOTH         GstPadProbeType = 1776
	GST_PAD_PROBE_TYPE_SCHEDULING       GstPadProbeType = 12288
)

func (s *C.GstPad) impl() *GstPad {
	return (*GstPad)(s)
}

func (p *GstPad) native() *C.GstPad {
	return (*C.GstPad)(p)
}

//Constructor

func PadNew(name string, direction GstPadDirection) *GstPad {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_pad_new((*C.gchar)(n), C.GstPadDirection(direction)).impl()
}

func PadNewFromStaticTemplate(templ *GstStaticPadTemplate, name string) *GstPad {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_pad_new_from_static_template(templ.native(), (*C.gchar)(n)).impl()
}

func PadNewFromTemplate(templ *GstPadTemplate, name string) *GstPad {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return C.gst_pad_new_from_template(templ.native(), (*C.gchar)(n)).impl()
}

//Methods

func (p *GstPad) ActivateMode(mode GstPadMode, active bool) bool {
	i := 0
	if active {
		i = 1
	}
	return C.gst_pad_activate_mode(p.native(), C.GstPadMode(mode), C.int(i)) != 0
}

// AddProbe

func (p *GstPad) CanLink(sinkpad *GstPad) bool {
	return C.gst_pad_can_link(p.native(), sinkpad.native()) != 0
}

func (p *GstPad) Chain(buffer *GstBuffer) GstFlowReturn {
	return GstFlowReturn(C.gst_pad_chain(p.native(), buffer.native()))
}

func (p *GstPad) ChainList(list *GstBufferList) GstFlowReturn {
	return GstFlowReturn(C.gst_pad_chain_list(p.native(), list.native()))
}

func (p *GstPad) CheckReconfigure() bool {
	return C.gst_pad_check_reconfigure(p.native()) != 0
}

func (p *GstPad) CreateStreamId(parent *GstElement, stream_id string) string {
	i := C.CString(stream_id)
	defer C.free(unsafe.Pointer(i))
	return C.GoString(C.gst_pad_create_stream_id(p.native(), parent.native(), (*C.gchar)(i)))
}

//CreateStreamIdPrintf

func (p *GstPad) EventDefault(parent *GstObject, event *GstEvent) bool {
	return C.gst_pad_event_default(p.native(), parent.native(), event.native()) != 0
}

// func (p *GstPad) Forward() // Idk how to implement user_data

func (p *GstPad) GetAllowedCaps() *GstCaps {
	return C.gst_pad_get_allowed_caps(p.native()).impl()
}

func (p *GstPad) GetCurrentCaps() *GstCaps {
	return C.gst_pad_get_current_caps(p.native()).impl()
}

func (p *GstPad) GetDirection() GstPadDirection {
	return GstPadDirection(C.gst_pad_get_direction(p.native()))
}

func (p *GstPad) GetElementPrivate() uintptr {
	return uintptr(C.gst_pad_get_element_private(p.native()))
}

func (p *GstPad) GetLastFlowReturn() GstFlowReturn {
	return GstFlowReturn(C.gst_pad_get_last_flow_return(p.native()))
}

func (p *GstPad) GetOffset() int64 {
	return int64(C.gst_pad_get_offset(p.native()))
}

func (p *GstPad) GetPadTemplate() *GstPadTemplate {
	return C.gst_pad_get_pad_template(p.native()).impl()
}

func (p *GstPad) GetPadTemplateCaps() *GstCaps {
	return C.gst_pad_get_pad_template_caps(p.native()).impl()
}

func (p *GstPad) GetParentElement() *GstElement {
	return C.gst_pad_get_parent_element(p.native()).impl()
}

func (p *GstPad) GetPeer() *GstPad {
	return C.gst_pad_get_peer(p.native()).impl()
}

//GetRange

//GetSingleInternalLink: idk

func (p *GstPad) GetStickyEvent(event_type GstEventType, idx uint) *GstEvent {
	return C.gst_pad_get_sticky_event(p.native(), C.GstEventType(event_type), C.uint(idx)).impl()
}

func (p *GstPad) GetStream() *GstStream {
	return C.gst_pad_get_stream(p.native()).impl()
}

func (p *GstPad) GetStreamId() string {
	return C.GoString(C.gst_pad_get_stream_id(p.native()))
}

func (p *GstPad) GetTaskState() GstTaskState {
	return GstTaskState(C.gst_pad_get_task_state(p.native()))
}

func (p *GstPad) HasCurrentCaps() bool {
	return C.gst_pad_has_current_caps(p.native()) != 0
}

func (p *GstPad) IsActive() bool {
	return C.gst_pad_is_active(p.native()) != 0
}

func (p *GstPad) IsBlocked() bool {
	return C.gst_pad_is_blocked(p.native()) != 0
}

func (p *GstPad) IsBlocking() bool {
	return C.gst_pad_is_blocking(p.native()) != 0
}

func (p *GstPad) IsLinked() bool {
	return C.gst_pad_is_linked(p.native()) != 0
}

func (p *GstPad) Link(sinkpad *GstPad) GstPadLinkReturn {
	return GstPadLinkReturn(C.gst_pad_link(p.native(), sinkpad.native()))
}

func (p *GstPad) LinkFull(sinkpad *GstPad, flags GstPadLinkCheck) GstPadLinkReturn {
	return GstPadLinkReturn(C.gst_pad_link_full(p.native(), sinkpad.native(), C.GstPadLinkCheck(flags)))
}

func (p *GstPad) LinkMaybeGhostingFull(sinkpad *GstPad, flags GstPadLinkCheck) bool {
	return C.gst_pad_link_maybe_ghosting_full(p.native(), sinkpad.native(), C.GstPadLinkCheck(flags)) != 0
}

func (p *GstPad) MarkReconfigure() {
	C.gst_pad_mark_reconfigure(p.native())
}

func (p *GstPad) NeedsReconfigure() bool {
	return C.gst_pad_needs_reconfigure(p.native()) != 0
}

func (p *GstPad) PauseTask() bool {
	return C.gst_pad_pause_task(p.native()) != 0
}

func (p *GstPad) PeerQuery(query *GstQuery) bool {
	return C.gst_pad_peer_query(p.native(), query.native()) != 0
}

func (p *GstPad) PeerQueryAcceptCaps(caps *GstCaps) bool {
	return C.gst_pad_peer_query_accept_caps(p.native(), caps.native()) != 0
}

func (p *GstPad) PeerQueryCaps(filter *GstCaps) *GstCaps {
	return C.gst_pad_peer_query_caps(p.native(), filter.native()).impl()
}

func (p *GstPad) PeerQueryConvert(src_format GstFormat, src_val int64, dest_format GstFormat) (bool, int64) {
	var dest_val C.long
	return C.gst_pad_peer_query_convert(p.native(), C.GstFormat(src_format), C.long(src_val),
		C.GstFormat(dest_format), (*C.gint64)(&dest_val)) != 0, int64(dest_val)
}

func (p *GstPad) PeerQueryDuration(format GstFormat) (bool, int64) {
	var duration C.long
	return C.gst_pad_peer_query_duration(p.native(), C.GstFormat(format), (*C.gint64)(&duration)) != 0, int64(duration)
}

func (p *GstPad) PeerQueryPosition(format GstFormat) (bool, int64) {
	var cur C.long
	return C.gst_pad_peer_query_position(p.native(), C.GstFormat(format), (*C.gint64)(&cur)) != 0, int64(cur)
}

func (p *GstPad) ProxyQueryAcceptCaps(query *GstQuery) bool {
	return C.gst_pad_proxy_query_accept_caps(p.native(), query.native()) != 0
}

func (p *GstPad) ProxyQueryCaps(query *GstQuery) bool {
	return C.gst_pad_proxy_query_caps(p.native(), query.native()) != 0
}

// PullRange: **GstBuffer

func (p *GstPad) Push(buffer *GstBuffer) GstFlowReturn {
	return GstFlowReturn(C.gst_pad_push(p.native(), buffer.native()))
}

func (p *GstPad) PushEvent(event *GstEvent) bool {
	return C.gst_pad_push_event(p.native(), event.native()) != 0
}

// PushList: BufferList

func (p *GstPad) Query(query *GstQuery) bool {
	return C.gst_pad_query(p.native(), query.native()) != 0
}

func (p *GstPad) QueryAcceptCaps(caps *GstCaps) bool {
	return C.gst_pad_query_accept_caps(p.native(), caps.native()) != 0
}

func (p *GstPad) QueryCaps(filter *GstCaps) *GstCaps {
	return C.gst_pad_query_caps(p.native(), filter.native()).impl()
}

func (p *GstPad) QueryConvert(src_format GstFormat, src_val int64, dest_format GstFormat) (bool, int64) {
	var dest_val C.long
	return C.gst_pad_query_convert(p.native(), C.GstFormat(src_format), C.gint64(src_val),
		C.GstFormat(dest_format), (*C.gint64)(&dest_val)) != 0, int64(dest_val)
}

func (p *GstPad) QueryDefault(parent *GstObject, query *GstQuery) bool {
	return C.gst_pad_query_default(p.native(), parent.native(), query.native()) != 0
}

func (p *GstPad) QueryDuration(format GstFormat) (bool, int64) {
	var duration C.long
	return C.gst_pad_query_duration(p.native(), C.GstFormat(format), (*C.gint64)(&duration)) != 0, int64(duration)

}

func (p *GstPad) QueryPosition(format GstFormat) (bool, int64) {
	var cur C.long
	return C.gst_pad_query_position(p.native(), C.GstFormat(format), (*C.gint64)(&cur)) != 0, int64(cur)
}

func (p *GstPad) RemoveProbe(id uint64) {
	C.gst_pad_remove_probe(p.native(), C.gulong(id))
}

func (p *GstPad) SendEvent(event *GstEvent) bool {
	return C.gst_pad_send_event(p.native(), event.native()) != 0
}

// ActivateFunctionFull: callback

func (p *GstPad) SetActive(active bool) bool {
	a := 0
	if active {
		a = 1
	}
	return C.gst_pad_set_active(p.native(), C.gboolean(a)) != 0
}

//SetChainFunctionFull: callback

//SetChainListFunctionFull: callback

//SetElementPrivate: gpointer

//SetEventFullFunctionFull: callback

//SetEventFunctionFull: callback

//SetGetrangeFunctionFull: callback

//SetIterateInternalLinksFunctionFull: callback

//SetLinkFunctionFull: callback

func (p *GstPad) SetOffset(offset int64) {
	C.gst_pad_set_offset(p.native(), C.gint64(offset))
}

//SetQueryFunctionFull: callback

//SetUnlinkFunctionFull: callback

//StartTask: callback

//StickyEventsForeach: callback

func (p *GstPad) StopTask() bool {
	return C.gst_pad_stop_task(p.native()) != 0
}

func (p *GstPad) StoreStickyEvent(event *GstEvent) GstFlowReturn {
	return GstFlowReturn(C.gst_pad_store_sticky_event(p.native(), event.native()))
}

func (p *GstPad) Unlink(sinkpad *GstPad) bool {
	return C.gst_pad_unlink(p.native(), sinkpad.native()) != 0
}

func (p *GstPad) UseFixedCaps() {
	C.gst_pad_use_fixed_caps(p.native())
}

// GstPadLinkReturn

func (p GstPadLinkReturn) GetName() string {
	return C.GoString(C.gst_pad_link_get_name(C.GstPadLinkReturn(p)))
}

// GstPadProbeInfo
// TODO

//TODO: Task, Event, Stream, PadTemplate(Static), Query, Format
