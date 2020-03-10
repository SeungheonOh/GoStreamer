package gostreamer

import (
	"github.com/gotk3/gotk3/glib"
	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

//type GstElement C.GstElement
type GstElement struct {
	glib.Object
}

type GstElementFlags int
type GstState int
type GstStateChange int
type GstStateChangeReturn int
type GstURIType int

const (
	GST_ELEMENT_FLAG_LOCKED_STATE       GstElementFlags      = 16
	GST_ELEMENT_FLAG_SINK               GstElementFlags      = 32
	GST_ELEMENT_FLAG_SOURCE             GstElementFlags      = 64
	GST_ELEMENT_FLAG_PROVIDE_CLOCK      GstElementFlags      = 128
	GST_ELEMENT_FLAG_REQUIRE_CLOCK      GstElementFlags      = 256
	GST_ELEMENT_FLAG_INDEXABLE          GstElementFlags      = 512
	GST_ELEMENT_FLAG_LAST               GstElementFlags      = 16384
	GST_STATE_VOID_PENDING              GstState             = 0
	GST_STATE_NULL                      GstState             = 1
	GST_STATE_READY                     GstState             = 2
	GST_STATE_PAUSED                    GstState             = 3
	GST_STATE_PLAYING                   GstState             = 4
	GST_STATE_CHANGE_NULL_TO_READY      GstStateChange       = 10
	GST_STATE_CHANGE_READY_TO_PAUSED    GstStateChange       = 19
	GST_STATE_CHANGE_PAUSED_TO_PLAYING  GstStateChange       = 28
	GST_STATE_CHANGE_PLAYING_TO_PAUSED  GstStateChange       = 35
	GST_STATE_CHANGE_PAUSED_TO_READY    GstStateChange       = 26
	GST_STATE_CHANGE_READY_TO_NULL      GstStateChange       = 17
	GST_STATE_CHANGE_NULL_TO_NULL       GstStateChange       = 9
	GST_STATE_CHANGE_READY_TO_READY     GstStateChange       = 18
	GST_STATE_CHANGE_PAUSED_TO_PAUSED   GstStateChange       = 27
	GST_STATE_CHANGE_PLAYING_TO_PLAYING GstStateChange       = 36
	GST_STATE_CHANGE_FAILURE            GstStateChangeReturn = 0
	GST_STATE_CHANGE_SUCCESS            GstStateChangeReturn = 1
	GST_STATE_CHANGE_ASYNC              GstStateChangeReturn = 2
	GST_STATE_CHANGE_NO_PREROLL         GstStateChangeReturn = 3
)

const (
	GST_STATE_UNKNOWN GstURIType = 0
	GST_STATE_SINK    GstURIType = 1
	GST_STATE_SRC     GstURIType = 2
)

func (s *C.GstElement) impl() *GstElement {
	return (*GstElement)(unsafe.Pointer(s))
}

func (e *GstElement) native() *C.GstElement {
	return (*C.GstElement)(unsafe.Pointer(e))
}

func (e *GstElement) AbortState() {
	C.gst_element_abort_state(e.native())
}

func (e *GstElement) AddPad(pad *GstPad) bool {
	return C.gst_element_add_pad(e.native(), pad.native()) != 0
}

func (e *GstElement) ChangeState(transition GstStateChange) GstStateChangeReturn {
	return GstStateChangeReturn(C.gst_element_change_state(e.native(), C.GstStateChange(transition)))
}

func (e *GstElement) ContinueState(ret GstStateChangeReturn) GstStateChangeReturn {
	return GstStateChangeReturn(C.gst_element_continue_state(e.native(), C.GstStateChangeReturn(ret)))
}

func (e *GstElement) CreateAllPads() {
	C.gst_element_create_all_pads(e.native())
}

func (e *GstElement) GetBaseTime() GstClockTime {
	return GstClockTime(C.gst_element_get_base_time(e.native()))
}

func (e *GstElement) GetBus() *GstBus {
	return (C.gst_element_get_bus(e.native())).impl()
}

func (e *GstElement) GetClock() *GstClock {
	return (C.gst_element_get_clock(e.native())).impl()
}

func (e *GstElement) GetContext(context_type string) *GstContext {
	c := C.CString(context_type)
	defer C.free(unsafe.Pointer(c))
	return (*GstContext)(C.gst_element_get_context(e.native(), (*C.gchar)(c)))
}

func (e *GstElement) GetContextUnlocked(context_type string) *GstContext {
	c := C.CString(context_type)
	defer C.free(unsafe.Pointer(c))
	return (*GstContext)(C.gst_element_get_context_unlocked(e.native(), (*C.gchar)(c)))
}

func (e *GstElement) GetMetadata(key string) string {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	return C.GoString(C.gst_element_get_metadata(e.native(), (*C.gchar)(k)))
}

func (e *GstElement) GetPadTemplate(name string) *GstPadTemplate {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return (*GstPadTemplate)(C.gst_element_get_pad_template(e.native(), (*C.gchar)(n)))
}

func (e *GstElement) GetStartTime() GstClockTime {
	return GstClockTime(C.gst_element_get_start_time(e.native()))
}

func (e *GstElement) GetState(timeout uint64) (GstState, GstState, GstStateChangeReturn) {
	var state C.GstState
	var pending C.GstState
	ret := GstStateChangeReturn(C.gst_element_get_state(e.native(), &state, &pending, C.ulong(timeout)))
	return GstState(state), GstState(pending), GstStateChangeReturn(ret)
}

func (e *GstElement) GetStaticPad(name string) *GstPad {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return (*GstPad)(C.gst_element_get_static_pad(e.native(), (*C.gchar)(n)))
}

func (e *GstElement) IsLockedState() bool {
	return C.gst_element_is_locked_state(e.native()) != 0
}

func (e *GstElement) Link(dests ...*GstElement) bool {
	for _, dest := range dests {
		if C.gst_element_link(e.native(), dest.native()) == 0 {
			return false
		}
	}
	return true
}

func (e *GstElement) LinkFiltered(dest *GstElement, filter *GstCaps) bool {
	return C.gst_element_link_filtered(e.native(), dest.native(), filter.native()) != 0
}

func (e *GstElement) Unlink(dests ...*GstElement) {
	for _, dest := range dests {
		C.gst_element_unlink(e.native(), dest.native())
	}
}

func (e *GstElement) LinkPads(dest *GstElement, srcpadname, destpadname string) bool {
	s := (*C.gchar)(C.CString(srcpadname))
	d := (*C.gchar)(C.CString(destpadname))
	defer C.g_free((C.gpointer)(unsafe.Pointer(s)))
	defer C.g_free((C.gpointer)(unsafe.Pointer(d)))
	return C.gst_element_link_pads(e.native(), s, dest.native(), d) != 0
}

func (e *GstElement) LinkPadsFiltered(dest *GstElement, srcpadname, destpadname string, filter *GstCaps) bool {
	s := (*C.gchar)(C.CString(srcpadname))
	d := (*C.gchar)(C.CString(destpadname))
	defer C.g_free((C.gpointer)(unsafe.Pointer(s)))
	defer C.g_free((C.gpointer)(unsafe.Pointer(d)))
	return C.gst_element_link_pads_filtered(e.native(), s, dest.native(), d, filter.native()) != 0
}

func (e *GstElement) UnlinkPads(dest *GstElement, srcpadname, destpadname string) {
	s := C.CString(srcpadname)
	d := C.CString(destpadname)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(d))
	C.gst_element_unlink_pads(e.native(), (*C.gchar)(s), dest.native(), (*C.gchar)(d))
}

