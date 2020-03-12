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
	return GstPadDirection(C.gst_pad_get_direction(p.native))
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

func (p *GstPad) GetSingleInternalLink() *GstPad {
	// TODO:idk
	//return C.gst_pad_get_single_internal_link(p.native()).impl()
}

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
	return C.gst_pad_has_current_caps() != 0
}

func (p *GstPad) IsActive() bool {
	return C.gst_pad_is_active() != 0
}

func (p *GstPad) IsBlocked() bool {
	return C.gst_pad_is_blocked() != 0
}

func (p *GstPad) IsBlocking() bool {
	return C.gst_pad_is_blocking() != 0
}

func (p *GstPad) IsLinked() bool {
	return C.gst_pad_is_linked() != 0
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

//func (p *GstPad)