// TODO:PadsFull

func (e *GstElement) RemovePad(pad *GstPad) bool {
	return C.gst_element_remove_pad(e.native(), pad.native()) != 0
}

func (e *GstElement) RequestPad(templ *GstPadTemplate, name string, caps *GstCaps) *GstPad {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return (*GstPad)(C.gst_element_request_pad(e.native(), templ.native(), (*C.gchar)(n), caps.native()))
}

func (e *GstElement) LostState() {
	C.gst_element_lost_state(e.native())
}

func (e *GstElement) MessageFull(message_type GstMessageType, domain uint32, code int, text, debug, file, function string, line int) {
	t := C.CString(text)
	d := C.CString(debug)
	fi := C.CString(file)
	fu := C.CString(function)
	defer C.free(unsafe.Pointer(t))
	defer C.free(unsafe.Pointer(d))
	defer C.free(unsafe.Pointer(fi))
	defer C.free(unsafe.Pointer(fu))
	C.gst_element_message_full(e.native(), C.GstMessageType(message_type), C.GQuark(domain), C.gint(C.int(code)), (*C.gchar)(t), (*C.gchar)(d), (*C.gchar)(fi), (*C.gchar)(fu), C.gint(C.int(line)))
}

func (e *GstElement) NoMorePads() {
	C.gst_element_no_more_pads(e.native())
}

func (e *GstElement) PostMessage(message *GstMessage) bool {
	return C.gst_element_post_message(e.native(), message.native()) != 0
}

func (e *GstElement) ProvideClock() *GstClock {
	c := C.gst_element_provide_clock(e.native())
	return c.impl()
}

func (e *GstElement) Query(query *GstQuery) bool {
	return C.gst_element_query(e.native(), query.native()) != 0
}

func (e *GstElement) SendEvent(event *GstEvent) bool {
	return C.gst_element_send_event(e.native(), event.native()) != 0
}

func (e *GstElement) SetBaseTime(time GstClockTime) {
	C.gst_element_set_base_time(e.native(), C.GstClockTime(time))
}

func (e *GstElement) SetBus(bus *GstBus) {
	C.gst_element_set_bus(e.native(), bus.native())
}

func (e *GstElement) SetClock(clock *GstClock) bool {
	return C.gst_element_set_clock(e.native(), clock.native()) != 0
}

func (e *GstElement) SetContext(context *GstContext) {
	C.gst_element_set_context(e.native(), context.native())
}

func (e *GstElement) SetLockedState(locked_state bool) bool {
	b := 0
	if locked_state {
		b = 1
	}
	return C.gst_element_set_locked_state(e.native(), C.gboolean(b)) != 0
}

func (e *GstElement) SetStartTime(time GstClockTime) {
	C.gst_element_set_start_time(e.native(), C.GstClockTime(time))
}

func (e *GstElement) SetState(state GstState) GstStateChangeReturn {
	return GstStateChangeReturn(C.gst_element_set_state(e.native(), C.GstState(state)))
}

func (e *GstElement) SyncStateWithParent() bool {
	return C.gst_element_sync_state_with_parent(e.native()) != 0
}

func GstElementMakeFromURI(uri_type GstURIType, uri string, elementname string) *GstElement {
	u := C.CString(uri)
	e := C.CString(elementname)
	defer C.free(unsafe.Pointer(u))
	defer C.free(unsafe.Pointer(e))
	element := C.gst_element_make_from_uri(C.GstURIType(uri_type), (*C.gchar)(u), (*C.gchar)(e), nil)
	return (*GstElement)(unsafe.Pointer(element))
}

func GstElementStateChangeReturnGetName(state_ret GstStateChangeReturn) string {
	return C.GoString(C.gst_element_state_change_return_get_name(C.GstStateChangeReturn(state_ret)))
}

func GstElementStateGetName(state GstState) string {
	return C.GoString(C.gst_element_state_get_name(C.GstState(state)))
}

func (state GstState) GetName() string {
	return C.GoString(C.gst_element_state_get_name(C.GstState(state)))
}

func (state_ret GstStateChangeReturn) GetName() string {
	return C.GoString(C.gst_element_state_change_return_get_name(C.GstStateChangeReturn(state_ret)))
}
